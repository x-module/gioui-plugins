package widgets

import (
	"gioui.org/font"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/x/component"
	"github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
	"image"
	"image/color"
)

type SearchDropDown struct {
	SearchDropDown  *widget.Bool
	menuContextArea component.ContextArea
	menu            component.MenuState
	list            *widget.List
	theme           *theme.Theme

	minWidth unit.Dp
	width    unit.Dp
	menuInit bool

	isOpen              bool
	selectedOptionIndex int
	lastSelectedIndex   int
	options             []*SearchDropDownOption

	size image.Point

	borderWidth  unit.Dp
	cornerRadius unit.Dp

	onValueChange func(value string)

	searchInput *Input
}

type SearchDropDownOption struct {
	Text       string
	Value      string
	Identifier string
	clickable  widget.Clickable

	Icon      *widget.Icon
	IconColor color.NRGBA

	isDivider bool
	isDefault bool
}

// SetWidth 设置width
func (c *SearchDropDown) SetWidth(width unit.Dp) {
	c.width = width
}

func NewSearchDropDownOption(text string) *SearchDropDownOption {
	return &SearchDropDownOption{
		Text:      text,
		Value:     text,
		isDivider: false,
	}
}

func NewSearchDropDownDivider() *SearchDropDownOption {
	return &SearchDropDownOption{
		isDivider: true,
	}
}

func (o *SearchDropDownOption) WithIdentifier(identifier string) *SearchDropDownOption {
	o.Identifier = identifier
	return o
}

func (o *SearchDropDownOption) WithValue(value string) *SearchDropDownOption {
	o.Value = value
	return o
}

func (o *SearchDropDownOption) WithIcon(icon *widget.Icon, color color.NRGBA) *SearchDropDownOption {
	o.Icon = icon
	o.IconColor = color
	return o
}

func (o *SearchDropDownOption) DefaultSelected() *SearchDropDownOption {
	o.isDefault = true
	return o
}

func (o *SearchDropDownOption) GetText() string {
	if o == nil {
		return ""
	}

	return o.Text
}

func (o *SearchDropDownOption) GetValue() string {
	if o == nil {
		return ""
	}

	return o.Value
}

func (c *SearchDropDown) SetSelected(index int) {
	c.selectedOptionIndex = index
	c.lastSelectedIndex = index
}

func (c *SearchDropDown) SetOnChanged(f func(value string)) {
	c.onValueChange = f
}

func (c *SearchDropDown) SetSelectedByTitle(title string) {
	if len(c.options) == 0 {
		return
	}

	for i, opt := range c.options {
		if opt.Text == title {
			c.selectedOptionIndex = i
			c.lastSelectedIndex = i
			break
		}
	}
}

func (c *SearchDropDown) SetSelectedByIdentifier(identifier string) {
	for i, opt := range c.options {
		if opt.Identifier == identifier {
			c.selectedOptionIndex = i
			c.lastSelectedIndex = i
			break
		}
	}
}

func (c *SearchDropDown) SetSelectedByValue(value string) {
	for i, opt := range c.options {
		if opt.Value == value {
			c.selectedOptionIndex = i
			c.lastSelectedIndex = i
			break
		}
	}
}

func NewSearchDropDown(th *theme.Theme, options ...string) *SearchDropDown {
	c := &SearchDropDown{
		SearchDropDown: &widget.Bool{Value: true},
		menuContextArea: component.ContextArea{
			Activation:       pointer.ButtonPrimary,
			AbsolutePosition: true,
		},
		list: &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
		borderWidth:  unit.Dp(1),
		cornerRadius: unit.Dp(4),
		theme:        th,
		menuInit:     true,
		// width:        th.Size.DefaultElementWidth,
		searchInput: NewInput(th, "Search..."),
	}
	if len(options) > 0 {
		for _, opt := range options {
			c.options = append(c.options, NewSearchDropDownOption(opt))
		}
	}
	// c.searchInput.SetBefore(func(gtx layout.Context) layout.Dimensions {
	// 	return resource.SearchIcon.Layout(gtx, th.Color.DefaultIconColor)
	// })
	c.searchInput.SetIcon(resource.SearchIcon)
	size := theme.ElementStyle{TextSize: unit.Sp(14), Height: unit.Dp(17), Inset: layout.UniformInset(unit.Dp(5)), IconSize: unit.Dp(18)}
	c.searchInput.SetSize(size)
	return c
}

func NewSearchDropDownWithoutBorder(th *theme.Theme, options ...string) *SearchDropDown {
	c := &SearchDropDown{
		SearchDropDown: &widget.Bool{Value: true},
		menuContextArea: component.ContextArea{
			Activation:       pointer.ButtonPrimary,
			AbsolutePosition: true,
		},
		list: &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
		borderWidth:  unit.Dp(0),
		cornerRadius: unit.Dp(4),
		theme:        th,
		menuInit:     true,
		// width:        th.Size.DefaultElementWidth,
		searchInput: NewInput(th, "Search..."),
	}
	if len(options) > 0 {
		for _, opt := range options {
			c.options = append(c.options, NewSearchDropDownOption(opt))
		}
	}
	return c
}

func (c *SearchDropDown) SelectedIndex() int {
	return c.selectedOptionIndex
}

func (c *SearchDropDown) SetOptions(options []*SearchDropDownOption) {
	c.selectedOptionIndex = 0
	c.options = options
	if len(c.options) > 0 {
		c.menuInit = true
	}
}

