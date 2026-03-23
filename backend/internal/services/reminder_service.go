package services

import (
	"context"
	"fmt"
	"log"

	"bookstore/backend/internal/repository"
)

type ReminderService struct {
	rentalRepo       *repository.RentalRepository
	notificationRepo *repository.NotificationRepository
}

func NewReminderService(
	rentalRepo *repository.RentalRepository,
	notificationRepo *repository.NotificationRepository,
) *ReminderService {
	return &ReminderService{
		rentalRepo:       rentalRepo,
		notificationRepo: notificationRepo,
	}
}

func (s *ReminderService) Process(ctx context.Context) error {
	if err := s.processExpiring(ctx); err != nil {
		return err
	}
	if err := s.processExpired(ctx); err != nil {
		return err
	}
	return nil
}

func (s *ReminderService) processExpiring(ctx context.Context) error {
	rentals, err := s.rentalRepo.FindExpiringSoon(ctx)
	if err != nil {
		return err
	}

	for _, rental := range rentals {
		exists, err := s.notificationRepo.ExistsByRentalAndType(ctx, rental.ID, "rent_expiring")
		if err != nil {
			return err
		}
		if exists {
			continue
		}

		message := fmt.Sprintf("Срок аренды книги \"%s\" скоро истекает.", rental.BookTitle)

		_, err = s.notificationRepo.Create(ctx, rental.UserID, rental.ID, "rent_expiring", message, "pending")
		if err != nil {
			return err
		}

		log.Printf("created expiring reminder for rental_id=%d", rental.ID)
	}

	return nil
}

func (s *ReminderService) processExpired(ctx context.Context) error {
	rentals, err := s.rentalRepo.FindExpiredActive(ctx)
	if err != nil {
		return err
	}

	for _, rental := range rentals {
		if err := s.rentalRepo.MarkExpired(ctx, rental.ID); err != nil {
			return err
		}

		exists, err := s.notificationRepo.ExistsByRentalAndType(ctx, rental.ID, "rent_expired")
		if err != nil {
			return err
		}
		if exists {
			continue
		}

		message := fmt.Sprintf("Срок аренды книги \"%s\" истёк.", rental.BookTitle)

		_, err = s.notificationRepo.Create(ctx, rental.UserID, rental.ID, "rent_expired", message, "pending")
		if err != nil {
			return err
		}

		log.Printf("created expired reminder for rental_id=%d", rental.ID)
	}

	return nil
}
