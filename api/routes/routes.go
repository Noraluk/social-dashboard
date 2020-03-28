package routes

import (
	"social-dashboard/api/controllers/user"

	"github.com/labstack/echo"
)

func Run() {
	e := echo.New()

	// Routes
	e.GET("/users/daily/message", user.GetMessagePerDay)
	e.GET("/users/accounts/10-message", user.Get10AccountByMessage)
	e.GET("/users/messages/10-engagement", user.Get10MessageByEngagement)
	e.GET("/users/messages/word-clouds", user.GetWordClouds)
	e.GET("/users/messages/hashtag-clouds", user.GetHashtagClouds)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
