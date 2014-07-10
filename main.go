package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/mattn/go-lingr"
)

func main() {
	lc := lingr.NewClient(c.User, c.Password, c.APIKey)
	err := lc.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	_, err = lc.GetRooms()
	if err != nil {
		log.Fatal(err)
	}
	err = lc.ShowRoom(strings.Join(lc.RoomIds, ","))
	if err != nil {
		log.Fatal(err)
	}
	err = lc.Subscribe(strings.Join(lc.RoomIds, ","))
	if err != nil {
		log.Fatal(err)
	}

	lc.OnMessage = func(room lingr.Room, message lingr.Message) {
		notify(
			fmt.Sprintf("%s@%s", message.Nickname, room.Name),
			message.Text,
			message.IconUrl,
		)
	}
	err = lc.Observe()
	if err != nil {
		log.Fatal(err)
	}
}
