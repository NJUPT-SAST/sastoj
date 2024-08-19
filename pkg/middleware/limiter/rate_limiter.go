package limiter

import (
	"sync"

	"golang.org/x/time/rate"
)

type UserRateLimiter struct {
	users map[int64]*rate.Limiter
	mu    *sync.RWMutex
	r     rate.Limit
	b     int
}

func NewUserRateLimiter(r rate.Limit, b int) *UserRateLimiter {
	u := &UserRateLimiter{
		users: make(map[int64]*rate.Limiter),
		mu:    &sync.RWMutex{},
		r:     r,
		b:     b,
	}

	return u
}

func (u *UserRateLimiter) AddUser(userID int64) *rate.Limiter {
	u.mu.Lock()
	defer u.mu.Unlock()

	limiter := rate.NewLimiter(u.r, u.b)

	u.users[userID] = limiter

	return limiter
}

func (u *UserRateLimiter) GetLimiter(userID int64) *rate.Limiter {
	u.mu.Lock()
	limiter, exists := u.users[userID]

	if !exists {
		u.mu.Unlock()
		return u.AddUser(userID)
	}

	u.mu.Unlock()

	return limiter
}

var userRateLimiters = map[int64]*UserRateLimiter{}
