package main

import (
	"github.com/gdamore/tcell/v2"
)

func main() {
	// 画面を開く
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer screen.Fini()

	for {
		// イベントを取得する
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				// ESCキーが押されたら終了する
				return
			}
		default:
			screen.SetContent(0, 0, 'H', nil, tcell.StyleDefault)
			screen.SetContent(1, 0, 'i', nil, tcell.StyleDefault)
			screen.SetContent(2, 0, '!', nil, tcell.StyleDefault)
		}

		screen.Show()
	}
}
