package database

type User struct {
	Id       int64  `xorm:"pk autoincr"`
	Username string `xorm:"UNIQUE VARCHAR(255)"`
	Password string `xorm:"VARCHAR(255)"`
}

func (User) TableName() string {
	return `user`
}
