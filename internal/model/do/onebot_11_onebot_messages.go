// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Onebot11OnebotMessages is the golang structure of table onebot_messages for DAO operations like Where/Data.
type Onebot11OnebotMessages struct {
	g.Meta        `orm:"table:onebot_messages, do:true"`
	Id            interface{} //
	CreatedAt     *gtime.Time //
	MessageSeq    interface{} //
	RealSeq       interface{} //
	MessageId     interface{} //
	RealId        interface{} //
	PostType      interface{} //
	EventType     interface{} //
	SubType       interface{} //
	Time          interface{} //
	MessageType   interface{} //
	MessageFormat interface{} //
	Message       interface{} //
	RawMessage    interface{} //
	Font          interface{} //
	SelfId        interface{} //
	UserId        interface{} //
	GroupId       interface{} //
	Sender        interface{} //
	Extra         interface{} //
}
