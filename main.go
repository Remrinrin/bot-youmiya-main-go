package main

import (
	_ "github.com/Remrinrin/bot-youmiya-main-go/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/Remrinrin/bot-youmiya-main-go/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
