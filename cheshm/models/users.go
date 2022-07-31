package models

type TypeUser uint64

const (
	TypeUserAdmin = iota
	TypeUserNormal
)

type User struct {
	Id        int
	FirstName string   `orm:"size(100)"`
	LastName  string   `orm:"size(100)"`
	Email     string   `orm:"size(100)"`
	Type      TypeUser `orm:"default(1)"`
}

func (u *User) TableName() string {
	return "auth_user"
}
