package main

import (
	_ "go-snake/app/svc/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"go-snake/app/svc/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
