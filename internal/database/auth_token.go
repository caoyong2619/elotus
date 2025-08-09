package database

type AuthToken struct {
	Id        int64  `xorm:"'id' pk autoincr"`
	UserId    int64  `xorm:"'user_id' index(idx_user_id)"`
	Token     string `xorm:"'token'"`
	ExpiredAt int64  `xorm:"'expired_at'"`
	CreatedAt int64  `xorm:"'created_at' created"`
	UpdatedAt int64  `xorm:"'updated_at' updated"`
}

func (a AuthToken) TableName() string {
	return "auth_token"
}
