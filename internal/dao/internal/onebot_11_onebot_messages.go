// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// Onebot11OnebotMessagesDao is the data access object for the table onebot_messages.
type Onebot11OnebotMessagesDao struct {
	table    string                        // table is the underlying table name of the DAO.
	group    string                        // group is the database configuration group name of the current DAO.
	columns  Onebot11OnebotMessagesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler            // handlers for customized model modification.
}

// Onebot11OnebotMessagesColumns defines and stores column names for the table onebot_messages.
type Onebot11OnebotMessagesColumns struct {
	Id            string //
	CreatedAt     string //
	MessageSeq    string //
	RealSeq       string //
	MessageId     string //
	RealId        string //
	PostType      string //
	EventType     string //
	SubType       string //
	Time          string //
	MessageType   string //
	MessageFormat string //
	Message       string //
	RawMessage    string //
	Font          string //
	SelfId        string //
	UserId        string //
	GroupId       string //
	Sender        string //
	Extra         string //
}

// onebot11OnebotMessagesColumns holds the columns for the table onebot_messages.
var onebot11OnebotMessagesColumns = Onebot11OnebotMessagesColumns{
	Id:            "id",
	CreatedAt:     "created_at",
	MessageSeq:    "message_seq",
	RealSeq:       "real_seq",
	MessageId:     "message_id",
	RealId:        "real_id",
	PostType:      "post_type",
	EventType:     "event_type",
	SubType:       "sub_type",
	Time:          "time",
	MessageType:   "message_type",
	MessageFormat: "message_format",
	Message:       "message",
	RawMessage:    "raw_message",
	Font:          "font",
	SelfId:        "self_id",
	UserId:        "user_id",
	GroupId:       "group_id",
	Sender:        "sender",
	Extra:         "extra",
}

// NewOnebot11OnebotMessagesDao creates and returns a new DAO object for table data access.
func NewOnebot11OnebotMessagesDao(handlers ...gdb.ModelHandler) *Onebot11OnebotMessagesDao {
	return &Onebot11OnebotMessagesDao{
		group:    "onebot11",
		table:    "onebot_messages",
		columns:  onebot11OnebotMessagesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *Onebot11OnebotMessagesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *Onebot11OnebotMessagesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *Onebot11OnebotMessagesDao) Columns() Onebot11OnebotMessagesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *Onebot11OnebotMessagesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *Onebot11OnebotMessagesDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *Onebot11OnebotMessagesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
