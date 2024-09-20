/**
 * Created by Goland
 * @file   Markdown1.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/9/18 12:05
 * @desc   Markdown1.go
 */

package widgets

import (
	"fmt"
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"gioui.org/x/richtext"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"image/color"
	"strings"
)

// Markdown1 renders the parsed Markdown1 to a list of widgets.
type Markdown1 struct {
	th     *theme.Theme
	source []byte

	widgets []layout.Widget

	fontStyle  font.Style
	fontWeight font.Weight
	fontColor  color.NRGBA
	fontSize   unit.Sp

	htmlTag []string
}

// NewMarkdown1 creates a new Markdown1.
func NewMarkdown1(theme *theme.Theme) *Markdown1 {
	return &Markdown1{
		th: theme,
	}
}

// Render parses the Markdown1 content and converts it to a list of widgets.
func (m *Markdown1) Render(content []byte) []layout.Widget {
	md := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)
	m.source = content
	document := md.Parser().Parse(text.NewReader(content))
	return m.walk(document)
}

const (
	StyleU     = "<u>"    // 下划线
	StyleS     = "<s>"    // 删除线
	StyleMark  = "<mark>" // 高亮
	StyleSmall = "small"  // 小字体
	StyleBig   = "big"    // 大字体
	StyleBold  = "bold"   // 大字体
)

func (m *Markdown1) normal(gtx layout.Context, node any) layout.Dimensions {
	fmt.Println("tags:", m.htmlTag)
	element, ok := node.(*ast.Text)
	if !ok {
		fmt.Println("not text node!!")
		return layout.Dimensions{}
	}
	fontWeight := m.fontWeight
	fontStyle := m.fontStyle
	fontColor := m.fontColor

	dims := NewRichText(m.th).AddSpan([]richtext.SpanStyle{
		{
			Content:     string(element.Text(m.source)),
			Size:        unit.Sp(14),
			Interactive: true,
			Color:       fontColor,
			Font: font.Font{
				Typeface: "go",
				Weight:   fontWeight,
				Style:    fontStyle,
			},
		},
	}).Layout(gtx)
	m.fontWeight = font.Normal
	m.fontStyle = font.Regular
	m.fontColor = m.th.Color.MarkdownDefaultColor
	return dims
}
func (m *Markdown1) getStyleElement(gtx layout.Context, style []string, node any) layout.Dimensions {
	fmt.Println("-------------------------------------------------------------------")
	fmt.Printf("all style:%s\n", style)
	if len(style) == 0 {
		return m.normal(gtx, node)
	} else {
		currentStyle := style[0]
		// 去掉第一个style后剩余的
		otherStyle := style[1:]
		fmt.Println("current style:", currentStyle)
		fmt.Printf("other style:%s\n", otherStyle)
		if currentStyle == StyleU { // 下划线
			fmt.Println("下划线")
			return m.underLine(gtx, func(gtx layout.Context) layout.Dimensions {
				return m.getStyleElement(gtx, otherStyle, node)
			})
		} else if currentStyle == StyleS { // 删除线
			return m.deleteLine(gtx, func(gtx layout.Context) layout.Dimensions {
				return m.getStyleElement(gtx, otherStyle, node)
			})
		} else if currentStyle == StyleMark { // 高亮
			m.fontColor = m.th.Color.DefaultWindowBgGrayColor
			return m.mark(gtx, func(gtx layout.Context) layout.Dimensions {
				return m.getStyleElement(gtx, otherStyle, node)
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

func (m *Markdown1) underLine(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	return NewRichText(m.th).UnderLineLayout(gtx, widget)
}
func (m *Markdown1) mark(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	return NewRichText(m.th).MarkLayout(gtx, widget)
}
func (m *Markdown1) deleteLine(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	return NewRichText(m.th).DeleteLineLayout(gtx, widget)
}

// walk traverses the AST and converts it to a list of widgets.
func (m *Markdown1) walk(node ast.Node) []layout.Widget {
	var widgets []layout.Widget
	for child := node.FirstChild(); child != nil; child = child.NextSibling() {
		// fmt.Println("type:", child.Kind().String())
		switch n := child.(type) {
		case *ast.Text:
			htmlTags := make([]string, len(m.htmlTag))
			count := copy(htmlTags, m.htmlTag)
			fmt.Println("copy count:", count)
			fmt.Println("tags:", htmlTags)
			widgets = append(widgets, func(gtx layout.Context) layout.Dimensions {
				return m.getStyleElement(gtx, htmlTags, n)
			})
			fmt.Println("-----------Text-------------")
		case *ast.TextBlock:
			widgets = append(widgets, func(gtx layout.Context) layout.Dimensions {
				return material.Body1(material.NewTheme(), string(n.Text(m.source))).Layout(gtx)
			})
			fmt.Println("-----------TextBlock-------------")
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
			widgets = append(widgets, m.walk(n)...)
		case *ast.List:
			listWidgets := m.walk(n)
			for i, item := range listWidgets {
				index := i + 1
				if n.IsOrdered() {
					item = m.decorateListItem(fmt.Sprintf("%d. ", index), item)
				} else {
					item = m.decorateListItem("• ", item)
				}
				widgets = append(widgets, item)
			}
		case *ast.ListItem:
			widgets = append(widgets, m.walk(n)...)
		case *ast.Emphasis:
			fmt.Println("n.Level", n.Level)
			if n.Level == 1 {
				m.fontStyle = font.Italic
			} else if n.Level == 2 {
				m.fontWeight = font.Bold
			}
			widgets = append(widgets, m.walk(n)...)
		case *ast.HTMLBlock, *ast.RawHTML:
			at := n.(*ast.RawHTML).Segments.At(0)
			tag := string(at.Value(m.source))
			fmt.Println("tag:", tag)
			if strings.Contains(tag, "/") {
				m.htmlTag = nil
			} else {
				m.htmlTag = append(m.htmlTag, tag)
			}
		}
		// if child.HasChildren() {
		// 	widgets = append(widgets, r.walk(child, source)...)
		// }
	}
	return widgets
}

// decorateListItem adds bullet points or numbers to list items.
func (m *Markdown1) decorateListItem(prefix string, item layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis: layout.Horizontal,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return material.Body1(material.NewTheme(), prefix).Layout(gtx)
			}),
			layout.Flexed(1, item),
		)
	}
}
