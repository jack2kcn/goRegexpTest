/*
This is a simple practice program
for checking the regexp

version 0.1
2016-4-27
by Jack Yip

*/

package main

import (
	"regexp"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var inTE, outTE *walk.TextEdit
	var lnEdit *walk.LineEdit

	MainWindow{
		Title:   "Go语言正则表达式测试工具",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout:  Grid{Columns: 3},
				MaxSize: Size{600, 20},
				Children: []Widget{
					Label{Text: "正则表达式"},
					LineEdit{
						AssignTo: &lnEdit,
					},
					PushButton{
						Text: "Check",
						OnClicked: func() {
							outTE.SetText("")
							//outTE.SetText(strings.ToUpper(inTE.Text()))
							part := regexp.MustCompile(lnEdit.Text())
							str := inTE.Text()
							outTE.SetText(part.FindString(str))
						},
					},
				},
			},
			VSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
		},
	}.Run()
}
