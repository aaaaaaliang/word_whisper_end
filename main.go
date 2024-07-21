package main

import (
	_ "word_whisper_end/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "word_whisper_end/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"word_whisper_end/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
