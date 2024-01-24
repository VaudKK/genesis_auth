package service

import (
	"context"
	"errors"
	"fmt"
	"genesis_auth/config"
	"genesis_auth/dto"
	"genesis_auth/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Email          string `json:"email"`
	OrganizationId string `json:"organizationId"`
	jwt.StandardClaims
}

type AuthenticationService interface {
	CreateUser(*dto.UserDto) (interface{}, error)
	GenerateToken(*model.User, string) (*dto.TokenDto, error)
	LogIn(*dto.LogInDto) (*dto.TokenDto, error)
}

type authenticationService struct {
	collection *mongo.Collection
}

func NewAuthenticationService(collection *mongo.Collection) AuthenticationService {
	return &authenticationService{
		collection: collection,
	}
}

func (authService *authenticationService) CreateUser(userDto *dto.UserDto) (interface{}, error) {
	filter := bson.D{
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "email", Value: userDto.Email}},
			bson.D{{Key: "phone", Value: userDto.Phone}},
		}},
	}

	var usrs []model.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := authService.collection.Find(ctx, filter)

	cursor.All(context.TODO(), &usrs)

	if err != nil {
		return nil, err
	}

	if len(usrs) > 0 {
		return nil, fmt.Errorf("user exists")
	} else {
		hash, hashErr := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
		if hashErr != nil {
			return nil, hashErr
		}

		usr := model.User{
			ID:        primitive.NewObjectID(),
			FirstName: userDto.FirstName,
			LastName:  userDto.LastName,
			Phone:     userDto.Phone,
			Email:     userDto.Email,
			Password:  string(hash),
		}

		result, insertErr := authService.collection.InsertOne(ctx, usr)

		if err != nil {
			return nil, insertErr
		}

		return result, nil
	}
}

func (authService *authenticationService) LogIn(logInDto *dto.LogInDto) (*dto.TokenDto, error) {

	filter := bson.D{
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "email", Value: logInDto.Identifier}},
			bson.D{{Key: "phone", Value: logInDto.Identifier}},
		}},
	}

	var usr model.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := authService.collection.FindOne(ctx, filter).Decode(&usr)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no user found")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(logInDto.Password))

	if err != nil {
		return nil, err
	}

	token, err := authService.GenerateToken(&usr, "access")

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (authService *authenticationService) GenerateToken(user *model.User, tokenType string) (*dto.TokenDto, error) {
	var expiry int64
	var subject string

	if tokenType == "access" {
		expiry = time.Now().Add(time.Hour * 1).Unix()
		subject = user.ID.String()
	} else if tokenType == "refresh" {
		expiry = time.Now().Add(time.Hour * 24 * 3).Unix()
		subject = "refresh"
	}

	claims := Claims{
		Email:          user.Phone,
		OrganizationId: user.OrganizationId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiry,
			Issuer:    "GenesisAuth",
			Subject:   subject,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(config.AppConfig.JWT_KEY))

	if err != nil {
		return nil, err
	}

	return &dto.TokenDto{
		Token:  accessToken,
		Expiry: expiry,
	}, nil

}

func ValidateToken(tokenString string) (*Claims, error) {
	parsedToken, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signature")
		}

		return []byte(config.AppConfig.JWT_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*Claims)

	if !ok || !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
