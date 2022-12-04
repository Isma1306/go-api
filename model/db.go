package db

import (
	"context"
	"fmt"
	"log"

	"github.com/Isma1306/go-api/types"
	"github.com/Isma1306/go-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection mongo.Collection
var ctx context.Context
var client mongo.Client

func Connect() {

	client, err := mongo.NewClient(options.Client().ApplyURI(utils.Uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx = context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection = createCollection("articles", client)

	if err != nil {
		log.Fatal(err)
	}

	//

}

func createCollection(name string, client *mongo.Client) mongo.Collection {
	database := client.Database("golang")

	return *database.Collection(name)
}

func GetAllArticles() []bson.M {
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		curErr := cursor.Err()
		fmt.Println(curErr)
		log.Fatal(err)
	}
	var articles []bson.M
	if err = cursor.All(ctx, &articles); err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	return articles
}

func CreateArticle(article types.Article) string {
	fmt.Println(" DB Create article")
	new := bson.D{{Key: "Title", Value: article.Title}, {Key: "Desc", Value: article.Desc}, {Key: "Content", Value: article.Content}}
	response, err := collection.InsertOne(ctx, new)
	if err != nil {
		log.Fatal(err)
	}
	newId := response.InsertedID.(primitive.ObjectID).Hex()

	return newId
}

func UpdateArticle(article types.Article, id string) (mongo.UpdateResult, error) {
	fmt.Println(" DB Update article")
	idPrimitive, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{
		"$set": article,
	}
	response, err := collection.UpdateByID(ctx, idPrimitive, update)
	if err != nil {
		return *response, err

	}
	fmt.Println(response)
	return *response, nil
}
func DeleteArticle(id string) (mongo.DeleteResult, error) {
	fmt.Println(" DB Delete article")
	idPrimitive, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: idPrimitive}}
	response, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return *response, err

	}
	fmt.Println(response)
	return *response, nil
}

func GetArticle(id string) (primitive.M, error) {
	fmt.Println(" DB Delete article")
	idPrimitive, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: idPrimitive}}
	var response bson.M
	err := collection.FindOne(ctx, filter).Decode(&response)
	if err != nil {
		return response, err

	}
	fmt.Println(response)
	return response, nil
}
