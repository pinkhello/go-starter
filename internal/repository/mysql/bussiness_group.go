package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
	"go-starter/internal/models"
	"go-starter/utils"
	"strconv"
)

type BusinessGroupRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (businessGroups []models.BusinessGroup, csr string, err error)
	GetById(ctx context.Context, id int64) (businessGroup models.BusinessGroup, err error)
	//Update(ctx context.Context, businessGroup *BusinessGroup) (err error)
	//Store(ctx context.Context, businessGroup *BusinessGroup) (err error)
	//Delete(ctx context.Context, id int64) (err error)
}

type mysqlBusinessGroupRepository struct {
	Conn *sql.DB
}

func NewBusinessGroupRepository(conn *sql.DB) BusinessGroupRepository {
	if conn == nil {
		panic("Database Connections is null")
	}
	return &mysqlBusinessGroupRepository{conn}
}

func (m *mysqlBusinessGroupRepository) fetch(ctx context.Context, query string, args ...interface{}) (businessGroups []models.BusinessGroup, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	businessGroups = make([]models.BusinessGroup, 0)
	for rows.Next() {
		t := models.BusinessGroup{}
		err = rows.Scan(
			&t.Id,
			&t.CreatedAt,
			&t.UpdatedAt,
			&t.DeletedAt,
			&t.Uid,
			&t.BusinessUid,
			&t.EnableMall,
			&t.EndOfMall,
			&t.DeliveryConfig,
			&t.Phones,
			&t.MallQrcode,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		businessGroups = append(businessGroups, t)
	}
	return businessGroups, nil
}

func (m *mysqlBusinessGroupRepository) Fetch(ctx context.Context, cursor string, num int64) (businessGroups []models.BusinessGroup, csr string, err error) {
	qb := squirrel.Select("id", "created_at", "updated_at", "deleted_at", "uid", "business_uid", "enable_mall", "end_of_mall", "delivery_config", "phones", "mall_qrcode").From("business_group")
	qb = qb.OrderBy("id DESC").Limit(uint64(num))

	if cursor != "" {
		decodedCursor, err := strconv.ParseInt(cursor, 10, 64)
		if err != nil && cursor != "" {
			return nil, "", utils.ErrBadParamInput
		}
		qb = qb.Where(squirrel.Lt{
			"id": decodedCursor,
		})
	}

	query, args, err := qb.ToSql()

	if err != nil {
		return
	}

	businessGroups, err = m.fetch(ctx, query, args...)
	if err != nil {
		return nil, "", err
	}
	csr = cursor
	if len(businessGroups) > 0 {
		csr = fmt.Sprintf("%d", businessGroups[len(businessGroups)-1].Id)
	}
	return

}

func (m *mysqlBusinessGroupRepository) GetById(ctx context.Context, id int64) (businessGroup models.BusinessGroup, err error) {
	query := `SELECT id, created_at,updated_at,deleted_at,uid,business_uid,enable_mall,end_of_mall,delivery_config,phones,mall_qrcode
  						FROM business_group WHERE id = ?`
	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return models.BusinessGroup{}, err
	}
	if len(list) > 0 {
		businessGroup = list[0]
	} else {
		return businessGroup, utils.ErrNotFound
	}
	return
}

//Update(ctx context.Context, businessGroup *BusinessGroup) (err error)
//Store(ctx context.Context, businessGroup *BusinessGroup) (err error)
//Delete(ctx context.Context, id int64) (err error)
