package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pranjalch99/ecommerce-golang/database"
	"github.com/pranjalch99/ecommerce-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	productCollection *mongo.Collection
	userCollection    *mongo.Collection
}

func NewApplication(productCollection *mongo.Collection, userCollection *mongo.Collection) *Application {

	return &Application{
		productCollection: productCollection,
		userCollection:    userCollection,
	}

}

func (app *Application) AddToCart() gin.HandlerFunc {

	return func(c *gin.Context) {

		productQueryId := c.Query("productId")
		if productQueryId == "" {
			log.Println("product id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}

		userQueryId := c.Query("userId")
		if userQueryId == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productId, err := primitive.ObjectIDFromHex(productQueryId)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.AddProductToCart(ctx, app.productCollection, app.userCollection, productId, userQueryId)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(200, "Successfully added to the cart")

	}

}

func (app *Application) RemoveItem() gin.HandlerFunc {

	return func(c *gin.Context) {

		productQueryId := c.Query("productId")
		if productQueryId == "" {
			log.Println("product id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("productId is empty"))
			return
		}

		userQueryId := c.Query("userId")
		if userQueryId == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("userId is empty"))
			return
		}

		productId, err := primitive.ObjectIDFromHex(productQueryId)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.RemoveCartItem(ctx, app.productCollection, app.userCollection, productId, userQueryId)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(200, "Successfully removed from the cart")

	}

}

func GetItemFromCart() gin.HandlerFunc {

	return func(c *gin.Context) {

		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid query"})
			c.Abort()
			return
		}

		userId, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var filledCart models.User

		err = UserCollection.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: userId}}).Decode(&filledCart)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(500, "not found")
			return
		}

		//Aggregate function of mongo has 3 stages -> filter matching, unwinding, grouping
		filter_match := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: userId}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "userCart"}}}}
		grouping := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$_id"}, {Key: "$total", Value: bson.D{primitive.E{Key: "$sum", Value: "$userCart.price"}}}}}}

		pointCursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{filter_match, unwind, grouping})
		if err != nil {
			log.Println(err)
		}

		var listing []bson.M

		err = pointCursor.All(ctx, &listing)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		for _, json := range listing {
			c.IndentedJSON(200, json["total"])
			c.IndentedJSON(200, filledCart.UserCart)
		}

		ctx.Done()

	}

}

func (app *Application) BuyFromCart() gin.HandlerFunc {

	return func(c *gin.Context) {

		userQueryId := c.Query("id")
		if userQueryId == "" {
			log.Println("userd Id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user Id is empty"))
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		err := database.BuyItemFromCart(ctx, app.userCollection, userQueryId)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(200, "Successfully placed the order")

	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {

	return func(c *gin.Context) {

		productQueryId := c.Query("productId")
		if productQueryId == "" {
			log.Println("productId is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("productId is empty"))
			return
		}

		userQueryId := c.Query("userId")
		if userQueryId == "" {
			log.Println("userId is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("userId is empty"))
			return
		}

		productId, err := primitive.ObjectIDFromHex(productQueryId)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = database.InstantBuyer(ctx, app.productCollection, app.userCollection, productId, userQueryId)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(200, "successfully placed the order")

	}

}
