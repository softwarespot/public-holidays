package holidays

import "fmt"

type Service struct {
	// No mutex is needed, as this is written to only on initialization
	handlers map[CountryCode]handlerFunc
}

func New() *Service {
	s := &Service{
		handlers: map[CountryCode]handlerFunc{},
	}

	// Registered country code handlers
	s.register("DK", dk)
	s.register("FI", fi)
	s.register("NO", no)
	s.register("SE", se)

	return s
}

func (s *Service) register(code CountryCode, handler handlerFunc) {
	s.handlers[code] = handler
}

func (s *Service) Get(code CountryCode, year int) ([]Holiday, error) {
	handlerFn, ok := s.handlers[code]
	if !ok {
		return nil, fmt.Errorf("unsupported country code of %q", code)
	}
	return handlerFn(year)
}
