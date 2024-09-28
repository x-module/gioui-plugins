/**
 * Created by Goland
 * @file   welcome.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/16 15:07
 * @desc   welcome.go
 */

package widgets

import (
	"bytes"
	"fmt"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

type Image struct {
	theme   *theme.Theme
	src     string
	imageOp paint.ImageOp
}

func NewImage(th *theme.Theme, src string) *Image {
	image := &Image{
		theme: th,
		src:   src,
	}
	data, err := image.LoadImage(src)
	if err != nil {
		panic(err)
	}
	image.imageOp = paint.NewImageOp(data)
	return image
}
func (i *Image) Layout(gtx layout.Context) layout.Dimensions {
	return widget.Image{
		Src:      i.imageOp,
		Fit:      widget.Unscaled,
		Position: layout.Center,
		Scale:    1.0,
	}.Layout(gtx)
}

func getDefaultImage() image.Image {
	data, err := resource.Asset("data/images/no-image.png")
	if err != nil {
		panic(err)
	}
	// 解码图片
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		fmt.Println("Error decoding image: ", err)
		panic(err)
	}
	return img
}
func (i *Image) LoadImage(fileName string) (image.Image, error) {
	file, err := os.ReadFile(fmt.Sprintf("%s", fileName))
	if err != nil {
		log.Println("load image err:", err.Error())
		return getDefaultImage(), nil
	}
	// 获取fileName后缀
	temp := strings.Split(fileName, ".")
	suffix := temp[len(temp)-1]

	var img image.Image
	if suffix == "png" {
		img, err = png.Decode(bytes.NewReader(file))
	} else if suffix == "jpg" {
		img, err = jpeg.Decode(bytes.NewReader(file))
	}
	if err != nil {
		return nil, err
	}
	return img, nil
}
