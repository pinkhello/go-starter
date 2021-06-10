package models

import (
	"time"
)

type BusinessGroup struct {
	ID        int64     `json:"id" xorm:"'id' pk autoincr"`
	Name      string    `json:"name" xorm:"VARCHAR(255) not null comment('商户名称')"`
	Status    string    `json:"status" xorm:"VARCHAR(10) not null comment('状态｜invalid｜valid')"`
	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
}
