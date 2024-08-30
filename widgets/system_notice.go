package widgets

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"os"
	"path"
	"path/filepath"
)

const (
	icon  = "/Users/lijin/go/ui/qrcodeSymbol.png"
	title = "Client"
)

type SystemNotice struct {
	iconPath string
	message  string
}

func NewSystemNotice() *SystemNotice {
	return &SystemNotice{
		// iconPath: filepath.Join(absolutePath, icon),
		iconPath: icon,
	}
}

func (s *SystemNotice) Notice(message string) error {
	err := beeep.Notify(title, message, s.iconPath)
	if err != nil {
		return err
	}
	return nil
}

func getAbsolutePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("error getting executable path: %s", err.Error())
	}
	exSym, err := filepath.EvalSymlinks(ex)
	if err != nil {
		return "", fmt.Errorf("error getting filepath after evaluating sym links")
	}
	return path.Dir(exSym), nil
}
