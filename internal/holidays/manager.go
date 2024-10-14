package holidays

import "fmt"

type Manager struct {
	// No mutex is needed, as this is written to only on initialization
	handlers map[CountryCode]handlerFunc
}

func NewManager() *Manager {
	m := &Manager{
		handlers: map[CountryCode]handlerFunc{},
	}

	// Registered country code handlers
	m.Register("DK", dk)
	m.Register("FI", fi)
	m.Register("IS", is)
	m.Register("NO", no)
	m.Register("SE", se)

	return m
}

func (m *Manager) Register(code CountryCode, handler handlerFunc) {
	m.handlers[code] = handler
}

func (m *Manager) Get(code CountryCode, year int) ([]Holiday, error) {
	handlerFn, ok := m.handlers[code]
	if !ok {
		return nil, fmt.Errorf("unsupported country code of %q", code)
	}
	return handlerFn(year)
}
