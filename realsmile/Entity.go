package realsmile

import (
	"time"
)

type Base struct {
	ID         int64      `gorm:"primaryKey;column:id" json:"id"`
	CreateDate *time.Time `gorm:"comment:创建日期;not null" json:"createDate"`
	Version    int32      `gorm:"comment:版本号;not null;default:1" json:"version"`
}

type Null struct {
}

type Result[T interface{}] struct {
	InnerCode int
	Msg       *string
	Data      *T
	Remark    *string
}

type PaginationResult[T interface{}] struct {
	Result[T]
	total int64
}
