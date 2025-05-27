package controllers
import(
"context"
"fmt"
"log"
"strconv"
"net/http"
"time"
"github.com/gin-gonic/gin"
"github.com/go-playground/validator/v10"
helper "github.com/akmmbh/golang-authentication/helpers"
"github.com/akmmbh/golang-authentication/models"
//"github/akmmbh/golang-authentication/helpers"
"golang.org/x/crypto/bcrypt"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/bson/primitive"
"github.com/go-playground/validator/v10"
"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client,"user")
var validate =validator.New()
func HashPassword()

func VerifyPassword()

func Signup()gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx,cancel= context.WithTimeout(context.Background(),100*time.Second)
		var user models.User
		if err:= c.BindJSON(&user);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return 
		}
		validateErr:=validate.Struct(user)
		if validateErr!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":validateErr.Error()})	
		return 
		}

		count,err:=userCollection.CountDocuments(ctx,bson.M{"email":user.Email})
		defer cancel()
		if err!=nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error occured while checking for the email"})

		}
      count1,err:= userCollection.CountDocuments(ctx,bson.M{"phone":user.Phone})
	  defer cancel()
	  if err!=nil{
		log.Panic(err)
		c.JSON(http.StatusInternalServerError,gin.H{"error":"error occured while checking th phone numbers"})

	  }
	  if count >0 || count1>0{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"this email or phone number already exists"})
	  }
	  user.Created_at ,_ =time.Parse(time.RFC3339 , time.Now().Format(time.RFC3339))
	  user.Updated_at ,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
	  user.ID=primitive.NewObjectID()
       user.User_id=user.ID.Hex() 
        token ,refreshToken,_ :=helper.GenerateAllTokens(*user.Email,*user.First_name,*user.Last_name,*user.User_type,*&user.User_id)
		user.Token=&token
		user.Refresh_token=&refreshToken
		resultInsertionNumber,insertErr :=userCollection.InsertOne(ctx,user)
		if insertErr!=nil{
			msg:=fmt.Sprintf("user item was not created")
			c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
			return 
		}
		defer cancel()
		c.JSON(http.StatusOK,resultInsertionNumber)

	}
}

func Login()

func GetUsers()
//only admin can get acces of user data no other can access it
func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId:= c.Param("userId")
	 if err:= helper.MatchUserTypeToUid(c,userId);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return 
	 }
	 var ctx,cancel =context.WithTimeout(context.Background(),100*time.Second)
	 defer cancel()

	 var user models.User
	 err:=userCollection.FindOne(ctx,bson.M{"user_id":userId}).Decode(&user)
     if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return 
	 }
	 c.JSON(http.StatusOK,user)


	}
}