package auth

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/golang-jwt/jwt/v5"
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
	UserID     int64   `json:"user_id"`
	GroupID    int64   `json:"group_id"`
	GroupName  string  `json:"group_name"`
	ContestIDs []int64 `json:"contest_ids"`
}

const (
	authorizationKey    string = "token"
	bearerWord          string = "Bearer"
	GROUPADMIN2CONTESTS string = "groupAdmin2contests:"
)

func Auth(keyFunc jwt.Keyfunc, opMap map[string]string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			requestFromServerContext, _ := http.RequestFromServerContext(ctx)
			token, err := getTokenFromTrans(ctx)
			if err != nil {
				return nil, err
			}
			claims, err := parseToken(token, keyFunc, jwt.SigningMethodHS256)
			if err != nil {
				return nil, err
			}
			claimsInfo, err := parseClaims(*claims)
			if err != nil {
				return nil, err
			}
			//put claims into context so that other service could retrieve it
			ctx = context.WithValue(ctx, "userInfo", claimsInfo)

			log.Infof("%v", requestFromServerContext)
			if opMap != nil {
				trans, ok := transport.FromServerContext(ctx)
				if !ok {
					return nil, ErrWrongContext
				}
				oper := strings.Split(trans.Operation(), "/")[2]
				role, exists := opMap[oper]
				if !exists {
					log.Infof("Accessing Public Method: %s", trans.Operation())
					return handler(ctx, req)
				}
				switch role {
				case "user":
					if !strings.HasPrefix(claimsInfo.GroupName, "user") {
						return nil, errors.Unauthorized("UNAUTHORIZED", "Operation not allowed")
					}
				case "manager":
					if !strings.HasPrefix(claimsInfo.GroupName, "manager") {
						return nil, errors.Unauthorized("UNAUTHORIZED", "Operation not allowed")
					}
				case "admin":
					if !strings.HasPrefix(claimsInfo.GroupName, "admin") {
						return nil, errors.Unauthorized("UNAUTHORIZED", "Operation not allowed")
					}
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

func parseToken(token string, keyFunc jwt.Keyfunc, signingMethod jwt.SigningMethod) (*jwt.Claims, error) {
	var (
		tokenInfo *jwt.Token
		err       error
	)
	tokenInfo, err = jwt.Parse(token, keyFunc, jwt.WithJSONNumber())
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

func parseClaims(claims jwt.Claims) (*Claims, error) {
	userId, err := claims.(jwt.MapClaims)["user_id"].(json.Number).Int64()
	if err != nil {
		return nil, ErrMissingClaims
	}
	groupId, err := claims.(jwt.MapClaims)["group_id"].(json.Number).Int64()
	if err != nil {
		return nil, ErrMissingClaims
	}
	groupName, ok := claims.(jwt.MapClaims)["group_name"].(string)
	if !ok {
		return nil, ErrMissingClaims
	}
	contestIDs := make([]int64, 0)
	for _, id := range claims.(jwt.MapClaims)["contest_ids"].([]interface{}) {
		contestID, err := id.(json.Number).Int64()
		if err != nil {
			return nil, ErrMissingClaims
		}
		contestIDs = append(contestIDs, contestID)
	}
	return &Claims{
		UserID:     userId,
		GroupID:    groupId,
		GroupName:  groupName,
		ContestIDs: contestIDs,
	}, nil
}
