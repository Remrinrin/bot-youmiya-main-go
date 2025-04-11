package onebot

import (
	"context"

	"github.com/Remrinrin/bot-youmiya-main-go/internal/dao"
	"github.com/Remrinrin/bot-youmiya-main-go/internal/model/do"
	"github.com/Remrinrin/bot-youmiya-main-go/internal/service/onebot"
)

type MsgServiceLogic struct {
	ctx context.Context
}

func NewMsgService(ctx context.Context) onebot.IMsgService {
	return &MsgServiceLogic{ctx}
}

func (l *MsgServiceLogic) InsertMsg(ctx context.Context, msg *do.Onebot11OnebotMessages) (Id int64, err error) {
	Id, err = dao.Onebot11OnebotMessages.Ctx(ctx).InsertAndGetId(msg)
	if err != nil {
		return 0, err
	}
	return Id, nil
}
