package mongo_entity

type UserFriendData struct {
	UserID         uint64   `bson:"_id"`
	Friends        []uint64 `bson:"friends"`
	FriendRequests []uint64 `bson:"friendRequests"`
}
