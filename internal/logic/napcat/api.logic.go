package napcat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Remrinrin/bot-youmiya-main-go/internal/model/onebot11"
	"github.com/Remrinrin/bot-youmiya-main-go/internal/service/napcat"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type NapcatServiceLogic struct {
	ctx context.Context
}

func NewNapcatService(ctx context.Context) napcat.INapcatService {
	return &NapcatServiceLogic{ctx}
}

func (l *NapcatServiceLogic) SendGroupMsg(ctx context.Context, groupID int64, selfID int64, msg *onebot11.MessageArray) error {
	// 获取所有 napcat 配置
	configs := g.Cfg().MustGet(ctx, "napcat").Array()
	if len(configs) == 0 {
		return fmt.Errorf("未找到 napcat 配置")
	}

	// 遍历配置，找到匹配的 self_id
	var targetConfig map[string]interface{}
	for _, config := range configs {
		cfg := config.(map[string]interface{})
		// 打印配置信息以便调试
		g.Log().Debugf(ctx, "配置信息: %+v", cfg)

		// 尝试获取 self_id
		cfgSelfID := cfg["self_id"]
		if cfgSelfID == nil {
			return fmt.Errorf("配置中缺少 self_id")
		}

		if gconv.Int64(cfgSelfID) == selfID {
			targetConfig = cfg
			break
		}
	}

	if targetConfig == nil {
		return fmt.Errorf("未找到匹配的 self_id: %d", selfID)
	}

	// 获取 HTTP URL 和 token
	httpURL := targetConfig["http_url"].(string)
	token := targetConfig["token"].(string)

	// 构造请求 URL
	url := strings.TrimRight(httpURL, "/") + "/send_group_msg"

	// 构造请求体
	reqBody := map[string]interface{}{
		"group_id": groupID,
		"message":  msg,
	}

	// 创建 HTTP 客户端
	client := &http.Client{}

	// 创建请求
	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// 设置请求体
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("序列化请求体失败: %v", err)
	}
	req.Body = io.NopCloser(io.NewSectionReader(bytes.NewReader(jsonData), 0, int64(len(jsonData))))

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("发送消息失败: %s", resp.Status)
	}

	return nil
}
