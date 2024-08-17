package friend

type UserFriendData struct {
	UserID         uint64   `bson:"uid"`
	Friends        []uint64 `bson:"friends"`
	FriendRequests []uint64 `bson:"friendRequests"`
}
