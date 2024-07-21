// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserProgress is the golang structure of table user_progress for DAO operations like Where/Data.
type UserProgress struct {
	g.Meta        `orm:"table:user_progress, do:true"`
	UserId        interface{} //
	LastWordId    interface{} //
	WordsToLearn  interface{} //
	WordsToReview interface{} //
}
