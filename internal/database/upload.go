package database

type Upload struct {
	Id        int64  `xorm:"'id' pk autoincr"`
	UserId    int64  `xorm:"'user_id' index(idx_user_id)"`
	MimeType  string `xorm:"'mime_type'"`
	Size      int64  `xorm:"'size'"`
	Filepath  string `xorm:"'file_path'"`
	CreatedAt int64  `xorm:"'created_at' created"`
}
