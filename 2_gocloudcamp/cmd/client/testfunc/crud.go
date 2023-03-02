package testfunc

import (
	"fmt"
	"gocloudcamp/internal/domain/song"
	"gocloudcamp/internal/handlers/plservice/client"
	"gocloudcamp/pkg/utils"
	"log"
	"strconv"
	"strings"
	"time"
)

func TestCreateSong(client *client.PlaylistClient) {
	name := utils.GetInputString("name: ")
	duration := utils.GetInputInt("duration: ")

	id, err := client.CreateSong(&song.Song{
		Name:     name,
		Duration: time.Second * time.Duration(duration),
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Println("song created with id:", id)
	}
}

func TestReadSong(client *client.PlaylistClient) {
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

func TestUpdateSong(client *client.PlaylistClient) {
	id := utils.GetInputInt("id: ")
	name := utils.GetInputString("name (blank if not changed): ")
	duration := utils.GetInputInt("duration (blank if not changed): ")

	updatedId, err := client.UpdateSong(&song.Song{
		Id:       uint64(id),
		Name:     name,
		Duration: time.Second * time.Duration(duration),
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Println("song updated with id:", updatedId)
	}
}

func TestDeleteSong(client *client.PlaylistClient) {
	id := utils.GetInputInt("id: ")
	deletedId, err := client.DeleteSong(uint64(id))
	if err != nil {
		log.Println(err)
	} else {
		log.Println("song deleted with id:", deletedId)
	}
}
