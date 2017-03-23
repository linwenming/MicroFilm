package model

import (
	"time"

	"github.com/gocraft/dbr"
	"MicroFilm/util"
)

type User struct {
	Id        int64   `json:"id"`
	UserType  int   `db:"user_type" json:"userType"`
	LoginType int   `db:"login_type"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Openid    string  `json:"openid"`
	Active    int   `json:"active"`
	Nickname  string  `json:"nickname"`
	Face      string  `json:"face"`
	Gender    int   `json:"gender"`
	Age       int   `json:"age"`
	Phone     string  `json:"phone"`
	Email     string  `json:"email"`
	QQ        string  `db:"qq" json:"qq"`
	Weixin    string  `json:"weixin"`
	CreatedAt int64  `db:"created_at" json:"createdAt"`
}

func NewUser() *User {
	return &User{
		UserType:  0, //临时用户、会员、高级会员、代理商用户
		LoginType: 0, //手机、邮箱、微信、QQ
		Active:  0,
		CreatedAt: time.Now().Unix(),
	}
}

func (m *User) Save(tx *dbr.Tx) error {

	_, err := tx.InsertInto("sys_users").
		Columns(util.BuildColumnName(m)...).
		Record(m).
		Exec()

	return err
}

func (m *User) Update(tx *dbr.Tx) error {

	_, err := tx.Update("sys_users").
		SetMap(util.StructMap(m)).
		Where("id = ?", m.Id).
		Exec()

	return err
}

func (m *User) Load(tx *dbr.Tx, id int64) error {

	return tx.Select("*").
		From("sys_users").
		Where("id = ?", id).
		LoadStruct(m)
}

func (m *User) LoadByUsername(tx *dbr.Tx, value string) error {

	return tx.Select("*").
		From("sys_users").
		Where("username = ?", value).
		LoadStruct(m)
}

type Users []User

func (m *Users) Load(tx *dbr.Tx, active int) error {

	return tx.Select("*").
		From("member").
		Where(dbr.Eq("active", active)).
		LoadStruct(m)
}

