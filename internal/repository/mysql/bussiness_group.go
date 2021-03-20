package mysql

import (
	"context"
	"go-starter/internal/models"
	"go-starter/utils"
	"xorm.io/xorm"
)

type BusinessGroupRepository interface {
	GetByID(ctx context.Context, id int64) (res models.BusinessGroup, err error)
}

type mysqlBusinessGroupRepository struct {
	engine *xorm.Engine
}

func NewBusinessGroupRepository(engine *xorm.Engine) BusinessGroupRepository {
	if engine == nil {
		panic("Database engine is null")
	}
	return &mysqlBusinessGroupRepository{engine: engine}
}

func (m *mysqlBusinessGroupRepository) GetByID(ctx context.Context, id int64) (res models.BusinessGroup, err error) {
	has, err := m.engine.ID(id).Get(&res)
	if err != nil {
		return models.BusinessGroup{}, err
	}
	if !has {
		return res, utils.ErrNotFound
	}
	return
}
