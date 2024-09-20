/**
 * Created by Goland
 * @file   markdown.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/9/18 12:05
 * @desc   markdown.go
 */

package widgets

import (
	"fmt"
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/richtext"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/helper/strutil"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	ast2 "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"image/color"
	"strings"
)

// Markdown renders the parsed Markdown to a list of widgets.
type Markdown struct {
	th         *theme.Theme
	widgets    []layout.Widget
	fontStyle  font.Style
	fontWeight font.Weight
	fontColor  color.NRGBA

	source []byte

	htmlTag []string

	taskCheckBox []int // 0 非任务 1 未选中 2 选中
}

// NewMarkdown creates a new Markdown.
func NewMarkdown(theme *theme.Theme) *Markdown {
	return &Markdown{
		th: theme,
	}
}

func (m *Markdown) underLine(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	return NewRichText(m.th).UnderLineLayout(gtx, widget)
}
func (m *Markdown) mark(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	return NewRichText(m.th).MarkLayout(gtx, widget)
}
func (m *Markdown) deleteLine(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	return NewRichText(m.th).DeleteLineLayout(gtx, widget)
}

// Render parses the Markdown content and converts it to a list of widgets.
func (m *Markdown) Render(content []byte) []layout.Widget {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	m.source = content
	document := md.Parser().Parse(text.NewReader(content))
	m.fontWeight = font.Normal
	m.fontStyle = font.Regular
	m.fontColor = m.th.Color.MarkdownDefaultColor

	return m.walk(document, 0, "entrance")
}

func (m *Markdown) filterContent(content string) string {
	return strutil.Replace(content, []string{
		"~~",
	}, []string{""})
}

func (m *Markdown) normal(gtx layout.Context, node any, font font.Font, color color.NRGBA) layout.Dimensions {
	if _, ok := node.(*ast.Text); ok {
		element, ok := node.(*ast.Text)
		if !ok {
			fmt.Println("not text node!!")
			return layout.Dimensions{}
		}
		dims := NewRichText(m.th).AddSpan([]richtext.SpanStyle{
			{
				Content:     string(element.Text(m.source)),
				Size:        unit.Sp(14),
				Interactive: true,
				Color:       color,
				Font:        font,
			},
		}).Layout(gtx)
		m.fontColor = m.th.Color.MarkdownDefaultColor
		return dims
	} else if _, ok = node.(*ast.TextBlock); ok {
		element, ok := node.(*ast.TextBlock)
		if !ok {
			fmt.Println("not text node!!")
			return layout.Dimensions{}
		}
		dims := NewRichText(m.th).AddSpan([]richtext.SpanStyle{
			{
				Content:     string(element.Text(m.source)),
				Size:        unit.Sp(14),
				Interactive: true,
				Color:       color,
				Font:        font,
			},
		}).Layout(gtx)
		m.fontColor = m.th.Color.MarkdownDefaultColor
		return dims
	} else {
		return layout.Dimensions{}
	}

}

func (m *Markdown) getStyleElement(gtx layout.Context, style []string, node any, font font.Font, color color.NRGBA) layout.Dimensions {
	if len(style) == 0 || style[0] == "" {
		return m.normal(gtx, node, font, color)
	} else {
		currentStyle := style[0]
		// 去掉第一个style后剩余的
		otherStyle := style[1:]
		if currentStyle == StyleU { // 下划线
			return m.underLine(gtx, func(gtx layout.Context) layout.Dimensions {
				return m.getStyleElement(gtx, otherStyle, node, font, color)
			})
		} else if currentStyle == StyleS { // 删除线
			return m.deleteLine(gtx, func(gtx layout.Context) layout.Dimensions {
				return m.getStyleElement(gtx, otherStyle, node, font, color)
			})
		} else if currentStyle == StyleMark { // 高亮
			color = m.th.Color.DefaultWindowBgGrayColor
			return m.mark(gtx, func(gtx layout.Context) layout.Dimensions {
				return m.getStyleElement(gtx, otherStyle, node, font, color)
			})
		}
		// else if currentStyle == StyleI { // 斜体
		// 	return m.italic(gtx, child)
		// } else if currentStyle == StyleSmall { // 小字体
		// 	return m.small(gtx, child)
		// } else if currentStyle == StyleBig { // 大字体
		// 	return m.big(gtx, child)
		// } else if currentStyle == StyleBold { // 大字体
		// 	return m.bold(gtx, child)
		// } else { // 普通文本
		// 	return m.normal(gtx, child)
		// }
	}
	return layout.Dimensions{}
}

