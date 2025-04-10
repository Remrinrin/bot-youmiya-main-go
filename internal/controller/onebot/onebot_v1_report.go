package onebot

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/Remrinrin/bot-youmiya-main-go/api/onebot/v1"
)

func (c *ControllerV1) Report(ctx context.Context, req *v1.ReportReq) (res *v1.ReportRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
