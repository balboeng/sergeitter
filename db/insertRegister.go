package db

import (
	"context"
	"time"

	"github.com/balboeng/sergeitter/db/utils"
	"github.com/balboeng/sergeitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func InsertRegister(user models.User) (string, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db:= MongoConnection.Database("sergeitter")
	col:= db.Collection("Users")
	user.Password, _ = utils.EncryptPassword(user.Password)
	result, err := col.InsertOne(ctx, user)
	if err != nil{
		return "", false, err
	}

	ObjID, _:= result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}