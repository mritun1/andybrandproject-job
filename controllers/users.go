package controllers

import (
	"andybrandproject/auth"
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
	if db.Con == nil {
		return c.Status(500).SendString("MongoDB collection is nil")
	}

	query := bson.D{{}}
	cursor, err := db.Con.Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString("Failed to retrieve users from MongoDB: " + err.Error())
	}
	defer cursor.Close(c.Context())

	var users []models.Users

	for cursor.Next(c.Context()) {
		var user models.Users
		if err := cursor.Decode(&user); err != nil {
			return c.Status(500).SendString("Failed to decode user from MongoDB: " + err.Error())
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return c.Status(500).SendString("Cursor error during MongoDB iteration: " + err.Error())
	}

	count, err := db.Con.CountDocuments(context.Background(), query)
	if err != nil {
		return c.Status(500).SendString("Failed to get user count from MongoDB: " + err.Error())
	}

	var code int
	if count > 0 {
		code = 1
	}

	jsonData := models.OutPut{
		Code:  code,
		Count: count,
		Data:  users,
	}

	return c.Status(200).JSON(jsonData)
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
	collection := db.Con
	if collection == nil {
		return c.Status(500).SendString("MongoDB collection is nil")
	}

	users := new(models.Users)
	if err := c.BodyParser(users); err != nil {
		return c.Status(404).SendString(err.Error())
	}

	users.ID = ""
	//CONVERT PASSWORD INTO HASH
	pass, _ := auth.HashPassword(users.Password)
	users.Password = pass

	insertionResult, err := collection.InsertOne(c.Context(), users)
	if err != nil {
		return c.Status(500).SendString("Failed to insert user into MongoDB: " + err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdUsers := &models.Users{}
	if err := createdRecord.Decode(createdUsers); err != nil {
		return c.Status(500).SendString("Failed to decode created user from MongoDB: " + err.Error())
	}

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

	if db.Con == nil {
		return c.Status(500).SendString("MongoDB collection is nil")
	}

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
			},
		},
	}
	err = db.Con.FindOneAndUpdate(c.Context(), query, update).Err()
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

	if db.Con == nil {
		return c.Status(500).SendString("MongoDB collection is nil")
	}

	userId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.SendStatus(404)
	}
	query := bson.D{{Key: "_id", Value: userId}}
	result, err := db.Con.DeleteOne(c.Context(), &query)
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
// ---------------------------------------------------------------------
//
//	USER LOGIN - START
//
// ---------------------------------------------------------------------
func LoginUsers(c *fiber.Ctx) error {
	users := new(models.Users)
	if err := c.BodyParser(users); err != nil {
		return c.Status(404).SendString(err.Error())
	}
	code := 0
	token := ""
	//CHECK IF THE ACCOUNT EXISTS
	filter := bson.M{"email": users.Email}
	var user models.Users
	err := db.Con.FindOne(context.Background(), filter).Decode(&user)
	if err == nil {
		if err != mongo.ErrNoDocuments {
			//GET THE PASSWORD HASH AND VERIFY
			verify := auth.VerifyPassword(user.Password, users.Password)
			if verify {
				code = 1
				t, _ := auth.CreateKey(user.Email, user.ID)
				token = t
			}
		}
	}
	d := models.Users{}
	if code > 0 {
		d = models.Users{
			Email:    user.Email,
			Password: user.Password,
		}
	}

	jsonData := models.LogOutPut{
		Code:  code,
		Token: token,
		Data:  d,
	}
	return c.Status(200).JSON(jsonData)
}

// ---------------------------------------------------------------------
//
//	USER LOGIN - END
//
// ---------------------------------------------------------------------
