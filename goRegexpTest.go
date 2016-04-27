package main

import (
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyMainWindow struct {
	*walk.MainWindow
}

func main() {
	var inTE, outTE *walk.TextEdit
	var lnEdit *walk.LineEdit

	mw := new(MyMainWindow)

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
							var hwnd winapi.HWND
							winapi.MessageBox(hwnd, Text("保存成功请重启生效"), Text("提示"), winapi.MB_OK)
						},
					},
				},
			},
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "Check",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}
