package onebot

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gogf/gf/v2/frame/g"
)

type IFileService interface {
	// DownloadAndSave 下载文件并保存到指定目录
	DownloadAndSave(ctx context.Context, url string, filename string) (string, error)
	// CopyFromPath 从本地路径复制文件
	CopyFromPath(ctx context.Context, srcPath string, filename string) (string, error)
}

type FileService struct {
	ctx context.Context
}

func NewFileService(ctx context.Context) IFileService {
	return &FileService{ctx}
}

func (s *FileService) DownloadAndSave(ctx context.Context, url string, filename string) (string, error) {
	// 获取存储目录配置
	baseDir := g.Cfg().MustGet(ctx, "storage.baseDir").String()
	if baseDir == "" {
		return "", fmt.Errorf("未配置存储目录")
	}

	// 生成文件路径
	filePath := filepath.Join(baseDir, filename)

	// 使用 curl 下载文件
	cmd := exec.Command("curl", "-L", "-o", filePath, url)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("下载文件失败: %v, 输出: %s", err, string(output))
	}

	return filePath, nil
}

func (s *FileService) CopyFromPath(ctx context.Context, srcPath string, filename string) (string, error) {
	// 获取存储目录配置
	baseDir := g.Cfg().MustGet(ctx, "storage.baseDir").String()
	if baseDir == "" {
		return "", fmt.Errorf("未配置存储目录")
	}

	// 生成目标文件路径
	dstPath := filepath.Join(baseDir, filename)

	// 检查源文件
	fileInfo, err := os.Lstat(srcPath)
	if err != nil {
		if os.IsPermission(err) {
			return "", fmt.Errorf("没有权限访问源文件: %s, 错误: %v", srcPath, err)
		}
		if os.IsNotExist(err) {
			return "", fmt.Errorf("源文件不存在: %s, 错误: %v", srcPath, err)
		}
		return "", fmt.Errorf("检查源文件失败: %s, 错误: %v", srcPath, err)
	}

	// 检查是否是符号链接
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		// 如果是符号链接，获取实际路径
		actualPath, err := os.Readlink(srcPath)
		if err != nil {
			return "", fmt.Errorf("读取符号链接失败: %s, 错误: %v", srcPath, err)
		}
		srcPath = actualPath
	}

	// 打开源文件
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return "", fmt.Errorf("打开源文件失败: %s, 错误: %v", srcPath, err)
	}
	defer srcFile.Close()

	// 创建目标文件
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return "", fmt.Errorf("创建目标文件失败: %s, 错误: %v", dstPath, err)
	}
	defer dstFile.Close()

	// 复制文件内容
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return "", fmt.Errorf("复制文件失败: %s -> %s, 错误: %v", srcPath, dstPath, err)
	}

	return dstPath, nil
}
