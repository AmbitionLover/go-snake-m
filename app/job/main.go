package main

import (
	_ "go-snake/app/job/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"go-snake/app/job/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
