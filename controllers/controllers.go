package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jpuseche/ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HashPassword(password string) string {
	return ""
}

func VerifyPassword(userPassword, introducedPassword string) (bool, string) {
	return false, ""
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already Exists"})
		}

		count, err = UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})

		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "this phone no. is already used"})
			return
		}
		password := HashPassword(user.Password)
		user.Password = password

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.UserID = user.ID.Hex()
		token, refreshtoken, _ := generate.TokenGenerator(user.Email, user.FirstName, user.LastName, user.UserID)
		user.Token = &token
		user.RefreshToken = &refreshtoken
		user.Cart = make([]models.ProductUser, 0)
		user.Addresses = make([]models.Address, 0)
		user.Order = make([]models.Order, 0)

		_, insertErr := UserCollection.InsertOne(ctx, user)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't create user"})
			return
		}
		defer cancel()

		c.JSON(http.StatusCreated, "Successfully signed in")
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		user := models.User{}
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password incorrect"})
			return
		}
	}

	PassswordIsValid, msg := VerifyPassword(*user.Password, *founduser.Password)
	defer cancel()

	if !PasswordIsValid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(msg)
		return
	}

	token, refreshToken, _ := generate.TokenGenerator(*founduser.Email, *founduser.FirstName, *founduser.LastName, *founduser.UserID)
	defer cancel()

	generate.UdpateAllTokens(token, refreshToken, founduser.UserID)
	c.JSON(http.StatusFound, founduser)
}

func ProductViewerAdmin() gin.HandlerFunc {
	return nil
}

func SearchPassword() gin.HandlerFunc {
	return nil
}

func SearchProductByQuery() gin.HandlerFunc {
	return nil
}