func intToRoman(num int) string {
	// Define Roman numerals for 1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	// Initialize the result string
	var result string

	// Loop through each value-symbol pair
	for i := 0; i < len(values); i++ {
		// While num is greater than or equal to the value
		for {
			if num >= values[i] {
				// Append the symbol to the result
				result += symbols[i]
				// Subtract the value from num
				num -= values[i]
			} else {
				break
			}
		}
	}

	return result
}
func numToLetter(num int) string {
	// 字母在 ASCII 表中的起始点是 'A' - 1，因为 1 应该对应 'A'
	offset := 'a' - 1
	// 计算给定数字对应的字母
	letter := rune(num) + offset
	// 将 rune 类型转换为 string 并返回
	return string(letter)
}

func getNumber(num int, level int) string {
	switch level {
	case 1:
		return fmt.Sprint(num)
	case 3:
		return intToRoman(num)
	default:
		return numToLetter(num)
	}
}

// walk traverses the AST and converts it to a list of widgets.
func (m *Markdown) walk(node ast.Node, level int, attr string) []layout.Widget {
	var widgets []layout.Widget
	fmt.Println("inter all attr:", attr)
	for child := node.FirstChild(); child != nil; child = child.NextSibling() {
		fmt.Println("inter type:", child.Kind().String(), " attr:", attr)
		switch n := child.(type) {
		case *ast.Text:
			fmt.Println("text all attr:", attr)
			htmlTags := make([]string, len(m.htmlTag))
			copy(htmlTags, m.htmlTag)
			func(font font.Font, color color.NRGBA) {
				widgets = append(widgets, func(gtx layout.Context) layout.Dimensions {
					return m.getStyleElement(gtx, htmlTags, n, font, color)
				})
			}(font.Font{
				Typeface: "go",
				Weight:   m.fontWeight,
				Style:    m.fontStyle,
			}, m.fontColor)
			m.fontWeight = font.Normal
			m.fontStyle = font.Regular
		case *ast.TextBlock:
			widgets = append(widgets, m.walk(n, 0, attr)...)
		case *ast.Heading:
			level := n.Level
			widgets = append(widgets, func(gtx layout.Context) layout.Dimensions {
				heading := material.H1(material.NewTheme(), string(n.Text(m.source)))
				switch level {
				case 2:
					heading = material.H2(material.NewTheme(), string(n.Text(m.source)))
				case 3:
					heading = material.H3(material.NewTheme(), string(n.Text(m.source)))
				case 4:
					heading = material.H4(material.NewTheme(), string(n.Text(m.source)))
				case 5:
					heading = material.H5(material.NewTheme(), string(n.Text(m.source)))
				case 6:
					heading = material.H6(material.NewTheme(), string(n.Text(m.source)))
				}
				return heading.Layout(gtx)
			})
		case *ast.Paragraph:
			var childs []layout.FlexChild
			for _, widget := range m.walk(n, 0, attr) {
				childs = append(childs, layout.Rigid(widget))
			}
			widgets = append(widgets, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, childs...)
			})
		case *ast2.Strikethrough:
			widgets = append(widgets, m.walk(n, 0, attr)...)
		case *ast.List:
			// widgets = m.walk(n)
			index := 1
			lv := level + 1
			listWidgets := m.walk(n, lv, attr)
			fmt.Println("------------list selected status:", m.taskCheckBox, " task count:", len(listWidgets))
			if len(m.taskCheckBox) > 0 {
				for key, item := range listWidgets {
					status := 1
					if key < len(m.taskCheckBox) {
						status = m.taskCheckBox[key]
					}
					// index := i + 1
					item = m.taskListItem(item, status == 2)
					widgets = append(widgets, item)
					index++
				}
			} else {
				for _, item := range listWidgets {
					// index := i + 1
					if n.IsOrdered() {
						item = m.decorateListItem(fmt.Sprintf("%s. ", getNumber(index, lv)), item)
					} else {
						item = m.decorateListItem("• ", item)
					}
					widgets = append(widgets, item)
					index++
				}
			}
		case *ast.ListItem:
			fmt.Println("-------attr:", attr)
			// widgets = append(widgets, m.walk(n)...)
			var childs []layout.FlexChild
			lv := level + 1
			for _, widget := range m.walk(n, lv, attr) {
				childs = append(childs, layout.Rigid(widget))
			}
			widgets = append(widgets, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx, childs...)
			})
		case *ast.Emphasis:
			if n.Level == 1 {
				m.fontStyle = font.Italic
			} else if n.Level == 2 {
				m.fontWeight = font.Bold
			}
			widgets = append(widgets, m.walk(n, 0, attr)...)
		case *ast.HTMLBlock, *ast.RawHTML:
			at := n.(*ast.RawHTML).Segments.At(0)
			tag := string(at.Value(m.source))
			if strings.Contains(tag, "/") {
				m.htmlTag = nil
			} else {
				m.htmlTag = append(m.htmlTag, tag)
			}
		case *ast2.TaskCheckBox:
			fmt.Println("-------------------IsChecked:", n.IsChecked)
			if n.IsChecked {
				m.taskCheckBox = append(m.taskCheckBox, 2)
			} else {
				m.taskCheckBox = append(m.taskCheckBox, 1)
			}
			widgets = append(widgets, m.walk(n, 0, "Task")...)
		case *ast.Image:
			widgets = append(widgets, func(gtx layout.Context) layout.Dimensions {
				return NewImage(m.th, string(n.Destination)).Layout(gtx)
			})
		case *ast2.Table:

			var childs []layout.FlexChild
			for _, widget := range m.walk(n, 0, attr) {
				childs = append(childs, layout.Rigid(widget))
			}
			widgets = append(widgets, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx, childs...)
			})

		case *ast2.TableHeader:
			var childs []layout.FlexChild
			for _, wd := range m.walk(n, 0, attr) {
				childs = append(childs, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return widget.Border{
						Color: m.th.Color.BorderLightGrayColor,
						Width: unit.Dp(1),
					}.Layout(gtx, wd)
				}))
			}
			widgets = append(widgets, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx, childs...)
			})
		case *ast2.TableCell:
			var childs []layout.FlexChild
			for _, wd := range m.walk(n, 0, attr) {
				childs = append(childs, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Min.X = 200
					gtx.Constraints.Min.Y = 60
					return widget.Border{
						Color: m.th.Color.BorderLightGrayColor,
						Width: unit.Dp(1),
					}.Layout(gtx, wd)
				}))
			}
			widgets = append(widgets, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx, childs...)
			})
		case *ast2.TableRow:
			var childs []layout.FlexChild
			for _, wd := range m.walk(n, 0, attr) {
				childs = append(childs, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return widget.Border{
						Color: m.th.Color.BorderLightGrayColor,
						Width: unit.Dp(1),
					}.Layout(gtx, wd)
				}))
			}
			widgets = append(widgets, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx, childs...)
			})
		}
	}
	return widgets
}

// decorateListItem adds bullet points or numbers to list items.
func (m *Markdown) decorateListItem(prefix string, item layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				label := material.Label(m.th.Material(), m.th.Size.DefaultTextSize, prefix)
				label.Color = m.th.Color.DefaultTextWhiteColor
				label.TextSize = m.th.Size.MarkdownPointSize
				return label.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return item(gtx)
			}),
		)
	}
}
func (m *Markdown) taskListItem(item layout.Widget, selected bool) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				box := NewCheckBox(m.th, &widget.Bool{Value: selected}, "")
				return box.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return item(gtx)
			}),
		)
	}
}
