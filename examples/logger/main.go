package main

import (
	"fmt"
	"github.com/gozelle/cobra"
	"github.com/gozelle/cobra/progress"
	"time"
)

func main() {
	logger := cobra.NewLogger()
	for i := 0; i < 100; i++ {
		logger.Progress(
			progress.Value("上传进度"),
			progress.Value(fmt.Sprintf("%d%%", i)).WithWidth(3),
		)
		time.Sleep(10 * time.Millisecond)
	}
	logger.NewLine()
	
	logger.Success("success")
	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
	
	cobra.Success("🚀", "success")
	cobra.Debug("🚀", "debug")
	cobra.Info("🚀", "info")
	cobra.Warn("🚀", "warn")
	cobra.Error("🚀", "error")
	
	logger = logger.WithModule("begin")
	logger.Success("😊", "success")
	logger.Debug("😊", "debug")
	logger.Info("😊", "info")
	logger.Warn("😊", "warn")
	logger.Error("😊", "error")
	logger.Fatal("😊", "fatal")
	
}
