package jobs

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func GetJob() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

	}
}
