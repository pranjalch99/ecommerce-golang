package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pranjalch99/ecommerce-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddAddress() gin.HandlerFunc {

	return func(c *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid query"})
			c.Abort()
			return
		}

		address, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}

		var addresses models.Address
		addresses.Address_ID = primitive.NewObjectID()

		err = c.BindJSON(&addresses)
		if err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		//there will be multiple addresses associated with a user id
		match_filter := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: address}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$address"}}}}
		//will find the count of all the addresses based on address id
		group := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "addressId"}, {Key: "count", Value: bson.D{primitive.E{Key: "$sum", Value: 1}}}}}}

		pointCursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{match_filter, unwind, group})
		if err != nil {
			c.IndentedJSON(500, "internal server error")
		}

		var addressInfo []bson.M

		err = pointCursor.All(ctx, &addressInfo)
		if err != nil {
			panic(err)
		}

		var size int32

		for _, addressNumber := range addressInfo {
			count := addressNumber["count"]
			size = count.(int32)
		}

		if size < 2 {
			filter := bson.D{{Key: "_id", Value: address}}
			update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}
			_, err = UserCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				log.Println(err)
			}

		} else {
			c.IndentedJSON(400, "Not Allowed")
		}

		defer cancel()
		ctx.Done()

	}
}

func ObjectIDFromHex(user_id string) {
	panic("unimplemented")
}

func EditHomeAddress() gin.HandlerFunc {

}

func EditWorkAddress() gin.HandlerFunc {

}

func DeleteAddress() gin.HandlerFunc {

	return func(c *gin.Context) {

		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid search index"})
			c.Abort()
			return
		}

		addresses := make([]models.Address, 0)

		userId, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Internal server error")
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{primitive.E{Key: "_id", Value: userId}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}

		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(404, "Wrong command, did not update")
			return
		}

		defer cancel()
		ctx.Done()

		c.IndentedJSON(200, "Successfully deleted the address")

	}

}
