package service

import (
	"context"
	"go-starter/internal/models"
	"go-starter/internal/repository/mysql"
	"time"
)

type (
	BusinessGroupService interface {
		GetByID(ctx context.Context, id int64) (businessGroup models.BusinessGroup, err error)
	}

	BusinessGroupServiceImpl struct {
		repo           mysql.BusinessGroupRepository
		contextTimeout time.Duration
	}
)

func NewBusinessGroupService(businessGroupRepo mysql.BusinessGroupRepository, timeout time.Duration) BusinessGroupService {
	if businessGroupRepo == nil {
		panic("BusinessGroup Repository is nil")
	}
	if timeout == 0 {
		panic("Timeout is empty")
	}
	return &BusinessGroupServiceImpl{
		repo:           businessGroupRepo,
		contextTimeout: timeout,
	}
}

func (s *BusinessGroupServiceImpl) GetByID(ctx context.Context, id int64) (businessGroup models.BusinessGroup, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	businessGroup, err = s.repo.GetByID(ctx, id)
	return
}
