package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"gocloudcamp/internal/domain/playlist"
	"gocloudcamp/internal/domain/song"
	"gocloudcamp/internal/player"
	"time"
)

func main() {
	var ch rune
	var err error
	pl := playlist.NewPlaylist()
	go player.Start(pl)

	fmt.Println("Player demo")

	for ch != 'q' {
		fmt.Println("p play, s stop, n next, b back, a add, q quit")

		ch, _, err = keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}

		if pl.Played() {
			fmt.Println()
		}

		switch ch {
		case 'p':
			out := "OK"
			err := pl.Play()
			if err != nil {
				out = err.Error()
			}
			fmt.Println("play:", out)

		case 's':
			out := "OK"
			err := pl.Pause()
			if err != nil {
				out = err.Error()
			}
			fmt.Println("pause:", out)

		case 'n', 'b':
			var out string
			if ch == 'n' {
				err = pl.Next()
				out = "next: "
			} else {
				err = pl.Prev()
				out = "prev: "
			}

			if err != nil {
				out += err.Error()
			} else {
				song, _ := pl.CurrentSong()
				out += "song: " + song.Name
			}

			fmt.Println(out)

		case 'a':
			played := pl.Played()
			if played {
				pl.Pause()
				time.Sleep(player.SongTick)
			}

			var name string
			var duration int
			fmt.Print("name: ")
			fmt.Scanf("%s\n", &name)
			fmt.Print("duration (sec): ")
			fmt.Scanf("%d\n", &duration)
			pl.AddSong(
				&song.Song{
					Name:     name,
					Duration: time.Duration(duration) * time.Second,
				},
			)

			if played {
				pl.Play()
			}
		}
	}
}
