package auth

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

var (
	ErrMissingClaims          = errors.Unauthorized("UNAUTHORIZED", "Claims is missing")
	ErrMissingJwtToken        = errors.Unauthorized("UNAUTHORIZED", "JWT token is missing")
	ErrMissingKeyFunc         = errors.Unauthorized("UNAUTHORIZED", "keyFunc is missing")
	ErrTokenInvalid           = errors.Unauthorized("UNAUTHORIZED", "Token is invalid")
	ErrTokenExpired           = errors.Unauthorized("UNAUTHORIZED", "JWT token has expired")
	ErrTokenParseFail         = errors.Unauthorized("UNAUTHORIZED", "Fail to parse JWT token ")
	ErrUnSupportSigningMethod = errors.Unauthorized("UNAUTHORIZED", "Wrong signing method")
	ErrWrongContext           = errors.Unauthorized("UNAUTHORIZED", "Wrong context for middleware")
	ErrNeedTokenProvider      = errors.Unauthorized("UNAUTHORIZED", "Token provider is missing")
	ErrSignToken              = errors.Unauthorized("UNAUTHORIZED", "Can not sign token.Is the key correct?")
	ErrGetKey                 = errors.Unauthorized("UNAUTHORIZED", "Can not get key while signing token")
)

type Claims struct {
	userId    int64
	groupId   int64
	groupName string
	//contestIds []int64
}

const (
	authorizationKey    string = "token"
	bearerWord          string = "Bearer"
	GROUPADMIN2CONTESTS string = "groupAdmin2contests:"
)

//type Repo struct {
//	data *Data
//	log  *log.Helper
//}

func Auth(keyFunc jwt.Keyfunc, operMap map[string]string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			requ, _ := http.RequestFromServerContext(ctx)

			log.Infof("s", requ)
			trans, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, ErrWrongContext
			}
			oper := strings.Split(trans.Operation(), "/")[2]
			role, exists := operMap[oper]
			if !exists {
				log.Infof("Accessing Public Method: ", trans.Operation())
				return handler(ctx, req)
			}

			token, err := getTokenFromTrans(ctx)
			if err != nil {
				return nil, err
			}
			claims, err := praseToken(token, keyFunc, jwt.SigningMethodHS256)
			if err != nil {
				return nil, err
			}
			claimsInfo, err := praseClaims(*claims)
			if err != nil {
				return nil, err
			}

			switch role {
			case "user":
				if !strings.HasPrefix(claimsInfo.groupName, "user") {
					return nil, errors.Unauthorized("UNAUTHORIZED", "Operation not allowed")
				}
			case "admin":
				if !strings.HasPrefix(claimsInfo.groupName, "admin") {
					return nil, errors.Unauthorized("UNAUTHORIZED", "Operation not allowed")
				}
			}
			return handler(ctx, req)
		}
	}
}
func getTokenFromTrans(ctx context.Context) (string, error) {
	if header, ok := transport.FromServerContext(ctx); ok {
		return header.RequestHeader().Get(authorizationKey), nil
	}
	log.Debug("Going Wrong when trying to get token from transport")
	return "", ErrWrongContext
}

func praseToken(token string, keyfunc jwt.Keyfunc, signingMethod jwt.SigningMethod) (*jwt.Claims, error) {
	var (
		tokenInfo *jwt.Token
		err       error
	)
	tokenInfo, err = jwt.Parse(token, keyfunc)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) || errors.Is(err, jwt.ErrTokenUnverifiable) {
			return nil, ErrTokenInvalid
		}
		if errors.Is(err, jwt.ErrTokenNotValidYet) || errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrTokenParseFail
	}
	if !tokenInfo.Valid {
		return nil, ErrTokenInvalid
	}
	if tokenInfo.Method != signingMethod {
		return nil, ErrUnSupportSigningMethod
	}
	return &tokenInfo.Claims, nil
}

func praseClaims(claims jwt.Claims) (*Claims, error) {

	userId, ok := claims.(jwt.MapClaims)["userId"].(float64)
	if !ok {
		return nil, ErrMissingClaims
	}
	groupId, ok := claims.(jwt.MapClaims)["groupId"].(float64)
	if !ok {
		return nil, ErrMissingClaims
	}
	groupName, ok := claims.(jwt.MapClaims)["groupName"].(string)
	if !ok {
		return nil, ErrMissingClaims
	}
	return &Claims{
		userId:    int64(userId),
		groupId:   int64(groupId),
		groupName: groupName,
	}, nil
}
