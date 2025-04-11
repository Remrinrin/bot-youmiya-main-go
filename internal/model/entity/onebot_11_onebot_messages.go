// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Onebot11OnebotMessages is the golang structure for table onebot11_onebot_messages.
type Onebot11OnebotMessages struct {
	Id            int64       `json:"id"            orm:"id"             description:""` //
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:""` //
	MessageSeq    int64       `json:"messageSeq"    orm:"message_seq"    description:""` //
	RealSeq       int64       `json:"realSeq"       orm:"real_seq"       description:""` //
	MessageId     int64       `json:"messageId"     orm:"message_id"     description:""` //
	RealId        int64       `json:"realId"        orm:"real_id"        description:""` //
	PostType      string      `json:"postType"      orm:"post_type"      description:""` //
	EventType     string      `json:"eventType"     orm:"event_type"     description:""` //
	SubType       string      `json:"subType"       orm:"sub_type"       description:""` //
	Time          int64       `json:"time"          orm:"time"           description:""` //
	MessageType   string      `json:"messageType"   orm:"message_type"   description:""` //
	MessageFormat string      `json:"messageFormat" orm:"message_format" description:""` //
	Message       string      `json:"message"       orm:"message"        description:""` //
	RawMessage    string      `json:"rawMessage"    orm:"raw_message"    description:""` //
	Font          int         `json:"font"          orm:"font"           description:""` //
	SelfId        int64       `json:"selfId"        orm:"self_id"        description:""` //
	UserId        int64       `json:"userId"        orm:"user_id"        description:""` //
	GroupId       int64       `json:"groupId"       orm:"group_id"       description:""` //
	Sender        string      `json:"sender"        orm:"sender"         description:""` //
	Extra         string      `json:"extra"         orm:"extra"          description:""` //
}
