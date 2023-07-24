package controllers

import (
	"andybrandproject/db"
	"andybrandproject/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ---------------------------------------------------------------------
//
//	GET LIST OF USERS - START
//
// ---------------------------------------------------------------------
func Users(c *fiber.Ctx) error {
	query := bson.D{{}}
	var collect = db.Mg.Db.Collection("users")
	cursor, err := collect.Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var users []models.Users = make([]models.Users, 0)
	if err := cursor.All(c.Context(), &users); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	//Count Documents num
	col, _ := collect.CountDocuments(context.Background(), query)
	var code int = 0
	if col > 0 {
		code = 1
	}
	jsonData := models.OutPut{
		Code:  code,
		Count: col,
		Data:  users,
	}
	return c.JSON(jsonData)

}

// ---------------------------------------------------------------------
//
//	GET LIST OF USERS - END
//
// ---------------------------------------------------------------------
// ---------------------------------------------------------------------
//
//	INSERT USERS DATA - START
//
// ---------------------------------------------------------------------
func CreateUsers(c *fiber.Ctx) error {

	collection := db.Mg.Db.Collection("users")
	users := new(models.Users)
	if err := c.BodyParser(users); err != nil {
		return c.Status(404).SendString(err.Error())
	}

	users.ID = ""
	insertionResult, err := collection.InsertOne(c.Context(), users)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdUsers := &models.Users{}
	createdRecord.Decode(createdUsers)

	jsonData := models.OutPut{
		Code:  1,
		Count: 1,
		Data:  createdUsers,
	}

	return c.Status(200).JSON(jsonData)
}

// ---------------------------------------------------------------------
//
//	INSERT USERS DATA - END
//
// ---------------------------------------------------------------------
// ---------------------------------------------------------------------
//
//	UPDATE USERS DATA - START
//
// ---------------------------------------------------------------------
func UpdateUsers(c *fiber.Ctx) error {
	idParam := c.Params("id")
	usersId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.SendStatus(404)
	}
	users := new(models.Users)
	if err := c.BodyParser(users); err != nil {
		return c.Status(404).SendString(err.Error())
	}
	query := bson.D{{Key: "_id", Value: usersId}}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "name", Value: users.Name},
				{Key: "dob", Value: users.Dob},
				{Key: "address", Value: users.Address},
				{Key: "description", Value: users.Description},
				{Key: "createdAt", Value: users.CreatedAt},
			},
		},
	}
	err = db.Mg.Db.Collection("users").FindOneAndUpdate(c.Context(), query, update).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.SendStatus(500)
	}
	users.ID = idParam

	jsonData := models.OutPut{
		Code:  1,
		Count: 1,
		Data:  users,
	}

	return c.Status(200).JSON(jsonData)
}

// ---------------------------------------------------------------------
//
//	UPDATE USERS DATA - END
//
// ---------------------------------------------------------------------
// ---------------------------------------------------------------------
//
//	DELETE USERS DATA - START
//
// ---------------------------------------------------------------------
func Delete(c *fiber.Ctx) error {
	userId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.SendStatus(404)
	}
	query := bson.D{{Key: "_id", Value: userId}}
	result, err := db.Mg.Db.Collection("users").DeleteOne(c.Context(), &query)
	if err != nil {
		return c.SendStatus(500)
	}
	var code int = 1
	if result.DeletedCount < 1 {
		code = 0
	}
	jsonData := models.StandardOutPut{
		Code:  code,
		Count: 1,
	}
	return c.Status(200).JSON(jsonData)
}

// ---------------------------------------------------------------------
//
//	DELETE USERS DATA - END
//
// ---------------------------------------------------------------------
