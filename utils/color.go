/**
 * Created by Goland
 * @file   color.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/23 15:28
 * @desc   color.go
 */

package utils

import "image/color"

func MulAlpha(c color.NRGBA, alpha uint8) color.NRGBA {
	c.A = uint8(uint32(c.A) * uint32(alpha) / 0xFF)
	return c
}
