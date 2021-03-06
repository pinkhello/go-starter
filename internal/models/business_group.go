package models

import (
	"time"
)

type BusinessGroup struct {
	Id             int64     `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
	Uid            string    `json:"uid"`
	BusinessUid    string    `json:"business_uid"`
	EnableMall     bool      `json:"enable_mall"`
	EndOfMall      time.Time `json:"end_of_mall"`
	DeliveryConfig int       `json:"delivery_config"`
	Phones         string    `json:"phones"`
	MallQrcode     string    `json:"mall_qrcode"`
}
