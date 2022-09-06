package main

import "fmt"

import "github.com/nsf/termbox-go"

func main33() {

	err := termbox.Init()

	if err != nil {

		panic(err)

	}

	defer termbox.Close()

	for {

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:

			switch ev.Key {

			case termbox.KeyEsc:

				fmt.Println("You press Esc")

			case termbox.KeyF1:

				fmt.Println("You press F1")

			default:

				fmt.Println(ev)
				//break Loop

			}

		}

	}

}
