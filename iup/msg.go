/**
 * Created by Goland
 * @file   msg.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/9/4 17:23
 * @desc   msg.go
 */

package main

import (
	"fmt"
	"github.com/gen2brain/dlgs"
)

// https://github.com/gen2brain/dlgs?tab=readme-ov-file
// Variables
// func Color(title, defaultColorHex string) (color.Color, bool, error)
// func Date(title, text string, defaultDate time.Time) (time.Time, bool, error)
// func Entry(title, text, defaultText string) (string, bool, error)
// func Error(title, text string) (bool, error)
// func File(title, filter string, directory bool) (string, bool, error)
// func FileMulti(title, filter string) ([]string, bool, error)
// func Info(title, text string) (bool, error)
// func List(title, text string, items []string) (string, bool, error)
// func ListMulti(title, text string, items []string) ([]string, bool, error)
// func MessageBox(title, text string) (bool, error)
// func Password(title, text string) (string, bool, error)
// func Question(title, text string, defaultCancel bool) (bool, error)
// func Warning(title, text string) (bool, error)

func main() {
	// _, _, err := dlgs.Password("Password", "Enter your API key:")
	// if err != nil {
	// 	panic(err)
	// }
	// _, _, err := dlgs.Color("Pick color", "#BEBEBE")
	// if err != nil {
	// 	panic(err)
	// }

	// _, err := dlgs.Error("Error", "Cannot divide by zero.")
	// if err != nil {
	// 	panic(err)
	// }
	_, err := dlgs.Question("Question", "Are you sure you want to format this media?", true)
	if err != nil {
		panic(err)
	}

	item, _, err := dlgs.List("List", "Select item from list:", []string{"Bug", "New Feature", "Improvement"})
	if err != nil {
		panic(err)
	}
	fmt.Println("Selected item:", item)

}
