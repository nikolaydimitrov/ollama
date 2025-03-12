package tray

import (
	"github.com/nikolaydimitrov/ollama/app/tray/commontray"
	"github.com/nikolaydimitrov/ollamaov/ollama/app/tray/wintray"
)

func InitPlatformTray(icon, updateIcon []byte) (commontray.OllamaTray, error) {
	return wintray.InitTray(icon, updateIcon)
}
