package auth

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrMissingClaims          = errors.Unauthorized("UNAUTHORIZED", "Claims is missing")
	ErrMissingJwtToken        = errors.Unauthorized("UNAUTHORIZED", "JWT token is missing")
	ErrTokenInvalid           = errors.Unauthorized("UNAUTHORIZED", "Token is invalid")
	ErrTokenExpired           = errors.Unauthorized("UNAUTHORIZED", "JWT token has expired")
	ErrTokenParseFail         = errors.Unauthorized("UNAUTHORIZED", "Fail to parse JWT token ")
	ErrUnSupportSigningMethod = errors.Unauthorized("UNAUTHORIZED", "Wrong signing method")
	ErrWrongContext           = errors.Unauthorized("UNAUTHORIZED", "Wrong context for middleware")
	ErrSignToken              = errors.Unauthorized("UNAUTHORIZED", "Can not sign token. Is the key correct?")
)

type Claims struct {
	UserId    int64   `json:"user_id"`
	GroupIds  []int64 `json:"group_ids"`
	GroupName string  `json:"group_name"`
}

// Authorization
const (
	authorizationKey string = "token"
)

// GroupName
const (
	PublicGroup string = "public"
	UserGroup   string = "user"
	AdminGroup  string = "admin"
)

// Auth middleware.
// secret: jwt secret;
// defaultRule: default rule for api;
// customApiMap: custom api rule.
func Auth(secret string, defaultRule string, customApiMap map[string]string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			// get token from transport and parse it
			token, err := getTokenFromTrans(ctx)
			if err != nil {
				return nil, err
			}
			keyFunc := func(token *jwt.Token) (interface{}, error) { return []byte(secret), nil }
			claims, err := parseToken(token, keyFunc, jwt.SigningMethodHS256)
			if err != nil {
				return nil, err
			}
			claimsInfo, err := parseClaims(*claims)
			if err != nil {
				return nil, err
			}
			log.Infof("receive request from user: %d", claimsInfo.UserId)

			//put claims into context so that other service could retrieve it
			ctx = context.WithValue(ctx, "userInfo", claimsInfo)

			// if the user is an admin, then return directly
			if strings.HasPrefix(claimsInfo.GroupName, AdminGroup) {
				log.Infof("receive request from admin: %d", claimsInfo.UserId)
				return handler(ctx, req)
			}

			// check the role of the user
			trans, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, ErrWrongContext
			}
			operation := trans.Operation()
			role := defaultRule
			if customApiMap != nil && customApiMap[operation] != "" {
				role = customApiMap[operation]
			}
			switch role {
			case PublicGroup:
				return handler(ctx, req)
			case UserGroup, AdminGroup:
				if !strings.HasPrefix(claimsInfo.GroupName, role) {
					return nil, errors.Unauthorized("PERMISSION_DENIED", "Wrong role")
				}
			default:
				return nil, errors.InternalServer("INTERNAL_SERVER", "Wrong rule")
			}
			return handler(ctx, req)
		}
	}
}

func getTokenFromTrans(ctx context.Context) (string, error) {
	if header, ok := transport.FromServerContext(ctx); ok {
		return header.RequestHeader().Get(authorizationKey), nil
	}
	log.Warn("Can not get token from transport")
	return "", ErrMissingJwtToken
}

func parseToken(token string, keyFunc jwt.Keyfunc, signingMethod jwt.SigningMethod) (*jwt.Claims, error) {
	var (
		tokenInfo *jwt.Token
		err       error
	)
	tokenInfo, err = jwt.Parse(token, keyFunc, jwt.WithJSONNumber())
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) || errors.Is(err, jwt.ErrTokenUnverifiable) {
			return nil, ErrSignToken
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

func parseClaims(claims jwt.Claims) (*Claims, error) {
	userID, err := claims.(jwt.MapClaims)["user_id"].(json.Number).Int64()
	if err != nil {
		return nil, ErrMissingClaims
	}
	groupIds := make([]int64, 0)
	for _, id := range claims.(jwt.MapClaims)["group_ids"].([]interface{}) {
		groupID, err := id.(json.Number).Int64()
		if err != nil {
			return nil, ErrMissingClaims
		}
		groupIds = append(groupIds, groupID)
	}
	log.Info(groupIds)
	if err != nil {
		return nil, ErrMissingClaims
	}
	groupName, ok := claims.(jwt.MapClaims)["group_name"].(string)
	if !ok {
		return nil, ErrMissingClaims
	}
	return &Claims{
		UserId:    userID,
		GroupIds:  groupIds,
		GroupName: groupName,
	}, nil
}
