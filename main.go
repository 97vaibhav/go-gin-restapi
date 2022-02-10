package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type tweet struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	User  string `json:"user"`
}

var tweets = []tweet{
	{ID: "1", Title: "anpad tweet", User: "daisuke"},
	{ID: "2", Title: "hello new tweet", User: "vaibhav"},
	{ID: "3", Title: "hey whats up", User: "Sankalp"},
}

func getTweets(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tweets)

}

func postTweets(c *gin.Context) {
	var newTweet tweet
	err := c.BindJSON(&newTweet)

	if err != nil {
		return
	}

	// Add the new tweet to the slice.
	tweets = append(tweets, newTweet)
	c.IndentedJSON(http.StatusCreated, newTweet)
}

func getTweetByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of tweets, looking for
	// an tweet whose ID value matches the parameter.
	for _, a := range tweets {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "tweet not found"})
}

func main() {
	router := gin.Default()
	router.GET("/tweets", getTweets)
	router.POST("/tweets", postTweets)
	router.GET("/tweets/:id", getTweetByID)

	router.Run("localhost:8080")
}
