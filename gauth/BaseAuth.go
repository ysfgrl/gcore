package gauth

import (
	"errors"
	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ysfgrl/gcore/gerror"
	"github.com/ysfgrl/gcore/gmodel"
)

var contextKey = "userCtx"

type BaseAuth struct {
	TokenLookup string
	AuthScheme  string
	Method      jwt.SigningMethod
	PublicKey   interface{}
	PrivateKey  interface{}
	isGuest     bool
}

// TODO create IToken for JWTToken and PasetoToken
func (auth *BaseAuth) CreateToken(payload Claims) (string, *gerror.Error) {
	if auth.isGuest {
		return "", gerror.PermissionDenied
	}
	if auth.PrivateKey == nil {
		return "", gerror.TokenKeyError
	}
	token := jwt.NewWithClaims(auth.Method, payload)
	jwtToken, err1 := token.SignedString(auth.PrivateKey)
	if err1 != nil {
		return "", gerror.GetError(err1)
	}
	return jwtToken, nil
}
func (auth *BaseAuth) keyFunc(token *jwt.Token) (interface{}, error) {
	if token.Header["alg"] != auth.Method.Alg() {
		return "", errors.New(gerror.TokenAlgError.Code)
	}
	return auth.PublicKey, nil
}

func (auth *BaseAuth) GetToken(ctx *fiber.Ctx) (string, *gerror.Error) {
	contextUser, ok := ctx.Locals(contextKey).(*jwt.Token)
	if !ok {
		return "", gerror.TokenTokenError
	}
	return contextUser.Raw, nil
}

func (auth *BaseAuth) GetUser(ctx *fiber.Ctx) (*Claims, *gerror.Error) {
	contextUser := ctx.Locals(contextKey).(*jwt.Token)
	user, ok := contextUser.Claims.(*Claims)
	if !ok {
		return nil, gerror.TokenUserError
	}
	return user, nil
}

func (auth *BaseAuth) RoleRequire(ctx *fiber.Ctx, roles []string) error {
	if auth.isGuest {
		return ctx.Next()
	}
	if len(roles) == 0 {
		return ctx.Next()
	}
	user, _ := auth.GetUser(ctx)
	if user == nil {
		return auth.forbidden(ctx)
	}
	if slices.Contains(roles, user.Role) {
		return ctx.Next()
	}
	return auth.forbidden(ctx)
}

func (auth *BaseAuth) forbidden(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(gmodel.Response{
		Code:    fiber.StatusUnauthorized,
		Content: nil,
		Error:   gerror.ForbiddenError,
	})
}

func (auth *BaseAuth) Require(ctx *fiber.Ctx) error {
	if auth.isGuest {
		token := &jwt.Token{
			Header: map[string]interface{}{},
			Claims: &GuestUser,
			Method: auth.Method,
			Raw:    "GUEST_TOKEN",
		}
		ctx.Locals(contextKey, token)
		return ctx.Next()
	}
	extractors := getExtractors(auth.TokenLookup, auth.AuthScheme)
	var tokenStr string
	var err error

	for _, extractor := range extractors {
		tokenStr, err = extractor(ctx)
		if tokenStr != "" && err == nil {
			break
		}
	}
	if err != nil {
		return auth.forbidden(ctx)
	}
	var token *jwt.Token

	//t := reflect.ValueOf(&UserPayload{}).Type().Elem()
	//claims := reflect.New(t).Interface().(jwt.Claims)
	claims := &Claims{}
	token, err = jwt.ParseWithClaims(tokenStr, claims, auth.keyFunc)
	if err == nil && token.Valid {
		ctx.Locals(contextKey, token)
		return ctx.Next()
	}
	return auth.forbidden(ctx)
}
