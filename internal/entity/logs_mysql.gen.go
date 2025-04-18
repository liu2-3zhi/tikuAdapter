// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameLog_Mysql = "logs"

//以后能不能把sqlite3和Mysql数据的注册代码分开啊
//autoincrement 是 SQLite 的自增语法
//AUTO_INCREMENT 是 Mysql 的自增语法

// Log mapped from table <logs>
type LogsMysql struct {
	ID         int32     `gorm:"type:integer primary key AUTO_INCREMENT" json:"id"`
	Qid        int32     `gorm:"column:qid;type:int(11);not null;comment:被修改的qid" json:"qid"`                // 被修改的qid
	OldAnswer  string    `gorm:"column:old_answer;type:longtext;not null;comment:修改之前的答案" json:"old_answer"` // 修改之前的答案
	NewAnswer  string    `gorm:"column:new_answer;type:longtext;not null;comment:修改之后的答案" json:"new_answer"` // 修改之后的答案
	UserID     int32     `gorm:"column:user_id;type:int(11);not null;comment:谁修改的答案" json:"user_id"`         // 谁修改的答案
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;comment:修改日期" json:"create_time"`  // 修改日期
	Action     int32     `gorm:"column:action;type:tinyint(4);not null;comment:新增0更新1删除2" json:"action"`     // 新增0更新1删除2
}

// TableName LogsMysql's table name
func (*LogsMysql) TableName() string {
	return TableNameLog_Mysql
}
