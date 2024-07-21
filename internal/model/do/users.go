// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta         `orm:"table:users, do:true"`
	Id             interface{} //
	Username       interface{} //
	Password       interface{} //
	Email          interface{} //
	WordsMemorized interface{} //
	WordsReviewed  interface{} //
}
