package onebot

import (
	"context"

	"github.com/Remrinrin/bot-youmiya-main-go/internal/model/do"
)

type IMsgService interface {
	// 消息
	InsertMsg(ctx context.Context, msg *do.Onebot11OnebotMessages) (Id int64, err error)
}
