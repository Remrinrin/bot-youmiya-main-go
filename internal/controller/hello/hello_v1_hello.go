package hello

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"

	v1 "github.com/Remrinrin/bot-youmiya-main-go/api/hello/v1"
)

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	request := g.RequestFromCtx(ctx)

	fmt.Println("=== 请求信息 ===")
	fmt.Printf("请求方法: %s\n", request.Method)
	fmt.Printf("请求URL: %s\n", request.URL.String())

	fmt.Println("=== 请求头 ===")
	for k, v := range request.Header {
		fmt.Printf("%s: %v\n", k, v)
	}

	fmt.Println("=== 请求体 ===")
	fmt.Println(request.GetBody())

	fmt.Println("=== 查询参数 ===")
	queryMap := request.GetQueryMap()
	for k, v := range queryMap {
		fmt.Printf("%s: %v\n", k, v)
	}

	fmt.Println("=== 表单数据 ===")
	formMap := request.GetFormMap()
	for k, v := range formMap {
		fmt.Printf("%s: %v\n", k, v)
	}

	fmt.Println("=== 请求结束 ===")

	return
}
