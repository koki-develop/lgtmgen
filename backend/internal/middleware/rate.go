package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koki-develop/lgtmgen/backend/internal/log"
	"github.com/koki-develop/lgtmgen/backend/internal/repo"
	"github.com/koki-develop/lgtmgen/backend/internal/service"
	"github.com/koki-develop/lgtmgen/backend/internal/util"
)

type RateLimitter struct {
	repo *repo.Repository
}

func NewRateLimitter(r *repo.Repository) *RateLimitter {
	return &RateLimitter{
		repo: r,
	}
}

func (m *RateLimitter) Apply(tier string, limit int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := util.GetClientIPFromContext(ctx)
		rate, err := m.repo.FindRate(ctx, ip, tier)
		if err != nil {
			log.Error(ctx, "failed to find rate", err)
		} else {
			if rate != nil && rate.LimitReached(limit) {
				log.Info(ctx, "rate limit reached", "rate", rate)
				ctx.JSON(http.StatusTooManyRequests, gin.H{"code": service.ErrCodeRateLimitReached})
				ctx.Abort()
				return
			}
			if err := m.repo.IncrementRate(ctx, ip, tier); err != nil {
				log.Error(ctx, "failed to increment rate", err)
			}
		}

		ctx.Next()
	}
}
