package napcat

import (
	"context"

	"github.com/Remrinrin/bot-youmiya-main-go/internal/model/onebot11"
)

type INapcatService interface {
	SendGroupMsg(ctx context.Context, groupID int64, selfID int64, msg *onebot11.MessageArray) error
	// SendPrivateMsg(ctx context.Context, userID string, msg *onebot11.MessageArray) error
}
