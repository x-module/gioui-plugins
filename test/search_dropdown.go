package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
)

func main() {
	th := theme.NewTheme()
	dropDown := widgets.NewSearchDropDown(th)
	dropDown.SetOnChanged(func(value widgets.SearchDropDownOption) {
		println(dropDown.GetSelected())
	})
	dropDown.SetOptions([]*widgets.SearchDropDownOption{
		{
			Value: "save/loots",
			Text:  "保存游戏内收集",
		},
		{
			Value: "warfare/receive/task",
			Text:  "ok-认领任务",
		},
		{
			Value: "join/scav/match",
			Text:  "ok-scav模式加入比赛",
		},
		{
			Value: "warfare/get/shop/level",
			Text:  "ok-获取商店等级",
		},
		{
			Value: "get/loots",
			Text:  "---获取游戏内收集",
		},
		{
			Value: "delete/loots",
			Text:  "删除游戏内收集",
		},
		{
			Value: "warfare/receive/task/rewards",
			Text:  "ok-领取任务奖励",
		},
		{
			Value: "warfare/get/shop/level/config",
			Text:  "ok-获取商店等级配置",
		},
		{
			Value: "warfare/sale/goods",
			Text:  "ok-卖出物品",
		},
		{
			Value: "warfare/get/player/level/config",
			Text:  "ok-获取玩家等级配置",
		},
		{
			Value: "warfare/get/all/task",
			Text:  "ok-拉取任务列表",
		},
		{
			Value: "warfare/get/buffs",
			Text:  "获取buff",
		},
		{
			Value: "get/scav/match/interval",
			Text:  "ok-获取scav时间间隔",
		},
		{
			Value: "warfare/get/sale/goods",
			Text:  "ok-获取自售卖商品",
		},
		{
			Value: "warfare/set/graphics/card",
			Text:  "ok-设置bitcoin矿区显卡",
		},
		{
			Value: "warfare/get/income",
			Text:  "ok-获取bitcoin矿区收益",
		},
		{
			Value: "warfare/complete/task/by/submit/goods",
			Text:  "ok-提交物品完成任务",
		},
		{
			Value: "warfare/complete/task",
			Text:  "ok-完成任务",
		},
		{
			Value: "warfare/get/task/list",
			Text:  "ok-拉取玩家任务",
		},
		{
			Value: "warfare/consume",
			Text:  "ok-消费Money",
		},
		{
			Value: "warfare/buy/goods",
			Text:  "ok-购买商店物品",
		},
		{
			Value: "warfare/area/upgrade",
			Text:  "区域升级",
		},
		{
			Value: "warfare/add/oil",
			Text:  "ok-发动机加油",
		},
		{
			Value: "warfare/save/buffs",
			Text:  "设置buff",
		},
		{
			Value: "warfare/get/all/goods",
			Text:  "ok获取商店商品",
		},
		{
			Value: "warfare/exchange",
			Text:  "ok-以物易物",
		},
	})
	dropDown.SetWidth(unit.Dp(300))
	card := widgets.NewCard(th)
	win := window.NewApplication(new(app.Window))
	win.Title("Hello, Gio!").Size(window.ElementStyle{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return dropDown.Layout(gtx, th)
					})
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
			)
		})
	})
	win.Run()
}
