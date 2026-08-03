package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	redisStore "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware() gin.HandlerFunc {
	store, err := redisStore.NewStore(
		10,
		"tcp",
		os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"),
		"",
		"",
		[]byte(os.Getenv("SESSION_SECRET")),
	)

	if err != nil {
		log.Fatalf("Failed to create redis session store: %v", err)
	}

	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 7,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return sessions.Sessions("job-monitoring-session", store)

}
