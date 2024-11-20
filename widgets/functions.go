/**
 * Created by Goland
 * @file   functions.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/16 15:43
 * @desc   functions.go
 */

package widgets

import (
	"gioui.org/app"
	"github.com/x-module/gioui-plugins/theme"
	"time"
)

var NotificationController = NewNotification()
var SystemNoticeController = &SystemNotice{}

func SendAppInfoNotice(win *app.Window, th *theme.Theme, text string, duration ...time.Duration) {
	dru := time.Second * 3
	if len(duration) > 0 {
		dru = duration[0]
	}
	notice := NewNoticeItem(win, th)
	notice.Text = text
	notice.msgType = InfoMsg
	notice.EndAt = time.Now().Add(dru)
	NotificationController.AppendNotice(notice)
}
func SendAppSuccessNotice(win *app.Window, th *theme.Theme, text string, duration ...time.Duration) {
	dru := time.Second * 3
	if len(duration) > 0 {
		dru = duration[0]
	}
	notice := NewNoticeItem(win, th)
	notice.Text = text
	notice.msgType = SuccessMsg
	notice.EndAt = time.Now().Add(dru)
	NotificationController.AppendNotice(notice)
}
func SendAppWaringNotice(win *app.Window, th *theme.Theme, text string, duration ...time.Duration) {
	dru := time.Second * 3
	if len(duration) > 0 {
		dru = duration[0]
	}
	notice := NewNoticeItem(win, th)
	notice.Text = text
	notice.msgType = WaringMsg
	notice.EndAt = time.Now().Add(dru)
	NotificationController.AppendNotice(notice)
}
func SendAppErrorNotice(win *app.Window, th *theme.Theme, text string, duration ...time.Duration) {
	dru := time.Second * 3
	if len(duration) > 0 {
		dru = duration[0]
	}
	notice := NewNoticeItem(win, th)
	notice.Text = text
	notice.msgType = ErrorMsg
	notice.EndAt = time.Now().Add(dru)
	NotificationController.AppendNotice(notice)
}

func SendSystemNotice(message string) {
	_ = SystemNoticeController.Notice(message)
}
