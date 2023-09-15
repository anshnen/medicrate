package tokens

import (
	"context"
	"log"
	"os"
	"time"
	"github.com/anshnen/medicrate/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct {
	Email    string
	First_Name string
	Last_Name  string
	Uid 	 string
	jwt.StandardClaims
}

var UserData *mongo.Collection = database.UserData(database.Client, "Users")

var SECRET_KEY = os.Getenv("SECRET_KEY")

func TokenGenerator(email string, firstname string, lastname string, uid string)(signedtoken string, signedrefreshtoken string, err error){
	claims := &SignedDetails{
		Email:    email,
		First_Name: firstname,
		Last_Name:  lastname,
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(12)).Unix(),
		},
	}

	refreshclaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "", "", err
	}

	refreshtoken, err := jwt.NewWithClaims(jwt.SigningMethodHS384, refreshclaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}
	return token, refreshtoken, err
}
	

func ValidateToken(signedToken string)(claims *SignedDetails, msg string){
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "token is invalid"
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token is expired"
		return
	}
	return claims, msg
}


func UpdateAllTokens(signedToken string, signedRefreshToken string, userid string){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var updateobj primitive.D

	updateobj = append(updateobj, bson.E{Key:"token", Value:signedToken})
	updateobj = append(updateobj, bson.E{Key:"refreshtoken", Value:signedRefreshToken})
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateobj = append(updateobj, bson.E{Key:"updated_at", Value:updated_at})

	upsert := true

	filter := bson.M{"user_id": userid}

	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := UserData.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateobj}}, &opt)

	defer cancel()

	if err != nil {
		log.Panic(err)
		return
	}
	
}
