package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pranjalch99/ecommerce-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HashPassword(password string) string {

}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {

}

func Signup() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User

		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err = Validate.Struct(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		emailCount, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if emailCount > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "this email is already in use"})
			return
		}

		phoneCount, err := UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel() //do you really need this extra defer statement?
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if phoneCount > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "this phone number is already in use"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		token, refreshtoken, _ := generate.TokenGenerator(*user.Email, *user.First_Name, *user.Last_Name, user.User_ID)
		user.Token = &token
		user.Refresh_Token = &refreshtoken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Details = make([]models.Order, 0)

		_, err = UserCollection.InsertOne(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "the user did not get created"})
			return
		}

		defer cancel() //do you really need this extra defer statement?

		c.JSON(http.StatusCreated, "Successfully signed in the user!")
	}

}

func Login() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User

		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err = UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel() //do you need this?

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password incorrect"}) //can you change the error message
			return
		}

		PasswordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()

		if !PasswordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			fmt.Println(msg)
			return
		}

		token, refreshToken, _ := generate.TokenGenerator(*foundUser.Email, *foundUser.First_Name, *foundUser.Last_Name, *foundUser.User_ID)
		defer cancel()

		generate.UpdateAllTokens(token, refreshToken, foundUser.User_ID)

		c.JSON(http.StatusFound, foundUser)
	}

}

func ProductViewerAdmin() gin.HandlerFunc {

}

func SearchProduct() gin.HandlerFunc {

}

func SearchProductByQuery() gin.HandlerFunc {

}
