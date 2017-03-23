package model

import (
	"github.com/gocraft/dbr"
	"MicroFilm/util"
)

// Go cron定时任务的用法
// http://www.cnblogs.com/zuxingyu/p/6023919.html
// https://github.com/robfig/cron

// 智能和功能强大的cron作业调度器
// http://www.ctolib.com/jobrunner.html


type Settlement struct {
	Id          int64  `json:"id"`
	Title       string   `json:"uid"`
	AgentId     int64   `json:"mid"`
	Commission  int64   `json:"cid"` // 佣金
	Quantity    int64   `json:"cid"` // 交易数量
	StartTime   int64   `db:"stime" json:"startTime"`
	EndTime     int64   `db:"etime" json:"endTime"`
	Mode        int64   `json:"mode"`
	Status      int64   `json:"status"`
	TotalAmount int64   `db:"total_amount" json:"totalAmount"`
	CreateTime  int64   `db:"create_time" json:"createTime"`
}

func (m *Settlement) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("bis_settlement").
		Columns(util.BuildColumnName(m)...).
		Record(m).
		Exec()
	return err
}

func (m *Settlement) Delete(tx *dbr.Tx, id int64) error {

	_, err := tx.DeleteFrom("bis_settlement").
		Where("id = ?", id).
		Exec()

	return err
}

func (m *Settlement) Update(tx *dbr.Tx) error {

	_, err := tx.Update("bis_settlement").
		SetMap(util.StructMap(m)).
		Where("id = ?", m.Id).
		Exec()

	return err
}

func (m *Settlement) Load(tx *dbr.Tx, id int64) error {

	return tx.Select(util.BuildColumnName(m)...).
		From("bis_settlement").
		Where("id = ?", id).
		LoadStruct(m)
}

type Settlements []Settlement

func (m *Settlements) LoadByTime(tx *dbr.Tx, stime int64,etime int64) error {

	return tx.Select(util.BuildColumnName(m)...).
		From("bis_settlement").
		Where("stime >= ? and etime <= ?", stime,etime).
		LoadStruct(m)
}

func (m *Settlements) LoadByPage(tx *dbr.Tx, pageSize uint64, pageNumber uint64) error {

	return tx.Select("*").
		From("bis_settlement").
		OrderBy("etime desc").
		Limit(pageSize).Offset((pageNumber - 1) * pageSize).
		LoadStruct(m)
}