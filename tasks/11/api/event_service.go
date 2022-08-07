package api

import (
	"time"

	"task11/models"
)

//EventService ...
type EventService struct {
	s Store
}

//NewEventService для дарльнейшего роутинга с http
func NewEventService() *EventService {
	return &EventService{s: NewEventStorage()}
}

// Интерфейс для реализации
type Store interface {
	CreateEvent(e models.Event) error
	UpdateEvent(e models.Event) error
	DeleteEvent(eventID string) error
	GetEventsForPeriod(userID string, p1 time.Time, p2 time.Time) ([]models.Event, error)
}

//CreateEvent реализует интерфейс Store
func (s *EventService) CreateEvent(e models.Event) error {
	return s.s.CreateEvent(e)
}

// UpdateEvent реализует интерфейс Store
func (s *EventService) UpdateEvent(e models.Event) error {
	return s.s.UpdateEvent(e)
}

// DeleteEvent реализует интерфейс Store
func (s *EventService) DeleteEvent(eventID string) error {
	return s.s.DeleteEvent(eventID)
}

// EventsForDay реализует интерфейс Store
func (s *EventService) EventsForDay(userID string) ([]models.Event, error) {
	es, err := s.s.GetEventsForPeriod(userID, time.Now(), time.Now().Add(24*time.Hour))
	return es, err
}

// EventsForWeek реализует интерфейс Store
func (s *EventService) EventsForWeek(userID string) ([]models.Event, error) {
	es, err := s.s.GetEventsForPeriod(userID, time.Now(), time.Now().Add(7*24*time.Hour))
	return es, err
}

// EventsForMonth реализует интерфейс Store
func (s *EventService) EventsForMonth(userID string) ([]models.Event, error) {
	es, err := s.s.GetEventsForPeriod(userID, time.Now(), time.Now().Add(30*24*time.Hour))
	return es, err
}
