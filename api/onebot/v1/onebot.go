package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

const (
	MsgTypePrivate = "private"
	MsgTypeGroup   = "group"
)

// ReportReq OneBot 上报请求
type ReportReq struct {
	g.Meta `path:"/onebot/report" tags:"OneBot" method:"post" summary:"OneBot 事件上报"`

	// 消息序列相关字段
	MessageSeq int64 `json:"message_seq" dc:"消息序列号"` // 消息序列号
	RealSeq    int64 `json:"real_seq" dc:"真实序列号"`    // 真实序列号
	MessageID  int64 `json:"message_id" dc:"消息 ID"`  // 消息 ID
	RealID     int64 `json:"real_id" dc:"真实 ID"`     // 真实 ID

	// 事件类型相关字段
	PostType  string `json:"post_type" dc:"上报类型"`  // 上报类型
	Time      int64  `json:"time" dc:"事件发生的时间戳"`   // 事件发生的时间戳
	EventType string `json:"event_type" dc:"事件类型"` // 事件类型
	SubType   string `json:"sub_type" dc:"子类型"`    // 子类型

	// 消息内容相关字段
	MessageType   string `json:"message_type" dc:"消息类型"`   // 消息类型
	MessageFormat string `json:"message_format" dc:"消息格式"` // 消息格式
	Message       string `json:"message" dc:"消息内容"`        // 消息内容
	RawMessage    string `json:"raw_message" dc:"原始消息"`    // 原始消息
	Font          int    `json:"font" dc:"字体"`             // 字体

	// 参与者信息字段
	SelfID  int64                  `json:"self_id" dc:"机器人 QQ 号"` // 机器人 QQ 号
	UserID  int64                  `json:"user_id" dc:"发送者 QQ 号"` // 发送者 QQ 号
	GroupID int64                  `json:"group_id" dc:"群号"`      // 群号
	Sender  map[string]interface{} `json:"sender" dc:"发送者信息"`     // 发送者信息

	// 扩展信息字段
	Extra map[string]interface{} `json:"extra" dc:"额外信息"` // 额外信息
}

// ReportRes OneBot 上报响应
type ReportRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Code   int         `json:"code" dc:"状态码"`
	Msg    string      `json:"msg" dc:"状态信息"`
	Data   interface{} `json:"data" dc:"响应数据"`
}
