package gacha

type GachaRecord struct {
	PoolId    int32 `bson:"poolId"`
	DropId    int32 `bson:"dropId"`
	CreatedAt int64 `bson:"createdAt"`
}
