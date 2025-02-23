package user

import (
	"app-admin/app/initialize"
	"fmt"
	"time"
)

type AppUsers struct {
	UsersId       int64     ` gorm:"primaryKey" json:"usersId"`
	UsersNickname string    `json:"usersNickname"`
	UsersPassword string    `json:"usersPassword"`
	UsersPhone    string    `json:"usersPhone"`
	UsersCreate   time.Time `gorm:"type:datetime" json:"usersCreate"`
	UsersPortrait string    `json:"usersPortrait"`
	UsersClient   string    `json:"usersClient"`
	UsersSource   string    `json:"usersSource"`
	UserClose     int       `json:"userClose"`
}

// 查询单个用户
func GetUserByID(id uint) (*AppUsers, error) {
	var user AppUsers
	result := initialize.DB.First(&user, id)
	fmt.Println(&user)
	return &user, result.Error
}
