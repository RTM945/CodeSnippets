package friend

import (
	"context"
	"errors"
	"fmt"
	"standard_layout_test/internal/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Friend = "Friend"

var FriendCollection *mongo.Collection

type ErrFriend struct {
	error
	Msg string
}

func InitFriendCollection() {
	FriendCollection = db.DB.Collection(Friend)
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "uid", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := FriendCollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}

}

// SendFriendRequest 不使用事务 直接在mongo操作中判断条件
func SendFriendRequest(uid, friendUID uint64, friendTotal, friendRequestTotal int) error {
	filter := bson.M{
		"uid": friendUID,
	}
	case1 := bson.M{
		"case": bson.M{
			"$and": []bson.M{
				// len(friends) < total
				bson.M{"$lt": bson.A{bson.M{"$size": "$friends"}, friendTotal}},
				// len(friendRequests) < total
				bson.M{"$lt": bson.A{bson.M{"$size": "$friendRequests"}, friendRequestTotal}},
				// !friendRequests.contains(uid)
				bson.M{"$not": bson.M{"$in": bson.A{uid, "$friendRequests"}}},
			},
		},
		"then": bson.M{"$concatArrays": bson.A{"$friendRequests", bson.A{uid}}},
	}

	cases := bson.A{case1}

	update := []bson.M{
		{
			"$set": bson.M{
				"friendRequests": bson.M{
					"$switch": bson.M{
						"branches": cases,
						"default":  "$friendRequests",
					},
				},
			},
		},
	}
	updateRes, err := FriendCollection.UpdateOne(
		context.TODO(),
		filter,
		update,
	)
	if err != nil {
		return err
	}

	if updateRes.ModifiedCount == 0 && updateRes.MatchedCount == 0 {
		err = initUserFriendDataWithRequest(friendUID, []uint64{uid})
		if err != nil {
			return err
		}
		err = initUserFriendDataWithRequest(uid, []uint64{})
		if err != nil {
			return err
		}
	}
	if updateRes.ModifiedCount != 1 {
		return ErrFriend{
			error: nil,
			Msg:   fmt.Sprintf("SendFriendRequest ModifiedCount = %d shoud be 1", updateRes.ModifiedCount),
		}
	}
	return nil
}

func initUserFriendDataWithRequest(uid uint64, friendRequests []uint64) error {
	_, err := FriendCollection.InsertOne(
		context.TODO(),
		&UserFriendData{
			UserID:         uid,
			Friends:        []uint64{},
			FriendRequests: friendRequests,
		},
	)
	if err != nil {
		if errors.Is(err, mongo.ErrInvalidIndexValue) {
			// 并发请求下可能会主键冲突 忽略
			return nil
		}
	}
	return err
}

// RejectFriendRequest 拒绝好友请求 直接从集合中删除
func RejectFriendRequest(uid, friendUID uint64) error {
	filter := bson.M{
		"uid": uid,
	}
	update := bson.M{
		"$pull": bson.M{"friendRequests": friendUID},
	}
	_, err := FriendCollection.UpdateOne(
		context.TODO(),
		filter,
		update,

		options.Update().SetUpsert(true),
	)
	return err
}

// AgreeFriendRequest 同意好友请求 不使用事务 双向操作
func AgreeFriendRequest(uid, friendUID uint64, friendTotal int) error {
	filter := bson.M{
		"uid": bson.M{
			"$in": bson.A{uid, friendUID},
		},
	}
	update := []bson.M{
		{
			"$set": bson.M{
				"friends": bson.M{
					"$switch": bson.M{
						"branches": bson.A{
							bson.M{
								"case": bson.M{
									"$and": []bson.M{
										// len(friends) < total
										bson.M{"$lt": bson.A{bson.M{"$size": "$friends"}, friendTotal}},
										bson.M{
											"$or": []bson.M{
												{
													"$and": []bson.M{
														bson.M{"$eq": bson.A{"$uid", uid}},
														bson.M{"$in": bson.A{friendUID, "$friendRequests"}},
														bson.M{"$not": bson.M{"$in": bson.A{friendUID, "$friends"}}},
													},
												},
												{
													"$and": []bson.M{
														bson.M{"$eq": bson.A{"$uid", friendUID}},
														bson.M{"$not": bson.M{"$in": bson.A{uid, "$friends"}}},
													},
												},
											},
										},
									},
								},
								"then": bson.M{
									"$switch": bson.M{
										"branches": bson.A{
											bson.M{
												"case": bson.M{"$eq": bson.A{"$uid", uid}},
												"then": bson.M{"$concatArrays": bson.A{"$friends", bson.A{friendUID}}},
											},
											bson.M{
												"case": bson.M{"$eq": bson.A{"$uid", friendUID}},
												"then": bson.M{"$concatArrays": bson.A{"$friends", bson.A{uid}}},
											},
										},
										"default": "$friends",
									},
								},
							},
						},
						"default": "$friends",
					},
				},
			},
		},
		{
			"$pull": bson.M{
				"friendRequests": bson.M{
					"$in": bson.A{uid, friendUID},
				},
			},
		},
	}

	updateRes, err := FriendCollection.UpdateMany(
		context.TODO(),
		filter,
		update,
	)
	if err != nil {
		return err
	}
	if updateRes.ModifiedCount != 2 {
		return ErrFriend{
			error: nil,
			Msg:   fmt.Sprintf("AgreeFriendRequest ModifiedCount = %d shoud be 2", updateRes.ModifiedCount),
		}
	}
}

// RemoveFriend 删好友应该就直接删除了 没有花头
func RemoveFriend(uid, friendUID uint64) {

}
