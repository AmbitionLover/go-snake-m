package main

import (
	_ "go-snake/app/api/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"go-snake/app/api/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
