package cmd

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/softwarespot/public-holidays/internal/env"
	"github.com/softwarespot/public-holidays/internal/holidays"
	"github.com/softwarespot/public-holidays/internal/logging"
	"github.com/softwarespot/public-holidays/internal/service"
	"github.com/softwarespot/public-holidays/internal/service/middleware"
	"github.com/softwarespot/public-holidays/internal/version"
)

func cmdServer(logger logging.Logger) error {
	port := env.Get("SERVER_PORT", "10000")
	maxConcurrency, err := parseMaxConcurrency(env.Get("SERVER_MAX_CONCURRENCY", "500"))
	if err != nil {
		return err
	}

	rw := service.NewResponseWriter(logger)
	s := service.NewServer(":"+port, logger)
	s.Use(
		middleware.NewPanicRecovery(
			func(w http.ResponseWriter, r *http.Request, err error) {
				rw.ErrorAsJSON(w, r, err)
			},
			logger,
		),
		middleware.NewMetrics(logger),
		middleware.NewConcurrentRequests(int32(maxConcurrency), logger),
	)

	hm := holidays.NewManager()
	s.HandleFunc("GET /holidays/v1/{countryCode}/{year}", func(w http.ResponseWriter, r *http.Request) error {
		headers := w.Header()
		headers.Set("Access-Control-Allow-Origin", "*")
		headers.Set("Access-Control-Allow-Methods", "GET")

		// Disable the API from being indexed by search engines e.g. Google.
		// See URL: https://developers.google.com/search/reference/robots_meta_tag
		headers.Set("X-Robots-Tag", "none")

		code, err := holidays.NewCountryCode(r.PathValue("countryCode"))
		if err != nil {
			return service.NewError(err, http.StatusBadRequest)
		}

		year, err := parseYear(r.PathValue("year"))
		if err != nil {
			return service.NewError(err, http.StatusBadRequest)
		}

		res, err := hm.Get(code, year)
		if err != nil {
			return service.NewError(err, http.StatusBadRequest)
		}
		return rw.WriteAsJSON(w, r, res)
	})
	s.HandleErrorFunc("GET /holidays/v1/{countryCode}/{year}", func(w http.ResponseWriter, r *http.Request, err error) {
		rw.ErrorAsJSON(w, r, err)
	})

	s.HandleFunc("GET /holidays/v1/version", func(w http.ResponseWriter, r *http.Request) error {
		headers := w.Header()
		headers.Set("Access-Control-Allow-Origin", "*")
		headers.Set("Access-Control-Allow-Methods", "GET")

		// Disable the API from being indexed by search engines e.g. Google.
		// See URL: https://developers.google.com/search/reference/robots_meta_tag
		headers.Set("X-Robots-Tag", "none")

		return rw.WriteAsJSON(w, r, map[string]any{
			"version":   version.Version,
			"buildTime": version.Time,
			"buildUser": version.User,
			"goVersion": version.GoVersion,
		})
	})

	s.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) error {
		if r.URL.Path != "/" {
			return service.NewError(errors.New("Not Found"), http.StatusNotFound)
		}

		http.Redirect(w, r, "https://github.com/softwarespot/public-holidays", http.StatusMovedPermanently)
		return nil
	})

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	return s.ListenAndServe(ctx)
}
