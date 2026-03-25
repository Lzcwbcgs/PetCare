package main

import (
	_ "PetCare/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"PetCare/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
