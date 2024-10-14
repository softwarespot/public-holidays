package cmd

import (
	"context"
	"net/http"
	"os"
	"syscall"

	"github.com/softwarespot/public-holidays/internal/env"
	"github.com/softwarespot/public-holidays/internal/helpers"
	"github.com/softwarespot/public-holidays/internal/holidays"
	"github.com/softwarespot/public-holidays/internal/logging"
	"github.com/softwarespot/public-holidays/internal/service"
	"github.com/softwarespot/public-holidays/internal/service/middleware"
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
		middleware.NewPanicRecovery(func(w http.ResponseWriter, r *http.Request, err error) { rw.ErrorAsJSON(w, r, err) }, logger),
		middleware.NewMetrics(logger),
		middleware.NewConcurrentRequests(int32(maxConcurrency), logger),
	)

	h := holidays.New()
	s.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) error {
		http.Redirect(w, r, "https://github.com/softwarespot/public-holidays", http.StatusMovedPermanently)
		return nil
	})

	routeHolidays := "GET /holidays/v1/{countryCode}/{year}"
	s.HandleFunc(routeHolidays, func(w http.ResponseWriter, r *http.Request) error {
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

		res, err := h.Get(code, year)
		if err != nil {
			return service.NewError(err, http.StatusBadRequest)
		}
		return rw.WriteAsJSON(w, r, res)
	})
	s.HandleErrorFunc(routeHolidays, func(w http.ResponseWriter, r *http.Request, err error) {
		rw.ErrorAsJSON(w, r, err)
	})

	ctx, cancel := helpers.SignalTrap(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	return s.ListenAndServe(ctx)
}
