package helpers
import(
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/akmmbh/golang-authentication/database"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type SignedDetails struct {
	Email string
	First_name string
	Last_name string
	Uid   string
	User_type string
	jwt.StandardClaims
}
var userCollection *mongo.Collection = database.OpenCollection(database.Client,"user")

var SECRET_KEY string= os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstname string, lastname string, userType string, uid string )(signedToken string , sighnedRefreshToken string, err error){
claims:=&SignedDetails{
	Email:email,
	First_name: firstname,
	Last_name:lastname,
	Uid:uid,
	User_type:userType,
	StandardClaims: jwt.StandardClaims{
		ExpiresAt:time.Now().Local().Add(time.Hour*time.Duration(24)).Unix(),
	},
}
refreshClaims:=&SignedDetails{
	StandardClaims:jwt.StandardClaims{
		ExpiresAt:time.Now().Local().Add(time.Hour*time.Duration(168)).Unix(),
	},
}
token, err:= jwt.NewWithClaims(jwt.SigningMethodES256,claims).SignedString([]byte(SECRET_KEY))
refreshtoken ,err:= jwt.NewWithClaims(jwt.SigningMethodES256,refreshClaims).SignedString([]byte(SECRET_KEY))
if err!=nil{
	log.Panic(err)
	return signedToken, sighnedRefreshToken, err
}
return token,refreshtoken,nil
}