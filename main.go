package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

func handleKeyEvent(s tcell.Screen, ch chan<- int) {
	for {
		ev := s.PollEvent() // イベントを取得する
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() { // 何のキーが押されたか？を調べる
			case tcell.KeyEscape: // ESCキーが押されたら終了する
				ch <- 1
			}
		}
	}
}

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}

	defer screen.Fini()

	ch := make(chan int)
	go handleKeyEvent(screen, ch)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	pos_x, pos_y := 5, 5
	ball := '●'

	for {
		select {
		case <-ticker.C:
			screen.Clear()
			screen.SetContent(pos_x, pos_y, ball, nil, tcell.StyleDefault)
			screen.Show()

			pos_x += 1
		case i := <-ch:
			if i == 1 {
				return
			}
		}
	}
}
