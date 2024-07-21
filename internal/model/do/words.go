// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Words is the golang structure of table words for DAO operations like Where/Data.
type Words struct {
	g.Meta      `orm:"table:words, do:true"`
	Id          interface{} //
	Word        interface{} //
	Translation interface{} //
	Difficulty  interface{} //
}
