package mongodb

import (
	"context"
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"micro_apps/micro_app/config"
	"micro_apps/micro_app/models"
)

var itemRepository ItemRepository

type ItemRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func init() {
	itemRepository = ItemRepository{
		config.GetMongodbClient(),
		"testing",
		config.GetConfiguration().MongoDB.Timeout,
	}
}

func GetItemRepository() *ItemRepository {
	return &itemRepository
}

func (r *ItemRepository) Items() ([]models.Item, error) {
	collection := r.client.Database(r.database).Collection("items")

	//var item models.Item
	ctx, _ := context.WithTimeout(context.Background(), r.timeout*time.Second)
	cur, err := collection.Find(ctx, bson.D{})
	var itemList []models.Item
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		var item models.Item
		record, _ := json.Marshal(result)
		err = json.Unmarshal(record, &item)
		if err != nil {
			return nil, err
		}
		itemList = append(itemList, item)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return itemList, nil
}

func (r *ItemRepository) ItemById(id string) (models.Item, error) {

	mc := config.GetMongodbClient()
	collection := mc.Database(r.database).Collection("items")

	var item models.Item
	objID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objID}
	ctx, _ := context.WithTimeout(context.Background(), r.timeout*time.Second)
	err := collection.FindOne(ctx, filter).Decode(&item)
	if err != nil {
		return item, err
	}
	item.Id = id
	return item, nil
}

func (r *ItemRepository) Store(itemJsonString string) {
	mc := config.GetMongodbClient()
	collection := mc.Database(r.database).Collection("items")

	collection = mc.Database("testing").Collection("items")

	//Save data into Job struct
	var item models.Item
	b := []byte(itemJsonString)
	err := json.Unmarshal(b, &item)
	if err != nil {
		log.Error(err)
	}

	//Insert item into MongoDB
	ctx, _ := context.WithTimeout(context.Background(), r.timeout*time.Second)
	res, err := collection.InsertOne(ctx, item)
	if res != nil {
		log.Info("ItemId : %s", res.InsertedID)
	}
}
