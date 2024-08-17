package main

import (
	"log"
	"standard_layout_test/internal/db"
	"standard_layout_test/internal/friend"
)

func main() {
	db.Init("mongodb://localhost:27017/", "test")

	var uid uint64 = 1
	var friendUID uint64 = 2

	_, err := friend.GetUserFriend(uid)
	if err != nil {
		log.Fatal(err)
	}
	_, err = friend.GetUserFriend(friendUID)
	if err != nil {
		log.Fatal(err)
	}
	limit := 1
	err = friend.SendFriendRequest(friendUID, uid, limit)
	if err != nil {
		log.Fatal(err)
	}
	err = friend.AgreeFriendRequest(uid, friendUID, limit)
	if err != nil {
		log.Fatal(err)
	}
}
