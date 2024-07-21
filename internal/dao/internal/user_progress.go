// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserProgressDao is the data access object for table user_progress.
type UserProgressDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns UserProgressColumns // columns contains all the column names of Table for convenient usage.
}

// UserProgressColumns defines and stores column names for table user_progress.
type UserProgressColumns struct {
	UserId        string //
	LastWordId    string //
	WordsToLearn  string //
	WordsToReview string //
}

// userProgressColumns holds the columns for table user_progress.
var userProgressColumns = UserProgressColumns{
	UserId:        "user_id",
	LastWordId:    "last_word_id",
	WordsToLearn:  "words_to_learn",
	WordsToReview: "words_to_review",
}

// NewUserProgressDao creates and returns a new DAO object for table data access.
func NewUserProgressDao() *UserProgressDao {
	return &UserProgressDao{
		group:   "default",
		table:   "user_progress",
		columns: userProgressColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserProgressDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserProgressDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserProgressDao) Columns() UserProgressColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserProgressDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserProgressDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserProgressDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
