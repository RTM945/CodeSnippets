package mongo_service

import (
	"context"
	"errors"
	"fmt"
	"standard_layout_test/internal/mongo_entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const Friend = "Friend"

var FriendCollection *mongo.Collection

func InitFriendCollection() {
	FriendCollection = DB.Collection(Friend)
}

var ErrFriendLimitExceeded = errors.New("friend limit exceeded")

func InitUserFriend(userID uint64) {
	doc := mongo_entity.UserFriendData{
		UserID:         userID,
		Friends:        make([]uint64, 0),
		FriendRequests: make([]uint64, 0),
	}
	FriendCollection.InsertOne(context.TODO(), doc)
}

func AddFriendSingle(userID, friendID uint64) error {
	_, err := FriendCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": userID},
		bson.M{
			"$addToSet": bson.M{"friends": friendID},
			"$pull":     bson.M{"friendRequests": friendID},
		},
	)
	return err
}

func RemoveFriendSingle(userID, friendID uint64) error {
	_, err := FriendCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": userID},
		bson.M{"$pull": bson.M{"friends": friendID}},
	)
	return err
}

func AddFriend(userID, friendID uint64) error {
	session, err := DBClient.StartSession()
	if err != nil {
		return fmt.Errorf("error starting session: %v", err)
	}
	defer session.EndSession(context.TODO())

	_, err = session.WithTransaction(context.TODO(), func(sessCtx mongo.SessionContext) (interface{}, error) {
		fmt.Println("Transaction started")
		// 检查 user1 的好友数量
		var user1Data mongo_entity.UserFriendData
		err := FriendCollection.FindOne(sessCtx, bson.M{"_id": userID}).Decode(&user1Data)
		if err != nil {
			return nil, fmt.Errorf("error fetching user1 data: %v", err)
		}
		fmt.Printf("user1Data: %+v\n", user1Data)
		if len(user1Data.Friends) >= 30 {
			return nil, ErrFriendLimitExceeded
		}

		// 检查 user2 的好友数量
		var user2Data mongo_entity.UserFriendData
		err = FriendCollection.FindOne(sessCtx, bson.M{"_id": friendID}).Decode(&user2Data)
		if err != nil {
			return nil, fmt.Errorf("error fetching user2 data: %v", err)
		}
		fmt.Printf("user2Data: %+v\n", user2Data)
		if len(user2Data.Friends) >= 30 {
			return nil, ErrFriendLimitExceeded
		}

		result, err := FriendCollection.UpdateOne(
			sessCtx,
			bson.M{"_id": userID},
			bson.M{
				"$addToSet": bson.M{"friends": friendID},
				"$pull":     bson.M{"friendRequests": friendID},
			},
		)
		if err != nil {
			return nil, fmt.Errorf("error adding friend userID=%d: friendId=%d: %v",
				userID, friendID, err)
		}
		fmt.Printf("UpdateOne result for user1: %+v\n", result)
		result, err = FriendCollection.UpdateOne(
			sessCtx,
			bson.M{"_id": friendID},
			bson.M{
				"$addToSet": bson.M{"friends": userID},
				"$pull":     bson.M{"friendRequests": userID},
			},
		)
		if err != nil {
			return nil, fmt.Errorf("error adding friend userID=%d: friendId=%d: %v",
				userID, friendID, err)
		}
		fmt.Printf("UpdateOne result for user2: %+v\n", result)
		return nil, nil
	})

	if err != nil {
		return fmt.Errorf("AddFriend transaction failed user1=%d, user2=%d: %v", userID, friendID, err)
	}

	return nil
}

func RemoveFriend(userID, friendID uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	session, err := DBClient.StartSession()
	if err != nil {
		return fmt.Errorf("error starting session: %v", err)
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		_, err := FriendCollection.UpdateOne(
			sessCtx,
			bson.M{"_id": userID},
			bson.M{"$pull": bson.M{"friends": friendID}},
		)
		if err != nil {
			return nil, fmt.Errorf("error removing friend userID=%d, friendId=%d: %v",
				userID, friendID, err)
		}
		_, err = FriendCollection.UpdateOne(
			sessCtx,
			bson.M{"_id": friendID},
			bson.M{"$pull": bson.M{"friends": userID}},
		)
		if err != nil {
			return nil, fmt.Errorf("error removing friend userID=%d: friendId=%d: %v",
				userID, friendID, err)
		}
		return nil, nil
	})

	if err != nil {
		return fmt.Errorf("AddFriend transaction failed user1=%d, user2=%d: %v",
			userID, friendID, err)
	}

	return nil
}

func AddFriendRequest(userID, friendID uint64) error {
	_, err := FriendCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": userID},
		bson.M{"$addToSet": bson.M{"friendRequests": friendID}},
	)
	return err
}

func RemoveFriendRequest(userID, friendID uint64) error {
	_, err := FriendCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": userID},
		bson.M{"$pull": bson.M{"friendRequests": friendID}},
	)
	return err
}

func GetUserFriendData(userID uint64) (*mongo_entity.UserFriendData, error) {
	var res mongo_entity.UserFriendData
	err := FriendCollection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &mongo_entity.UserFriendData{UserID: userID}, nil
		}
		return nil, fmt.Errorf("error fetching UserFriendData: %v", err)
	}
	return &res, nil
}
