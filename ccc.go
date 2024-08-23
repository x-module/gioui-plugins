/**
 * Created by Goland
 * @file   ccc.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/23 14:42
 * @desc   ccc.go
 */

package main

import "image/color"

type aaa struct {
	DefaultWindowBgGrayColor  color.NRGBA
	DefaultContentBgGrayColor color.NRGBA
	DefaultBgGrayColor        color.NRGBA
	DefaultTextWhiteColor     color.NRGBA
	DefaultBorderGrayColor    color.NRGBA
	DefaultBorderBlueColor    color.NRGBA

	DefaultLineColor   color.NRGBA
	DefaultMaskBgColor color.NRGBA

	IconGrayColor            color.NRGBA
	BorderBlueColor          color.NRGBA
	BorderLightGrayColor     color.NRGBA
	HoveredBorderBlueColor   color.NRGBA
	FocusedBorderBlueColor   color.NRGBA
	ActivatedBorderBlueColor color.NRGBA
	FocusedBgColor           color.NRGBA
	TextSelectionColor       color.NRGBA
	HintTextColor            color.NRGBA

	DropDownBgGrayColor          color.NRGBA
	DropDownItemHoveredGrayColor color.NRGBA

	GreenColor   color.NRGBA
	ErrorColor   color.NRGBA
	WarningColor color.NRGBA
	SuccessColor color.NRGBA
	InfoColor    color.NRGBA

	ActionTipsBgGrayColor color.NRGBA
	ProgressBarColor      color.NRGBA

	MenuHoveredBgColor  color.NRGBA
	MenuSelectedBgColor color.NRGBA
	LogTextWhiteColor   color.NRGBA

	NotificationBgColor        color.NRGBA
	NotificationTextWhiteColor color.NRGBA
	ModalBgGrayColor           color.NRGBA

	DropdownMenuBgColor color.NRGBA
	DropdownTextColor   color.NRGBA

	NoticeInfoColor    color.NRGBA
	NoticeSuccessColor color.NRGBA
	NoticeWaringColor  color.NRGBA
	NoticeErrorColor   color.NRGBA

	JsonStartEndColor color.NRGBA
	JsonKeyColor      color.NRGBA
	JsonStringColor   color.NRGBA
	JsonNumberColor   color.NRGBA
	JsonBoolColor     color.NRGBA
	JsonNullColor     color.NRGBA
}
