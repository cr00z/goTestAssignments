package testfunc

import (
	"fmt"
	"gocloudcamp/internal/handlers/plservice/client"
	"gocloudcamp/pkg/plservice"
	"log"
)

func TestControl(client *client.PlaylistClient, ctrl plservice.ControlRequest_Action) {
	status, err := client.Control(ctrl)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(status)
}
