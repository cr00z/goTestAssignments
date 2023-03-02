package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"gocloudcamp/internal/domain/song"
	"gocloudcamp/internal/handlers/plservice/client"
	"gocloudcamp/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	serverAddr = "0.0.0.0:50051"
)

func testCreateSong(client *client.PlaylistClient) {
	name := utils.GetInputString("name: ")
	duration := utils.GetInputInt("duration: ")

	id, err := client.CreateSong(&song.Song{
		Name:     name,
		Duration: time.Second * time.Duration(duration),
	})
	if err != nil {
		log.Println(err)
	}

	log.Println("song created with id:", id)
}

func testReadSong(client *client.PlaylistClient) {
	idsStr := utils.GetInputString("comma separated ids (e.g. 1,2,3) or blank for all: ")

	ids := make([]uint64, 0)
	if idsStr != "" {
		idsSlice := strings.Split(idsStr, ",")
		for _, elem := range idsSlice {
			id, err := strconv.ParseUint(elem, 10, 64)
			if err != nil {
				fmt.Printf("incorrect input: %v", err)
				return
			}
			ids = append(ids, id)
		}
	}

	songs, err := client.ReadSong(ids)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("___ Playlist ___")
	for _, song := range songs {
		fmt.Println(song.Id, "\t", song.Name, "\t", song.Duration)
	}
}

func main() {
	transportOpts := grpc.WithTransportCredentials(insecure.NewCredentials())

	playlistConn, err := grpc.Dial(serverAddr, transportOpts)
	if err != nil {
		log.Fatal("cannot connect to playlist server: ", err)
	}

	playlistClient := client.NewPlaylistClient(playlistConn)

	var ch rune
	for ch != 'q' {
		fmt.Print("song menu: [c] create, [r] read, [u] update, [d] delete / ")
		fmt.Print("player menu: [p] play, [s] stop, [n] next, [b] back / ")
		fmt.Println("[q] for quit")

		ch, _, err = keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}

		switch ch {
		case 'c':
			testCreateSong(playlistClient)
		case 'r':
			testReadSong(playlistClient)
		case 'p':
		case 's':
		case 'n', 'b':
		}
	}
}
