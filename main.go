package main

import (
	_ "gf-shop/internal/packed"

	_ "gf-shop/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-shop/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
