/**
 * Created by Goland
 * @file   functions.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/16 15:43
 * @desc   functions.go
 */

package widgets

import (
	"github.com/gioui-plugins/theme"
	"time"
)

var NotificationController = NewNotification()
var SystemNoticeController = &SystemNotice{}

func SendAppInfoNotice(th *theme.Theme, text string, duration ...time.Duration) {
	dru := time.Second * 4
	if len(duration) > 0 {
		dru = duration[0]
	}
	notice := NewNoticeItem(th)
	notice.Text = text
	notice.msgType = InfoMsg
	notice.EndAt = time.Now().Add(dru)
	NotificationController.AppendNotice(notice)
}
func SendAppSuccessNotice(th *theme.Theme, text string, duration ...time.Duration) {
	dru := time.Second * 4
	if len(duration) > 0 {
		dru = duration[0]
	}
	notice := NewNoticeItem(th)
	notice.Text = text
	notice.msgType = SuccessMsg
	notice.EndAt = time.Now().Add(dru)
	NotificationController.AppendNotice(notice)
}
func SendAppWaringNotice(th *theme.Theme, text string, duration ...time.Duration) {
	dru := time.Second * 4
	if len(duration) > 0 {
		dru = duration[0]
	}
	notice := NewNoticeItem(th)
	notice.Text = text
	notice.msgType = WaringMsg
	notice.EndAt = time.Now().Add(dru)
	NotificationController.AppendNotice(notice)
}
func SendAppErrorNotice(th *theme.Theme, text string, duration ...time.Duration) {
	dru := time.Second * 4
	if len(duration) > 0 {
		dru = duration[0]
	}
	notice := NewNoticeItem(th)
	notice.Text = text
	notice.msgType = ErrorMsg
	notice.EndAt = time.Now().Add(dru)
	NotificationController.AppendNotice(notice)
}

func SendSystemNotice(message string) {
	_ = SystemNoticeController.Notice(message)
}
