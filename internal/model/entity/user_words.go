// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserWords is the golang structure for table user_words.
type UserWords struct {
	UserId       int         `json:"userId"       ` //
	WordId       int         `json:"wordId"       ` //
	Status       int         `json:"status"       ` //
	LastReviewed *gtime.Time `json:"lastReviewed" ` //
}
