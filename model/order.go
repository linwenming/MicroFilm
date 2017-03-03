package model

import (
	"github.com/gocraft/dbr"
	"fmt"
	"time"
	"MicroFilm/util"
)

type OrderDetail struct {
	Id           int64   `json:"id"`
	OrderSn      string   `db:"order_sn" json:"orderSn"`
	Uid          int64    `json:"uid"`
	Status       int64    `json:"status"`
	CreateTime   int64    `db:"create_time" json:"createTime"`
	PayTime      int64    `db:"pay_time" json:"payTime"`
	AgentId      int64    `json:"status"`
	ProductId    int64    `db:"product_id" json:"productId"`
	ProductName  string    `db:"product_name" json:"productName"`
	ProductPrice int64    `db:"product_price" json:"productPrice"`
	TotalPrice   int64    `db:"total_price" json:"totalPrice"`
	Quantity     int64    `json:"quantity"`
	Platform     string    `json:"platform"`
	ReferUrl     string    `db:"refer_url" json:"referUrl"`
}

type OrderLog struct {
	Id         int64   `json:"id"`
	OrderSn    string   `db:"order_sn" json:"orderSn"`
	Content    string    `json:"content"`
	CreateTime int64    `db:"create_time"  json:"createTime"`
}

func (m *OrderDetail) Save(tx *dbr.Tx) error {

	sn, err0 := NextBillNumber(tx)
	if err0 != nil {
		return err0
	}
	m.OrderSn = sn

	_, err1 := tx.InsertInto("bis_order").
		Columns(util.BuildColumnName(m)...).
		Record(m).
		Exec()

	return err1
}

func (m *OrderLog) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("bis_order_log").
		Columns(util.BuildColumnName(m)...).
		Record(m).
		Exec()

	return err
}

func NextBillNumber(tx *dbr.Tx) (string, error) {
	var id int64
	err := tx.SelectBySql("SELECT auto_increment as value FROM information_schema.`TABLES` WHERE TABLE_SCHEMA='micro_movie' AND TABLE_NAME='bis_order'").
		LoadValue(&id)

	if (err == nil) {
		y, m, d := time.Now().Date()
		return fmt.Sprintf("%d%02d%02d%08d", y, m, d, id), nil
	} else {
		return "", err
	}
}

func (m *OrderDetail) LoadBySn(tx *dbr.Tx, sn string) error {

	return tx.Select(util.BuildColumnName(m)...).
		From("bis_order").
		Where("orderSn = ?", sn).
		LoadStruct(m)
}

func (m *OrderDetail) UpdateBy(tx *dbr.Tx, value map[string]interface{}) error {

	_, err := tx.Update("bis_order").
		SetMap(value).
		Where("id = ?", m.Id).
		Exec()

	return err
}

type OrderDetails []OrderDetail

func (m *OrderDetails) Load(tx *dbr.Tx, pageSize uint64, pageNumber uint64) error {

	return tx.Select("*").
		From("bis_order").
		Limit(pageSize).Offset((pageNumber - 1) * pageSize).
		LoadStruct(m)
}