package realsmile

import (
	"database/sql"
	"time"
)

type Base struct {
	ID         int64        `gorm:"primaryKey;column:id" json:"id"`
	CreateDate time.Time    `gorm:"comment:创建日期;not null" json:"createDate"`
	UpdateDate sql.NullTime `gorm:"comment:更新日期" json:"updateDate"`
	Version    int32        `gorm:"comment:版本号;not null;default:1" json:"version"`
}

type Null struct {
}

type Result[T any] struct {
	InnerCode int
	Msg       string
	Data      *T
	Remark    string
}

type PaginationResult[T any] struct {
	Result[T]
	total int64
}
