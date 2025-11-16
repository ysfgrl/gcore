package gauth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var GuestUser = Claims{
	UserId:    primitive.NilObjectID,
	SessionId: primitive.NilObjectID,
	Role:      "guest",
}

type Claims struct {
	UserId    primitive.ObjectID `json:"userId" bson:"userId"`
	SessionId primitive.ObjectID `json:"sessionId" bson:"sessionId"`
	Role      string             `json:"role" bson:"role"`
	Dna       string             `json:"dna" bson:"dna"`
	IssuedAt  time.Time          `json:"iat" bson:"iat"`
	ExpiredAt time.Time          `json:"exp" bson:"exp"`
	NotBefore time.Time          `json:"nbf" bson:"nbf"`
	Audiences []string           `json:"aud" bson:"aud"`
	Scopes    []string           `json:"scope" bson:"scope"`
	Issuer    string             `json:"issuer" bson:"issuer"`
}

func (t Claims) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(t.ExpiredAt), nil
}

func (t Claims) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(t.IssuedAt), nil
}

func (t Claims) GetNotBefore() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(t.NotBefore), nil
}

func (t Claims) GetIssuer() (string, error) {
	return t.Issuer, nil
}

func (t Claims) GetSubject() (string, error) {
	return t.UserId.Hex(), nil
}
func (t Claims) GetAudience() (jwt.ClaimStrings, error) {
	return t.Audiences, nil
}
