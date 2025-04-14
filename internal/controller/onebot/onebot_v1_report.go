package onebot

import (
	"context"
	"encoding/json"
	"time"

	v1 "github.com/Remrinrin/bot-youmiya-main-go/api/onebot/v1"
	"github.com/Remrinrin/bot-youmiya-main-go/internal/logic/onebot"
	"github.com/Remrinrin/bot-youmiya-main-go/internal/model/do"
	"github.com/Remrinrin/bot-youmiya-main-go/internal/model/onebot11"
	fileService "github.com/Remrinrin/bot-youmiya-main-go/internal/service/onebot"
	"github.com/gogf/gf/v2/frame/g"
)

// Report 处理 OneBot 上报
func (c *ControllerV1) Report(ctx context.Context, req *v1.ReportReq) (res *v1.ReportRes, err error) {
	// 解析消息段
	var segments []onebot11.MessageSegment
	if err := json.Unmarshal([]byte(req.Message), &segments); err != nil {
		return nil, err
	}

	// 处理消息段
	for _, segment := range segments {
		if data, ok := segment.Data.(map[string]interface{}); ok {
			// 处理url消息
			if url, ok := data["url"].(string); ok && url != "" {
				// 获取文件名
				filename := ""
				if file, ok := data["file"].(string); ok {
					filename = file
				} else {
					continue
				}

				// 下载并保存图片
				service := fileService.NewFileService(ctx)
				_, err := service.DownloadAndSave(ctx, url, filename)
				if err != nil {
					g.Log().Error(ctx, "下载并保存图片失败", err)
				}
			}

			// 处理path消息
			if path, ok := data["path"].(string); ok && path != "" {
				// 获取文件名
				filename := ""
				if file, ok := data["file"].(string); ok {
					filename = file
				}

				// 复制文件
				service := fileService.NewFileService(ctx)
				_, err := service.CopyFromPath(ctx, path, filename)
				if err != nil {
					g.Log().Error(ctx, "复制文件失败", err)
				}
			}
		}
	}

	// 更新消息内容
	messageBytes, err := json.Marshal(segments)
	if err != nil {
		return nil, err
	}
	req.Message = string(messageBytes)

	// 保存消息到数据库
	err = saveMessage(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.ReportRes{
		Code: 0,
		Msg:  "success",
		Data: nil,
	}, nil
}

// saveMessage 保存消息到数据库
func saveMessage(ctx context.Context, req *v1.ReportReq) error {
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
	_, err := onebotService.InsertMsg(ctx, msg)
	return err
}

func handleUrlFileStorage(ctx context.Context, url string, filename string) error {
	service := fileService.NewFileService(ctx)
	_, err := service.DownloadAndSave(ctx, url, filename)
	if err != nil {
		return err
	}
	return nil
}

func handlePathFileStorage(ctx context.Context, path string, filename string) error {
	service := fileService.NewFileService(ctx)
	_, err := service.CopyFromPath(ctx, path, filename)
	if err != nil {
		return err
	}
	return nil
}
