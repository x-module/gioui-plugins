/**
 * Created by Goland
 * @file   notice.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/9/4 17:38
 * @desc   notice.go
 */

package main

import "github.com/gen2brain/beeep"

func main() {
	beeep.Alert("Title", "Message body", "assets/warning.png")
	// err := beeep.Notify("Notify title", "Message body", "assets/information.png")
	// if err != nil {
	// 	panic(err)
	// }
}
