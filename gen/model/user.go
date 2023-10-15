// Code generated by meta-egg. DO NOT EDIT.
// WILL BE replace after re-generated!
// DO NOT EDIT.
// Version: v1.9.2-IE
// Author: meta-egg
// Generated at: 2023-10-15 22:49

package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	ColUserID        = "id"         //nolint
	ColUserName      = "name"       // 用户名
	ColUserGender    = "gender"     // 性别
	ColUserAge       = "age"        // 年龄
	ColUserIsOnJob   = "is_on_job"  // 是否在职
	ColUserBirthday  = "birthday"   // 生日
	ColUserCreatedBy = "created_by" // 创建者
	ColUserCreatedAt = "created_at" // 创建时间
	ColUserUpdatedBy = "updated_by" // 更新者
	ColUserUpdatedAt = "updated_at" // 更新时间
	ColUserDeletedBy = "deleted_by" // 删除者
	ColUserDeletedAt = "deleted_at" // 删除时间
)

// 用户
type User struct {
	ID        uint64         `gorm:"column:id"`
	Name      *string        `gorm:"column:name"`       // 用户名 [default NULL]
	Gender    uint64         `gorm:"column:gender"`     // 性别
	Age       uint8          `gorm:"column:age"`        // 年龄
	IsOnJob   bool           `gorm:"column:is_on_job"`  // 是否在职
	Birthday  *time.Time     `gorm:"column:birthday"`   // 生日 [default NULL]
	CreatedBy *uint64        `gorm:"column:created_by"` // 创建者 [default NULL]
	CreatedAt time.Time      `gorm:"column:created_at"` // 创建时间
	UpdatedBy *uint64        `gorm:"column:updated_by"` // 更新者 [default NULL]
	UpdatedAt time.Time      `gorm:"column:updated_at"` // 更新时间
	DeletedBy *uint64        `gorm:"column:deleted_by"` // 删除者 [default NULL]
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"` // 删除时间
}

func (t *User) TableName() string {
	return "user"
}
