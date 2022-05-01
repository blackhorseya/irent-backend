package rate

import (
	"sync"

	"github.com/blackhorseya/irent/internal/pkg/entity/er"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// IP is a ip rate limiter
type IP struct {
	ips   map[string]*rate.Limiter
	mu    *sync.Mutex
	limit rate.Limit
	burst int
}

// NewIP serve caller to new an IP rate limiter
func NewIP(limit rate.Limit, burst int) *IP {
	return &IP{
		ips:   make(map[string]*rate.Limiter),
		mu:    &sync.Mutex{},
		limit: limit,
		burst: burst,
	}
}

// Limiter serve caller get limiter by IP
func (i *IP) Limiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, ok := i.ips[ip]
	if !ok {
		i.mu.Unlock()
		return i.addIP(ip)
	}

	i.mu.Unlock()
	return limiter
}

func (i *IP) addIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.limit, i.burst)
	i.ips[ip] = limiter

	return limiter
}

var (
	ip *IP
)

// IPRateLimitMiddleware serve caller to added ip rate limit into gin
func IPRateLimitMiddleware(limit int, burst int) gin.HandlerFunc {
	if ip == nil {
		ip = NewIP(rate.Limit(limit), burst)
	}

	return func(c *gin.Context) {
		limiter := ip.Limiter(c.ClientIP())

		if !limiter.Allow() {
			c.Error(er.ErrRateLimit)
			return
		}

		c.Next()
	}
}
