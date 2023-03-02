package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"gocloudcamp/cmd/client/testfunc"
	"gocloudcamp/internal/handlers/plservice/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	serverAddr = "0.0.0.0:50051"
)

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
			testfunc.TestCreateSong(playlistClient)
		case 'r':
			testfunc.TestReadSong(playlistClient)
		case 'u':
			testfunc.TestUpdateSong(playlistClient)
		case 'd':
			testfunc.TestDeleteSong(playlistClient)
		case 'p', 's', 'n', 'b':
			// TODO: implement me
			fmt.Print("not implemented")
		}
	}
}
