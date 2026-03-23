package scheduler

import (
	"context"
	"log"
	"time"

	"bookstore/backend/internal/services"
)

type ReminderScheduler struct {
	service  *services.ReminderService
	interval time.Duration
}

func NewReminderScheduler(service *services.ReminderService, interval time.Duration) *ReminderScheduler {
	return &ReminderScheduler{
		service:  service,
		interval: interval,
	}
}

func (s *ReminderScheduler) Start(ctx context.Context) {
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	log.Printf("reminder scheduler started, interval=%s", s.interval)

	if err := s.service.Process(ctx); err != nil {
		log.Printf("reminder scheduler initial run error: %v", err)
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("reminder scheduler stopped")
			return
		case <-ticker.C:
			if err := s.service.Process(ctx); err != nil {
				log.Printf("reminder scheduler run error: %v", err)
			}
		}
	}
}
