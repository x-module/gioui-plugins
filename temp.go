package main

import (
	"bytes"
	"fmt"
	config "github.com/x-module/gioui-plugins/resource"
	"image"
	// 导入image包以支持基本的图片操作
	_ "image/png"
)

func main() {
	data, err := config.Asset("data/images/no-image.png")
	if err != nil {
		panic(err)
	}
	// 解码图片
	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		fmt.Println("Error decoding image: ", err)
		return
	}
	// 打印图片的格式和尺寸
	fmt.Println("Image format: ", format)
	fmt.Printf("Image size: %dx%d\n", img.Bounds().Dx(), img.Bounds().Dy())
}
