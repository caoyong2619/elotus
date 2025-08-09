package database

type User struct {
	Id        int64  `xorm:"pk autoincr"`
	Username  string `xorm:"UNIQUE VARCHAR(255)"`
	Password  string `xorm:"VARCHAR(255)"` // plain text password
	CreatedAt int64  `xorm:"'created_at' created"`
	UpdatedAt int64  `xorm:"'updated_at' updated"`
}

func (User) TableName() string {
	return `user`
}
