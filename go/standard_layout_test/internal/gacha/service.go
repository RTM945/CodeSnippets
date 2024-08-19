package gacha

import (
	"context"
	"errors"
	"standard_layout_test/internal/db"
  "log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const GachaHistory = "UserGachaHistory"

var UserGachaHistoryCollection *mongo.Collection

func InitUserGachaHistoryCollection() {
	FriendCollection = db.DB.Collection(GachaHistory)
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "uid", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := UserGachaHistoryCollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}
}

func SaveUserGachaHistory(uid uint64, poolId int32, dropList []int32, createdAt int64) error {
	filter := bson.M{"uid": uid}
	var flat []*GachaRecord
	for _, dropId := range dropList {
		flat = append(flat, &GachaRecord{
			PoolId:    poolId,
			DropId:    dropId,
			CreatedAt: createdAt,
		})
	}
	update := bson.M{"$push": bson.M{"gachaRecord": bson.M{"$each": flat}}}
	opts := options.Update().SetUpsert(true)
	_, err := UserGachaHistoryCollection.UpdateOne(context.TODO(), filter, update, opts)

	return err
}

func GetUserGachaHistory(uid uint64, pageNum, pageSize int32) ([]*GachaRecord, int32, error) {
  if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 {
		pageSize = 1
	}
	pipeline := mongo.Pipeline{
		{{"$match", bson.D{{"uid", uid}}}},
		{{"$unwind", "$gachaRecord"}},
		{{"$facet", bson.D{
			{"totalCount", bson.A{
				bson.D{{"$count", "count"}},
			}},
			{"records", bson.A{
				bson.D{{"$skip", (pageNum - 1) * pageSize}},
				bson.D{{"$limit", pageSize}},
			}},
		}}},
		{{"$project", bson.D{
			{"total", bson.D{{"$arrayElemAt", bson.A{"$totalCount.count", 0}}}},
			{"records", "$records.gachaRecord"},
		}}},
	}
	cursor, err := UserGachaHistoryCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, 0, err
	}
	type dto struct {
		Records []*GachaRecord `bson:"records"`
		Total   int32          `bson:"total"`
	}
	var result dto
	if cursor.Next(context.TODO()) {
		err := cursor.Decode(&result)
		if err != nil {
			log.Printf("Error decoding: %v", err)
		}
	}

	return result.Records, result.Total, nil
}

