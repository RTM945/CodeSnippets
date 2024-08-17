package friend

import (
	"context"
	"errors"
	"standard_layout_test/internal/db"

	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Friend = "Friend"

var FriendCollection *mongo.Collection

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

func GetUserFriend(uid uint64) (*UserFriendData, error) {
	var res UserFriendData
	filter := bson.M{
		"uid": uid,
	}
	update := bson.M{
		"$setOnInsert": UserFriendData{
			UserID:         uid,
			Friends:        make([]uint64, 0),
			FriendRequests: make([]uint64, 0),
		},
	}
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	err := FriendCollection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&res)
	return &res, err
}

// SendFriendRequest 不使用事务 直接在mongo操作中判断条件
func SendFriendRequest(uid, friendUID uint64, friendTotal int) error {
	filter := bson.M{
		"uid":            uid,
		"friendRequests": bson.M{"$ne": friendUID},
		"friends":        bson.M{"$ne": friendUID},
		"$expr": bson.M{
			"$lt": bson.A{
				bson.M{"$size": "$friends"},
				friendTotal,
			},
		},
	}
	update := bson.M{
		"$push": bson.M{"friendRequests": friendUID},
	}

	res, err := FriendCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount != 1 && res.ModifiedCount != 1 {
		return errors.New("SendFriendRequest fail")
	}
	return nil
}

// RejectFriendRequest 拒绝好友请求 直接从集合中删除
func RemoveFriendRequest(uid, friendUID uint64) error {
	_, err := FriendCollection.UpdateOne(
		context.TODO(),
		bson.M{"uid": uid},
		bson.M{"$pull": bson.M{"friendRequests": friendUID}},
	)
	return err
}

// AgreeFriendRequest 同意好友请求 不使用事务 双向操作
func AgreeFriendRequest(uid, friendUID uint64, friendTotal int) error {
	// updateMany不能保证每个文档的原子性
	// 先更新对方
	res, err := FriendCollection.UpdateOne(
		context.TODO(),
		bson.M{
			"uid":            friendUID,
			"friendRequests": bson.M{"$ne": uid},
			"friends":        bson.M{"$ne": uid},
			"$expr": bson.M{
				"$lt": bson.A{
					bson.M{"$size": "$friends"},
					friendTotal,
				},
			},
		},
		bson.M{
			"$push": bson.M{"friend": uid},
			"$pull": bson.M{"friendRequests": uid},
		},
	)
	if err != nil {
		return err
	}
	if res.MatchedCount != 1 && res.ModifiedCount != 1 {
		return errors.New("AgreeFriendRequest add me fail")
	}
	// 再更新自己
	res, err = FriendCollection.UpdateOne(
		context.TODO(),
		bson.M{
			"uid":            uid,
			"friendRequests": friendUID,
			"friends":        bson.M{"$ne": friendUID},
			"$expr": bson.M{
				"$lt": bson.A{
					bson.M{"$size": "$friends"},
					friendTotal,
				},
			},
		},
		bson.M{
			"$push": bson.M{"friend": friendUID},
			"$pull": bson.M{"friendRequests": friendUID},
		},
	)
	if err != nil || res.MatchedCount != 1 && res.ModifiedCount != 1 {
		// 回滚
		my, err := GetUserFriend(uid)
		if err != nil {
			return err
		}
		if lo.Contains(my.Friends, friendUID) {
			// 如果已经有好友了 就清理下好友请求
			RemoveFriendRequest(uid, friendUID)
		} else {
			// 如果没有加上好友 就删除对方的好友
			_, err = FriendCollection.UpdateOne(
				context.TODO(),
				bson.M{"uid": friendUID},
				bson.M{"$pull": bson.M{"friends": uid}},
			)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// RemoveFriend 双向删除
func RemoveFriend(uid, friendUID uint64) error {
	_, err := FriendCollection.UpdateOne(
		context.TODO(),
		bson.M{"uid": uid},
		bson.M{"$pull": bson.M{"friends": friendUID}},
	)
	if err != nil {
		return err
	}
	_, err = FriendCollection.UpdateOne(
		context.TODO(),
		bson.M{"uid": friendUID},
		bson.M{"$pull": bson.M{"friends": uid}},
	)
	if err != nil {
		return err
	}
	return nil
}
