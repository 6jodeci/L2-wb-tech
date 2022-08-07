package api

import (
	"errors"
	"task11/models"
	"sync"
	"time"
)

var NotFound = errors.New("event not found")

// Структура для эвента
type EventStorage struct {
	events map[string]models.Event
	mx     *sync.RWMutex
}

// Реализует структуру EventStorage
func NewEventStorage() *EventStorage {
	return &EventStorage{
		events: make(map[string]models.Event),
		mx:     &sync.RWMutex{},
	}
}

// Создает новый эвент и пишет в EventStorage
func (s *EventStorage) CreateEvent(e models.Event) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	if _, ok := s.events[e.ID]; !ok {
		s.events[e.ID] = e
		return nil
	}
	return NotFound
}

// UpdateEvent обновляет эвент и пишет в EventStorage
func (s *EventStorage) UpdateEvent(e models.Event) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	if _, ok := s.events[e.ID]; ok {
		s.events[e.ID] = e
		return nil
	}
	return NotFound
}

// DeleteEvent удаляет эвент и пишет в EventStorage
func (s *EventStorage) DeleteEvent(eventID string) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	if _, ok := s.events[eventID]; ok {
		delete(s.events, eventID)
		return nil
	}
	return NotFound
}

// GetEventsForPeriod возвращает эвент за нужный период 
func (s *EventStorage) GetEventsForPeriod(userID string, p1 time.Time, p2 time.Time) ([]models.Event, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	res := make([]models.Event, 0)
	for _, v := range s.events {
		if v.Date.Before(p2) && v.Date.After(p1) && v.CreatorID == userID {
			res = append(res, v)
		}
	}
	return res, nil
}