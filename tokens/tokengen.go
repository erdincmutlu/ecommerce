package tokens

import (
	"log"
	"os"
	"time"

	"github.com/erdincmutlu/ecommerce/database"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	Uid       string
	jwt.Claims
}

var UserData *mongo.Collection = database.UserData(database.Client, "Users")

var SECRET_KEY = os.Getenv("SECRET_KEY")

func TokenGenerator(email string, firstName string, lastName string, uid string,
) (string, string, error) {
	claims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Uid:       uid,
		Claims: jwt.Claims{
			ExpiresAt: time.Now().Local().Add(24 * time.Hour).Unix(),
		},
	}

	refreshClaims := SignedDetails{
		Claims: jwt.Claims{
			ExpiresAt: time.Now().Local().Add(168 * time.Hour).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Println(err.Error())
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Println(err.Error())
		return "", "", err
	}

	return token, refreshToken, nil
}

func ValidateToken(signedToken string) (*SignedDetails, string) {
	token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return &SignedDetails{}, err.Error()
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return &SignedDetails{}, "the token is invalid"
	}

	expirationTime, err := claims.GetExpirationTime()
	if err != nil {
		return &SignedDetails{}, err.Error()
	}
	if expirationTime.Before(time.Now()) {
		return &SignedDetails{}, "token is already expired"
	}
	return claims, ""
}

func UpdateAllTokens() {

}
