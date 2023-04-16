package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/pranjalch99/ecommerce-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrCantFindProduct    = errors.New("can't find the cart item")
	ErrCantDecodeProducts = errors.New("can't decode the cart item")
	ErrUserIdInvalid      = errors.New("the user ID is not invalid")
	ErrCantUpdateUser     = errors.New("cannot add this product to the cart")
	ErrCantRemoveCartItem = errors.New("cannot remove this item from the cart")
	ErrCantGetItem        = errors.New("not able to get the item from the cart")
	ErrCantBuyCartItem    = errors.New("cannot update the purchase")
)

func AddProductToCart(ctx context.Context, prodCollection, userCollection *mongo.Collection, productId primitive.ObjectID, userId string) error {

	searchFromDb, err := prodCollection.Find(ctx, bson.M{"_id": productId})
	if err != nil {
		log.Println(err)
		return ErrCantFindProduct
	}

	var productCart []models.ProductUser

	err = searchFromDb.All(ctx, &productCart)
	if err != nil {
		log.Println(err)
		return ErrCantDecodeProducts
	}

	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Println(err)
		return ErrUserIdInvalid
	}

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "userCart", Value: bson.D{primitive.E{Key: "each", Value: productCart}}}}}}

	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return ErrCantUpdateUser
	}

	return nil
}

func RemoveCartItem(ctx context.Context, prodCollection, userCollection *mongo.Collection, productId primitive.ObjectID, userId string) error {

	userID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Println(err)
		return ErrUserIdInvalid
	}

	filter := bson.D{primitive.E{Key: "_id", Value: userID}}
	update := bson.M{"$pull": bson.M{"userCart": bson.M{"_id": productId}}}

	_, err = userCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		log.Println(err)
		return ErrCantRemoveCartItem
	}

	return nil

}

func BuyItemFromCart(ctx context.Context, userCollection *mongo.Collection, userId string) error {

	//fetch the cart of the user
	//find the cart total
	//create an order with these items
	//empty up the cart

	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Println(err)
		return ErrUserIdInvalid
	}

	var getCartItems models.User
	var orderCart models.Order

	orderCart.Order_ID = primitive.NewObjectID()
	orderCart.Ordered_At = time.Now()
	orderCart.Payment_Method.COD = true
	orderCart.Order_Cart = make([]models.ProductUser, 0)

}

func InstanBuyer() {

}
