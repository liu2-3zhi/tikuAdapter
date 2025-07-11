// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID       int32  `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"column:username;type:varchar(255)" json:"username"`
	Password string `gorm:"column:password;type:varchar(255)" json:"password"`
	Perms    string `gorm:"column:perms;type:varchar(255)" json:"perms"`
	ParentID int32  `gorm:"column:parent_id;type:int(11)" json:"parent_id"`
	Nickname string `gorm:"column:nickname;type:varchar(255)" json:"nickname"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
