// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package onebot

import (
	"context"

	"github.com/Remrinrin/bot-youmiya-main-go/api/onebot/v1"
)

type IOnebotV1 interface {
	Report(ctx context.Context, req *v1.ReportReq) (res *v1.ReportRes, err error)
}
