package middleware

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// TokenMiddleware 校验 token 的中间件
func TokenMiddleware(r *ghttp.Request) {
	g.Log().Info(r.Context(), "=== Token 验证开始 ===")
	g.Log().Infof(r.Context(), "请求路径: %s", r.URL.Path)
	g.Log().Infof(r.Context(), "请求方法: %s", r.Method)

	// 获取请求体
	body := r.GetBody()
	if len(body) == 0 {
		g.Log().Error(r.Context(), "错误: 请求体为空")
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    401,
			Message: "请求体为空",
		})
		r.Exit()
		return
	}
	g.Log().Infof(r.Context(), "请求体长度: %d bytes", len(body))

	// 从配置文件获取 header 名称
	signatureHeader := g.Cfg().MustGet(r.Context(), "onebot11.signature_header").String()
	// selfIDHeader := g.Cfg().MustGet(r.Context(), "onebot.self_id_header", "x-self-id").String()

	// 获取 signature header
	signature := r.Header.Get(signatureHeader)
	if signature == "" {
		g.Log().Errorf(r.Context(), "错误: 缺少 %s header", signatureHeader)
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    401,
			Message: "缺少 " + signatureHeader + " header",
		})
		r.Exit()
		return
	}
	g.Log().Infof(r.Context(), "收到签名: %s", signature)

	// 从配置文件获取 token
	token := g.Cfg().MustGet(r.Context(), "onebot11.signature_token").String()
	if token == "" {
		g.Log().Error(r.Context(), "错误: 服务器未配置 token")
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    500,
			Message: "服务器未配置 token",
		})
		r.Exit()
		return
	}

	// 计算 HMAC-SHA1
	h := hmac.New(sha1.New, []byte(token))
	h.Write(body)
	expectedSignature := "sha1=" + hex.EncodeToString(h.Sum(nil))
	g.Log().Infof(r.Context(), "计算签名: %s", expectedSignature)

	// 验证签名
	if !strings.EqualFold(signature, expectedSignature) {
		g.Log().Error(r.Context(), "错误: 签名验证失败")
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    401,
			Message: "签名验证失败",
		})
		r.Exit()
		return
	}

	g.Log().Info(r.Context(), "签名验证通过")
	g.Log().Info(r.Context(), "=== Token 验证结束 ===")

	// 验证通过，继续处理请求
	r.Middleware.Next()
}
