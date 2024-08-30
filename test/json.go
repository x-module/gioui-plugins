package main

//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"gioui.org/app"
// 	"gioui.org/layout"
// 	"gioui.org/op"
// 	"gioui.org/op/clip"
// 	"gioui.org/op/paint"
// 	"gioui.org/widget"
// 	"gioui.org/widget/material"
// 	"github.com/x-module/gioui-plugins/theme"
// 	"github.com/x-module/gioui-plugins/widgets"
// 	"github.com/x-module/gioui-plugins/window"
// 	"log"
// 	"strings"
// )
//
// func main() {
// 	var clickable widget.Clickable
// 	th := theme.NewTheme()
// 	win := window.NewApplication(new(app.Window))
// 	win.Title("Hello, Gio!").Size(window.ElementSize{
Height: 600,
Width:  800,
})
// 	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
// 	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
// 		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
// 			layout.Rigid(layout.Spacer{Height: 100}.Layout),
// 			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 				gtx.Constraints.Max.X = 200
// 				return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
// 					return material.Button(th.Material(), &clickable, "Hello, Gio!").Layout(gtx)
// 				})
// 			}),
// 		)
// 	})
// 	win.Run()
// }
// /**
//  * Created by Goland
//  * @file   json.go
//  * @author 李锦 <lijin@cavemanstudio.net>
//  * @date   2024/8/20 12:21
//  * @desc   json.go
//  */
//
// package main
//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"gioui.org/app"
// 	"gioui.org/layout"
// 	"gioui.org/op"
// 	"gioui.org/op/clip"
// 	"gioui.org/op/paint"
// 	"gioui.org/unit"
// 	"gioui.org/widget"
// 	"gioui.org/widget/material"
// 	"github.com/TylerBrock/colorjson"
// 	"github.com/fatih/color"
// 	"github.com/x-module/gio/resource"
// 	"github.com/x-module/gio/theme"
// 	"github.com/x-module/gio/widgets"
// 	"log"
// 	"strings"
// )
//
// type Obj struct {
// 	Name      string      `json:"name"`
// 	Age       int         `json:"age"`
// 	Married   bool        `json:"married"`
// 	Test      interface{} `json:"test"`
// 	City      string      `json:"city"`
// 	Languages []string    `json:"languages"`
// }
//
// var obj = Obj{
// 	Name:      "John Doe111",
// 	Age:       30,
// 	Married:   false,
// 	Test:      nil,
// 	City:      "New York",
// 	Languages: []string{"English", "Spanish", "German"},
// }
//
// func formatJson(target any) string {
// 	jsonData, err := json.MarshalIndent(target, "", "") // 缩进格式化
// 	if err != nil {
// 		log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
// 	}
// 	str := string(jsonData)
// 	// 将JSON字符串解析到interface{}中
// 	var obj map[string]interface{}
// 	json.Unmarshal([]byte(str), &obj)
// 	// 创建一个colorjson.Formatter对象，并自定义颜色
// 	f := colorjson.NewFormatter()
// 	f.Indent = 2                             // 设置缩进
// 	f.KeyColor = color.New(color.FgWhite)    // 设置键的颜色
// 	f.StringColor = color.New(color.FgGreen) // 设置字符串值的颜色
// 	f.NumberColor = color.New(color.FgBlue)  // 设置数字值的颜色
// 	f.BoolColor = color.New(color.FgYellow)  // 设置布尔值的颜色
// 	f.NullColor = color.New(color.FgMagenta) // 设置null值的颜色
// 	// 使用Formatter将对象格式化为彩色JSON字符串
// 	s, _ := f.Marshal(obj)
// 	// 打印彩色JSON字符串
// 	fmt.Println(string(s))
// 	ss := strings.ReplaceAll(string(s), "\u001B[37m", "#key#")
// 	ss = strings.ReplaceAll(ss, "\u001B[34m", "#number#")
// 	ss = strings.ReplaceAll(ss, "\u001B[32m", "#string#")
// 	ss = strings.ReplaceAll(ss, "\u001B[33m", "#bool#")
// 	ss = strings.ReplaceAll(ss, "\u001B[35m", "#null#")
// 	ss = strings.ReplaceAll(ss, "\u001B[0m", "||")
// 	return ss
// }
//
// var jsonList = &widget.List{
// 	List: layout.List{
// 		Axis: layout.Vertical,
// 		// ScrollToEnd: true,
// 	},
// }
// var widgetList []layout.Dimensions
// var ths = theme.New(material.NewTheme(), true)
//
// func main() {
// 	// w := new(app.Window)
// 	var ops op.Ops
// 	initialize()
// 	go func() {
// 		w := new(app.Window)
// 		for {
// 			e := w.Event()
// 			switch e := e.(type) {
// 			case app.DestroyEvent:
// 				panic(e.Err)
// 			case app.FrameEvent:
// 				gtx := app.NewContext(&ops, e)
// 				//
// 				rect := clip.Rect{
// 					Max: gtx.Constraints.Max,
// 				}
// 				paint.FillShape(gtx.Ops, resource.DefaultContentBgGrayColor, rect.Op())
// 				fmt.Println("widgetList length: ", len(flexChildList))
// 				material.List(ths.Material(), jsonList).Layout(gtx, len(flexChildList), func(gtx layout.Context, i int) layout.Dimensions {
// 					temp := flexChildList[i]
// 					return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, temp...)
// 				})
// 				// ==============================================
// 				e.Frame(gtx.Ops)
// 			}
// 		}
// 	}()
// 	app.Main()
// }
//
// var flexChildList [][]layout.FlexChild
//
// func initialize() {
// 	if len(flexChildList) > 0 {
// 		return
// 	}
// 	strs := strings.Split(formatJson(obj), "\n")
// 	for _, str := range strs {
// 		temp := strings.Split(str, "||")
// 		var temps []layout.FlexChild
// 		child := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
// 				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 					return layout.Spacer{Width: unit.Dp(5)}.Layout(gtx)
// 				}),
// 				// layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 				// 	return getLabel(fmt.Sprint(key)).Layout(gtx)
// 				// }),
// 				// layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 				// 	return layout.Spacer{Width: unit.Dp(10)}.Layout(gtx)
// 				// }),
// 			)
// 		})
// 		temps = append(temps, child)
// 		for _, t := range temp {
// 			child := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 				return getLabel(t).Layout(gtx)
// 			})
// 			temps = append(temps, child)
// 		}
// 		flexChildList = append(flexChildList, temps)
// 	}
// }
//
// func getLabel(str string) material.LabelStyle {
// 	if strings.Contains(str, "#key#") {
// 		value := strings.ReplaceAll(str, "#key#", "")
// 		label := widgets.Label(ths, value)
// 		// label.Font.Weight = font.SemiBold
// 		// label.TextSize = ths.TextSize
// 		label.Color = resource.JsonKeyColor
// 		// label.Shaper = ths.Shaper
// 		// Shaper:         th.Shaper,
// 		return label
// 		return widgets.H6(ths, value)
//
// 	} else if strings.Contains(str, "#string#") {
// 		value := strings.ReplaceAll(str, "#string#", "")
// 		label := widgets.Label(ths, value)
// 		label.Color = resource.JsonStringColor
// 		return label
// 	} else if strings.Contains(str, "#number#") {
// 		value := strings.ReplaceAll(str, "#number#", "")
// 		label := widgets.Label(ths, value)
// 		label.Color = resource.JsonNumberColor
// 		return label
// 	} else if strings.Contains(str, "#bool#") {
// 		value := strings.ReplaceAll(str, "#bool#", "")
// 		label := widgets.Label(ths, value)
// 		label.Color = resource.JsonBoolColor
// 		return label
// 	} else if strings.Contains(str, "#null#") {
// 		value := strings.ReplaceAll(str, "#null#", "")
// 		label := widgets.Label(ths, value)
// 		label.Color = resource.JsonNullColor
// 		return label
// 	} else {
// 		label := widgets.Label(ths, str)
// 		label.Color = resource.JsonStartEndColor
// 		return label
// 	}
// }
