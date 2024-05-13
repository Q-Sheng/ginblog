package model

import (
	"fmt"
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20); not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500); not null" json:"password" validate:"required,min=4,max=12" label:"用户名"`
	Role     int    `gorm:"type:int; DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
	//Username string `gorm:"type:varchar(20);not null" json:"username"`
	//Password string `gorm:"type:varchar(20);not null" json:"password"`
	//Role     int    `gorm:"type:int" json:"role"`
	//Username string `json:"username"`
	//Password string `json:"password"`
	//Role     int    `json:"role"`
}

// 查询用户是否存在
func CheckUser(name string) int {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

// 新增用户
func CreateUser(data *User) int {
	//新建数据库
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE

}

// 查询用户列表 涉及分页
func GetUsers(pageSize int, pageNum int) []User {
	var users []User

	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Errorf("ERROR:", err)
		return users
	}
	return users
}

// 编辑用户

// 删除用户
