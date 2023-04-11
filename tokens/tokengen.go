package tokens

import (
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pranjalch99/ecommerce-golang/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	Uid       string
	jwt.StandardClaims
}

var SECRET_KEY = os.Getenv("SECRET_KEY")

var UserData *mongo.Collection = database.UserData(database.Client, "Users")

func TokenGenerator(email string, firstName string, lastName string, uid string) (signedToken string, signedRefreshToken string, err error) {

	claims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Uid:       uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func ValidateToken() {

}

func UpdateAllTokens() {

}
