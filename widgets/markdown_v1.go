/**
 * Created by Goland
 * @file   markdown.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/9/18 12:05
 * @desc   markdown.go
 */

package widgets

//
// import (
// 	"gioui.org/font"
// 	"gioui.org/layout"
// 	"gioui.org/unit"
// 	"github.com/x-module/gioui-plugins/theme"
// 	"golang.org/x/exp/slices"
// 	"image/color"
// 	"strings"
// )
//
// type Element struct {
// 	Type       string            `json:"type"`
// 	Attributes map[string]string `json:"attributes"`
// 	Content    string            `json:"content"`
// 	Children   []Element         `json:"children"`
// }
//
// type Markdown struct {
// 	th       *theme.Theme
// 	elements []Element
//
// 	weight font.Weight
// 	style  font.Style
// 	color  color.NRGBA
// }
//
// // NewMarkdown creates a new Markdown widget.
// func NewMarkdown(th *theme.Theme) *Markdown {
// 	return &Markdown{
// 		th: th,
// 	}
// }
//
// func (m *Markdown) SetElements(elements []Element) {
// 	m.weight = font.Normal
// 	m.style = font.Regular
// 	m.color = m.th.Color.MarkdownDefaultColor
// 	m.elements = elements
// }
//
// const (
// 	Paragraph = "Paragraph"
// )
//
// func (m *Markdown) Layout(gtx layout.Context) layout.Dimensions {
// 	var lineElement []layout.FlexChild
// 	for _, elem := range m.elements {
// 		// 	每个一行
// 		if elem.Type == Paragraph { // 普通段落
// 			lineElement = append(lineElement, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 				return m.paragraph(gtx, elem)
// 			}))
// 		}
// 	}
// 	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, lineElement...)
// }
//
// const (
// 	StyleU     = "u"     // 下划线
// 	StyleI     = "i"     // 斜体
// 	StyleS     = "s"     // 删除线
// 	StyleMark  = "mark"  // 高亮
// 	StyleSmall = "small" // 小字体
// 	StyleBig   = "big"   // 大字体
// 	StyleBold  = "bold"  // 大字体
// )
//
// func (m *Markdown) parseTags(gtx layout.Context, tags string, child Element, son layout.Dimensions) layout.Dimensions {
// 	tagList := strings.Split(tags, "@")
// 	if len(tagList) == 0 {
// 		return m.normal(gtx, child)
// 	}
// 	var commonTag []string
// 	var specTag []string
// 	sp := []string{"u", "s", "mark"}
// 	for _, tag := range tagList {
// 		if slices.Contains(sp, tag) {
// 			specTag = append(specTag, tag)
// 		} else {
// 			commonTag = append(commonTag, tag)
// 		}
// 	}
// 	commonTag = append(commonTag, specTag...)
// 	for _, tag := range commonTag {
// 		son = m.getTagText(gtx, tag, child, son)
// 	}
// 	return son
// }
//
// func (m *Markdown) getTagText(gtx layout.Context, style string, child Element, son layout.Dimensions) layout.Dimensions {
// 	if len(child.Attributes) == 0 { // 普通文本
// 		return m.normal(gtx, child)
// 	} else if style == StyleU { // 下划线
// 		return m.underLine(gtx, func(gtx layout.Context) layout.Dimensions {
// 			if strings.TrimSpace(child.Content) != "" {
// 				return m.normal(gtx, child)
// 			}
// 			return son
// 		})
// 	} else if style == StyleS { // 删除线
// 		return m.deleteLine(gtx, func(gtx layout.Context) layout.Dimensions {
// 			if strings.TrimSpace(child.Content) != "" {
// 				return m.normal(gtx, child)
// 			} else {
// 				return son
// 			}
// 		})
// 	} else if style == StyleMark { // 高亮
// 		m.color = m.th.Color.DefaultWindowBgGrayColor
// 		return m.mark(gtx, func(gtx layout.Context) layout.Dimensions {
// 			if strings.TrimSpace(child.Content) != "" {
// 				return m.normal(gtx, child)
// 			} else {
// 				return son
// 			}
// 		})
// 	} else if style == StyleI { // 斜体
// 		return m.italic(gtx, child)
// 	} else if style == StyleSmall { // 小字体
// 		return m.small(gtx, child)
// 	} else if style == StyleBig { // 大字体
// 		return m.big(gtx, child)
// 	} else if style == StyleBold { // 大字体
// 		return m.bold(gtx, child)
// 	} else { // 普通文本
// 		return m.normal(gtx, child)
// 	}
// }
// func (m *Markdown) paragraph(gtx layout.Context, element Element) layout.Dimensions {
// 	var lineElement []layout.FlexChild
// 	for _, child := range element.Children {
// 		switch child.Type {
// 		case "Text":
// 			lineElement = append(lineElement, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 				return m.parseTags(gtx, child.Attributes["style"], child, layout.Dimensions{})
// 			}))
// 		case "Emphasis":
// 			lineElement = append(lineElement, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 				return m.parseTags(gtx, child.Attributes["style"], child, layout.Dimensions{})
// 			}))
// 		}
// 	}
// 	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, lineElement...)
// }
//
// func (m *Markdown) normal(gtx layout.Context, element Element) layout.Dimensions {
// 	if element.Content == "" {
// 		return m.parseTags(gtx, element.Children[0].Attributes["style"], element.Children[0], layout.Dimensions{})
// 	}
// 	dims := NewRichText(m.th).AddSpan([]SpanStyle{
// 		{
// 			Content:     element.Content,
// 			Color:       m.color,
// 			Size:        unit.Sp(14),
// 			Interactive: true,
// 			Font: font.Font{
// 				Typeface: "go",
// 				Weight:   m.weight,
// 				Style:    m.style,
// 			},
// 		},
// 	}).Layout(gtx)
// 	m.weight = font.Normal
// 	m.style = font.Regular
// 	m.color = m.th.Color.MarkdownDefaultColor
// 	return dims
// }
//
// func (m *Markdown) small(gtx layout.Context, element Element) layout.Dimensions {
// 	return NewRichText(m.th).AddSpan([]SpanStyle{
// 		{
// 			Content:     element.Content,
// 			Color:       m.th.Color.DefaultTextWhiteColor,
// 			Size:        unit.Sp(10),
// 			Interactive: true,
// 		},
// 	}).Layout(gtx)
// }
// func (m *Markdown) big(gtx layout.Context, element Element) layout.Dimensions {
// 	return NewRichText(m.th).AddSpan([]SpanStyle{
// 		{
// 			Content:     element.Content,
// 			Color:       m.th.Color.DefaultTextWhiteColor,
// 			Size:        unit.Sp(18),
// 			Interactive: true,
// 		},
// 	}).Layout(gtx)
// }
//
// func (m *Markdown) italic(gtx layout.Context, element Element) layout.Dimensions {
// 	return NewRichText(m.th).AddSpan([]SpanStyle{
// 		{
// 			Content:     element.Content,
// 			Color:       m.th.Color.DefaultTextWhiteColor,
// 			Size:        unit.Sp(14),
// 			Interactive: true,
// 			Font: font.Font{
// 				Typeface: "go",
// 				Weight:   font.Normal,
// 				Style:    font.Italic,
// 			},
// 		},
// 	}).Layout(gtx)
// }
// func (m *Markdown) bold(gtx layout.Context, element Element) layout.Dimensions {
// 	m.weight = font.Bold
// 	m.style = font.Regular
// 	return m.paragraph(gtx, element)
// }
// func (m *Markdown) underLine(gtx layout.Context, widget layout.Widget) layout.Dimensions {
// 	return NewRichText(m.th).UnderLineLayout(gtx, widget)
// }
// func (m *Markdown) mark(gtx layout.Context, widget layout.Widget) layout.Dimensions {
// 	return NewRichText(m.th).MarkLayout(gtx, widget)
// }
// func (m *Markdown) deleteLine(gtx layout.Context, widget layout.Widget) layout.Dimensions {
// 	return NewRichText(m.th).DeleteLineLayout(gtx, widget)
// }