func (c *SearchDropDown) GetSelected() *SearchDropDownOption {
	if len(c.options) == 0 {
		return nil
	}

	return c.options[c.selectedOptionIndex]
}

func (c *SearchDropDown) box1(gtx layout.Context, theme *theme.Theme, text string, maxWidth unit.Dp) layout.Dimensions {
	borderColor := c.theme.Color.DropDownBorderColor
	if c.isOpen {
		borderColor = c.theme.Color.DropDownHoveredBorderColor
	}
	border := widget.Border{
		Color:        borderColor,
		Width:        c.borderWidth,
		CornerRadius: c.cornerRadius,
	}

	if maxWidth == 0 {
		maxWidth = unit.Dp(gtx.Constraints.Max.X)
	}

	c.size.X = gtx.Dp(maxWidth)
	return c.SearchDropDown.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		if c.SearchDropDown.Hovered() {
			border.Color = c.theme.Color.DropDownHoveredBorderColor
		}
		return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			// calculate the minimum width of the box, considering icon and padding
			return layout.Background{}.Layout(gtx,
				func(gtx layout.Context) layout.Dimensions {
					defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, gtx.Dp(c.theme.Size.DefaultElementRadiusSize)).Push(gtx.Ops).Pop()
					paint.Fill(gtx.Ops, c.theme.Color.DefaultBgGrayColor)
					return layout.Dimensions{Size: gtx.Constraints.Min}
				},
				func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Min.X = gtx.Dp(maxWidth) - gtx.Dp(8)
					return layout.Inset{Top: unit.Dp(3), Bottom: unit.Dp(3)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						c.searchInput.SetWidth(c.width).SetHasBorder(false)
						return c.searchInput.Layout(gtx)
					})
				},
			)
		})
	})
}
func (c *SearchDropDown) box(gtx layout.Context, theme *theme.Theme, text string, maxWidth unit.Dp) layout.Dimensions {
	c.searchInput.SetWidth(c.width).SetHasBorder(false)
	return c.searchInput.Layout(gtx)
}

func (c *SearchDropDown) SetSize(size image.Point) {
	c.size = size
}

func (c *SearchDropDown) SetMinWidth(minWidth unit.Dp) {
	c.minWidth = minWidth
}

// Layout the SearchDropDown.
func (c *SearchDropDown) Layout(gtx layout.Context, theme *theme.Theme) layout.Dimensions {
	c.isOpen = c.menuContextArea.Active()

	for i, opt := range c.options {
		if opt.isDefault {
			c.selectedOptionIndex = i
		}

		for opt.clickable.Clicked(gtx) {
			c.isOpen = false
			c.selectedOptionIndex = i
		}
	}

	if c.selectedOptionIndex != c.lastSelectedIndex {
		if c.onValueChange != nil {
			go c.onValueChange(c.options[c.selectedOptionIndex].Value)
		}
		c.lastSelectedIndex = c.selectedOptionIndex
	}

	// Update menu items only if options change
	if c.menuInit {
		c.menuInit = false
		c.updateMenuItems(theme)
	}

	if c.minWidth == 0 {
		c.minWidth = unit.Dp(150)
	}

	text := ""
	if c.selectedOptionIndex >= 0 && c.selectedOptionIndex < len(c.options) {
		text = c.options[c.selectedOptionIndex].Text
	}
	box := c.box1(gtx, theme, text, c.width)
	return layout.Stack{}.Layout(gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			// c.searchInput.SetWidth(c.width).SetHasBorder(false)
			// return c.searchInput.Layout(gtx)
			return box
		}),
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return c.menuContextArea.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				offset := layout.Inset{
					Top:  unit.Dp(float32(box.Size.Y)/gtx.Metric.PxPerDp + 1),
					Left: unit.Dp(0),
				}
				return offset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Min.X = gtx.Dp(c.minWidth)
					if c.width != 0 {
						gtx.Constraints.Max.X = gtx.Dp(c.width)
					}
					m := component.Menu(theme.Material(), &c.menu)
					m.SurfaceStyle.Fill = c.theme.Color.DropDownBgGrayColor
					return m.Layout(gtx)
				})
			})
		}),
	)
}

// updateMenuItems creates or updates menu items based on options and calculates minWidth.
func (c *SearchDropDown) updateMenuItems(th *theme.Theme) {
	c.menu.Options = c.menu.Options[:0]
	for _, opt := range c.options {
		opt := opt
		c.menu.Options = append(c.menu.Options, func(gtx layout.Context) layout.Dimensions {
			if opt.isDivider {
				dv := component.Divider(th.Material())
				dv.Fill = c.theme.Color.DefaultBorderGrayColor
				dv.Layout(gtx)
			}
			itm := component.MenuItem(th.Material(), &opt.clickable, opt.Text)
			itm.HoverColor = c.theme.Color.DropDownItemHoveredGrayColor
			if opt.Icon != nil {
				itm.Icon = opt.Icon
				itm.IconColor = opt.IconColor
			}
			itm.Label.TextSize = c.theme.Size.DropdownTextSize
			if c.GetSelected().Text == opt.Text {
				itm.Label.Color = c.theme.Color.DropDownSelectedItemBgColor
				itm.Label.Font.Weight = font.Bold
				// itm.Icon = widgets.ActionStarRateIcon
				// itm.IconSize = unit.Dp(16)
				// itm.IconInset = outlay.Inset{}
				// itm.IconColor = opt.IconColor
			} else {
				itm.Label.Color = c.theme.Color.DefaultTextWhiteColor
			}
			return itm.Layout(gtx)
		})
	}
}
