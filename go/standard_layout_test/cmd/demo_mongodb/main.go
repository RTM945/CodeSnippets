package main

import (
	"fmt"
	"standard_layout_test/internal/mongo_service"
)

func main() {
	mongo_uri := "mongodb://localhost:27017"
	db := "test"
	mongo_service.Init(mongo_uri, db)

	mongo_service.InitFriendCollection()

	fmt.Println("==========================初始化==========================")

	var user1 uint64 = 123
	var user2 uint64 = 456
	mongo_service.InitUserFriend(user1)
	mongo_service.InitUserFriend(user2)

	show(user1, user2)

	fmt.Println("==========================user2向user1发请求==========================")
	// user2 加 user1好友
	mongo_service.AddFriendRequest(user1, user2)

	show(user1, user2)

	fmt.Println("==========================user1同意==========================")

	mongo_service.AddFriendSingle(user1, user2)

	// user1同意
	// mongo_service.AddFriend(user1, user2)

	show(user1, user2)

	fmt.Println("==========================删除好友==========================")

	// 删除好友
	mongo_service.RemoveFriendSingle(user1, user2)
	// mongo_service.RemoveFriend(user1, user2)

	show(user1, user2)

	// mongodb 单机模式不支持事务...需要另想办法...

}

func show(user1, user2 uint64) {
	user1_friend_data, err := mongo_service.GetUserFriendData(user1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("user1 = %v\r\n", user1_friend_data)

	user2_friend_data, err := mongo_service.GetUserFriendData(user2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("user2 = %v\r\n", user2_friend_data)
}
