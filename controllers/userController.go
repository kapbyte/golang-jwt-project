package controllers

import (
	"context"
	"net/http"
	"time"

	helper "github.com/kapbyte/golang-jwt-project/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kapbyte/golang-jwt-project/database"
	"github.com/kapbyte/golang-jwt-project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

// This function is not completed...
// func GetUsers() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)

// 		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
// 		if err != nil || recordPerPage < 1 {
// 			recordPerPage = 10
// 		}

// 	}
// }

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
