package onebot

import (
	"context"
	"time"

	v1 "github.com/Remrinrin/bot-youmiya-main-go/api/onebot/v1"
	"github.com/Remrinrin/bot-youmiya-main-go/internal/logic/onebot"
	"github.com/Remrinrin/bot-youmiya-main-go/internal/model/do"
)

func (c *ControllerV1) Report(ctx context.Context, req *v1.ReportReq) (res *v1.ReportRes, err error) {
	msg := &do.Onebot11OnebotMessages{
		MessageSeq:    req.MessageSeq,
		RealSeq:       req.RealSeq,
		MessageId:     req.MessageID,
		RealId:        req.RealID,
		PostType:      req.PostType,
		EventType:     req.EventType,
		SubType:       req.SubType,
		Time:          time.Now().Unix(),
		MessageType:   req.MessageType,
		MessageFormat: req.MessageFormat,
		Message:       req.Message,
		RawMessage:    req.RawMessage,
		Font:          req.Font,
		SelfId:        req.SelfID,
		UserId:        req.UserID,
		GroupId:       req.GroupID,
		Sender:        req.Sender,
		Extra:         req.Extra,
	}
	onebotService := onebot.NewMsgService(ctx)
	_, err = onebotService.InsertMsg(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &v1.ReportRes{
		Code: 0,
		Msg:  "success",
		Data: "success",
	}, nil
}
