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
	"gioui.org/widget/material"
	"gioui.org/x/richtext"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/helper/strutil"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	ast2 "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/parser"
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
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)
	m.source = content
	document := md.Parser().Parse(text.NewReader(content))
	m.fontWeight = font.Normal
	m.fontStyle = font.Regular
	m.fontColor = m.th.Color.MarkdownDefaultColor

	return m.walk(document)
}

func (m *Markdown) filterContent(content string) string {
	return strutil.Replace(content, []string{
		"~~",
	}, []string{""})
}

func (m *Markdown) normal(gtx layout.Context, node any, font font.Font, color color.NRGBA) layout.Dimensions {
	fmt.Println("tags:", m.htmlTag)
	if _, ok := node.(*ast.Text); ok {
		element, ok := node.(*ast.Text)
		if !ok {
			fmt.Println("not text node!!")
			return layout.Dimensions{}
		}
		fmt.Println("===fontweight:", m.fontWeight*1)
		fmt.Println("normal content:", string(element.Text(m.source)))
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
		fmt.Println("------------------size--------------")
		fmt.Println("size:", dims.Size)
		fmt.Println("------------------size--------------")

		return dims
	} else if _, ok = node.(*ast.TextBlock); ok {
		element, ok := node.(*ast.TextBlock)
		if !ok {
			fmt.Println("not text node!!")
			return layout.Dimensions{}
		}
		fmt.Println("===fontweight:", m.fontWeight*1)
		fmt.Println("normal content:", string(element.Text(m.source)))
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
		fmt.Println("----------------not text type-------------------------")
		return layout.Dimensions{}
	}

}

func (m *Markdown) getStyleElement(gtx layout.Context, style []string, node any, font font.Font, color color.NRGBA) layout.Dimensions {
	fmt.Println("-------------------------------------------------------------------")
	fmt.Printf("all style:%s  lenght:%d \n", style, len(style))
	if len(style) == 0 || style[0] == "" {
		return m.normal(gtx, node, font, color)
	} else {
		currentStyle := style[0]
		// 去掉第一个style后剩余的
		otherStyle := style[1:]
		fmt.Println("current style:", currentStyle)
		fmt.Printf("other style:%s\n", otherStyle)
		if currentStyle == StyleU { // 下划线
			fmt.Println("下划线")
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

// walk traverses the AST and converts it to a list of widgets.
func (m *Markdown) walk(node ast.Node) []layout.Widget {
	var widgets []layout.Widget
	for child := node.FirstChild(); child != nil; child = child.NextSibling() {
		fmt.Println("type:", child.Kind().String())
		switch n := child.(type) {
		case *ast.Text:
			fmt.Println("text all tags:", m.htmlTag)
			htmlTags := make([]string, len(m.htmlTag))
			copy(htmlTags, m.htmlTag)
			fmt.Println("----------content:", string(n.Text(m.source)))
			func(font font.Font, color color.NRGBA) {
				fmt.Println("======exec=================")
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
			fmt.Println("-----------Text-------------")
		case *ast.TextBlock:
			widgets = append(widgets, m.walk(n)...)
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
			var childs []layout.FlexChild
			for _, widget := range m.walk(n) {
				childs = append(childs, layout.Rigid(widget))
			}

			aaa := func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, childs...)
			}

			widgets = append(widgets, aaa)
		case *ast2.Strikethrough:
			widgets = append(widgets, m.walk(n)...)
		case *ast.List:
			// widgets = m.walk(n)
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
				fmt.Println("===fontweight--s:", m.fontWeight*1)

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

	return func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis: layout.Horizontal,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				body := material.Body1(material.NewTheme(), prefix)
				body.Color = m.th.Color.MarkdownDefaultColor
				return body.Layout(gtx)
			}),
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				dims := item(gtx)
				fmt.Println("dims:", dims.Size.X)
				gtx.Constraints.Min.X = dims.Size.X
				return dims
			}),
		)
	}
}
