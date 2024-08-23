package widgets

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gen2brain/dlgs"
	"github.com/gioui-plugins/resource"
	"github.com/gioui-plugins/theme"
)

type FileSelector struct {
	theme        *theme.Theme
	input        *Input
	fileName     string
	actionClick  widget.Clickable
	windowTitle  string
	onSelectFile func(fileName string)
	filter       string
	changed      bool
	width        unit.Dp
	errLog       func(log error)
}

func NewFileSelector(th *theme.Theme, hint string, fileName ...string) *FileSelector {
	bf := &FileSelector{
		theme:       th,
		input:       NewInput(th, hint, fileName...),
		width:       th.Size.DefaultElementWidth,
		windowTitle: "Select file",
		errLog:      func(log error) {},
	}
	if len(fileName) > 0 {
		bf.fileName = fileName[0]
		bf.input.SetText(fileName[0])
	}
	bf.updateIcon()
	return bf
}

func (b *FileSelector) SetFilter(filter string) *FileSelector {
	b.filter = filter
	return b
}

// SetWidth 设置width
func (b *FileSelector) SetWidth(width unit.Dp) *FileSelector {
	b.width = width
	return b
}

// SetWindowTitle 设置windowTitle
func (b *FileSelector) SetWindowTitle(title string) *FileSelector {
	b.windowTitle = title
	return b
}

func (b *FileSelector) SetOnSelectFile(f func(fileName string)) *FileSelector {
	b.onSelectFile = f
	return b
}

func (b *FileSelector) action(gtx layout.Context) {
	if b.actionClick.Clicked(gtx) {
		if b.fileName != "" {
			b.removeFile()
			b.changed = true
			return
		} else {
			file, _, err := dlgs.File(b.windowTitle, b.filter, false)
			if err != nil {
				b.errLog(err)
				return
			}
			if file == "" {
				return
			}
			b.setFileName(file)
			b.changed = true
			if b.onSelectFile != nil {
				b.onSelectFile(file)
			}
		}
	}
}

func (b *FileSelector) setFileName(name string) {
	b.fileName = name
	b.input.SetText(name)
	b.updateIcon()
	b.changed = true
}

func (b *FileSelector) Changed() bool {
	out := b.changed
	b.changed = false
	return out
}

func (b *FileSelector) removeFile() {
	b.fileName = ""
	b.input.SetText("")
	b.updateIcon()
	b.changed = true
}

func (b *FileSelector) GetFileName() string {
	return b.fileName
}

func (b *FileSelector) updateIcon() {
	if b.fileName != "" {
		b.input.SetAfter(func(gtx layout.Context) layout.Dimensions {
			return b.actionClick.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Max.X = gtx.Dp(b.theme.Size.DefaultIconSize)
				return resource.DeleteIcon.Layout(gtx, b.theme.Color.IconGrayColor)
			})
		})
	} else {
		b.input.SetAfter(func(gtx layout.Context) layout.Dimensions {
			return b.actionClick.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Max.X = gtx.Dp(b.theme.Size.DefaultIconSize)
				return resource.UploadIcon.Layout(gtx, b.theme.Color.IconGrayColor)
			})
		})
	}
}

func (b *FileSelector) Layout(gtx layout.Context) layout.Dimensions {
	// gtx.Constraints.Max.Y = gtx.Dp(42)
	b.action(gtx)
	gtx.Constraints.Max.X = gtx.Dp(b.width)
	return b.input.Layout(gtx)
}
