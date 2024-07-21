// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserWords is the golang structure of table user_words for DAO operations like Where/Data.
type UserWords struct {
	g.Meta       `orm:"table:user_words, do:true"`
	UserId       interface{} //
	WordId       interface{} //
	Status       interface{} //
	LastReviewed *gtime.Time //
}
