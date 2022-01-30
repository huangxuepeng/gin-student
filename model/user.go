package model

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        int32     `gorm: "primarykey"`
	CreatedAt time.Time `gorm: "column: add_time"`
	UpdatedAt time.Time `gorm: "column: add_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

type Student struct {
	Base
	Mobile   string `gorm: "index:idx_mobile;unique; type:varchar(11); not null"`
	Password string `gorm: "type:varchar(100); not null"`
	Name     string `gorm: "type:varchar(20)"`
	StuNum   string `gorm:"type:varchar(9); not null comment '学号'"`
	Gender   string `gorm:"column:gender;default:male;type:varchar(6) comment 'female表示女, male表示男'"`
	Role     int    `gorm:"column:role; default:1; type:int comment '1表示普通用户, 2表示管理员'"`
}
