// Code generated by meta-egg. DO NOT EDIT.
// WILL BE replace after re-generated!
// DO NOT EDIT.
// Version: v1.1.0-EE
// Author: meta-egg
// Generated at: 2023-04-27 19:17

package option

import "gorm.io/gorm"

type Option func(tx *gorm.DB) *gorm.DB

// example: Select("id", "name")
// 不和Distinct一起使用
func Select(cols ...interface{}) Option {
	return func(tx *gorm.DB) *gorm.DB {
		if len(cols) > 0 {
			return tx.Select(cols[0], cols[1:]...)
		} else {
			return tx
		}
	}
}

// example: Order("id desc, name asc")
// example: Order("field(id, 45, 44)")
func Order(order string) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Order(order)
	}
}

// example: Limit(10)
func Limit(limit int) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Limit(limit)
	}
}

// example: Offset(10)
func Offset(offset int) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Offset(offset)
	}
}

// 参数定义同gorm的Where
// example: Where("id = ?", 1)
func Where(query interface{}, args ...interface{}) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(query, args...)
	}
}

// 参数定义同gorm的Unscoped, 可忽略软删除
func Unscoped() Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Unscoped()
	}
}

// 参数定义同gorm的Distinct
// example: Distinct("id", "name")
// 不和Select一起使用
func Distinct(cols ...interface{}) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Distinct(cols...)
	}
}

// 参数定义同gorm的Group
// example: Group("name")
// 每个查询只能有一个Group
func Group(group string) Option {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Group(group)
	}
}
