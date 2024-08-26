package widgets

import (
	"github.com/x-module/gioui-plugins/theme"
	"image/color"

	"gioui.org/io/semantic"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
)

type IconButton struct {
	theme      *theme.Theme
	icon       *widget.Icon
	size       unit.Dp
	color      color.NRGBA
	colorHover color.NRGBA
	Clickable  widget.Clickable
	onClick    func(gtx layout.Context)
}

func NewIconButton(th *theme.Theme, icon *widget.Icon) *IconButton {
	return &IconButton{
		theme:      th,
		icon:       icon,
		color:      th.Color.IconGrayColor,
		colorHover: th.Color.GreenColor,
		size:       th.Size.DefaultIconSize,
		Clickable:  widget.Clickable{},
	}
}

func (ib *IconButton) SetOnClick(f func(gtx layout.Context)) *IconButton {
	ib.onClick = f
	return ib
}

func (ib *IconButton) SetColor(color color.NRGBA) *IconButton {
	ib.color = color
	return ib
}
func (ib *IconButton) SetColorHover(color color.NRGBA) *IconButton {
	ib.colorHover = color
	return ib
}
func (ib *IconButton) SetSize(size unit.Dp) *IconButton {
	ib.size = size
	return ib
}
func (ib *IconButton) SetIcon(icon *widget.Icon) *IconButton {
	ib.icon = icon
	return ib
}

func (ib *IconButton) Layout(gtx layout.Context) layout.Dimensions {
	for ib.Clickable.Clicked(gtx) {
		if ib.onClick != nil {
			ib.onClick(gtx)
		}
	}
	defaultColor := ib.color
	if ib.Clickable.Hovered() {
		defaultColor = ib.colorHover
	}
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Min.X = gtx.Dp(ib.size)
			return ib.Clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				semantic.Button.Add(gtx.Ops)
				return layout.Background{}.Layout(gtx,
					func(gtx layout.Context) layout.Dimensions {
						return layout.Dimensions{Size: gtx.Constraints.Min}
					},
					func(gtx layout.Context) layout.Dimensions {
						return ib.icon.Layout(gtx, defaultColor)
					},
				)
			})

		}),
	)
}
