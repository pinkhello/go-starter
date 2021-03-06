package service

import (
	"context"
	"go-starter/internal/models"
	"go-starter/internal/repository/mysql"
	"time"
)

type (
	BusinessGroupService interface {
		Fetch(ctx context.Context, cursor string, num int64) (businessGroups []models.BusinessGroup, csr string, err error)
		GetById(ctx context.Context, id int64) (businessGroup models.BusinessGroup, err error)
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

func (s *BusinessGroupServiceImpl) Fetch(ctx context.Context, cursor string, num int64) (businessGroups []models.BusinessGroup, csr string, err error) {
	if num == 0 {
		num = 10
	}
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	if err != nil {
		return nil, "", err
	}
	businessGroups, csr, err = s.repo.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}
	return

}

func (s *BusinessGroupServiceImpl) GetById(ctx context.Context, id int64) (businessGroup models.BusinessGroup, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	businessGroup, err = s.repo.GetById(ctx, id)
	return
}
