// // package controllers
package archive

// import (
// 	"github.com/fachi-r/go-crud/archive"
// 	"github.com/fachi-r/go-crud/database"

// 	// "github.com/fachi-r/go-crud/models"
// 	"github.com/gin-gonic/gin"
// )

// func CreatePosts(c *gin.Context) {
// 	// Get Data from POST request
// 	var body struct {
// 		Title string
// 		Body  string
// 	}

// 	c.Bind(&body)

// 	// Create a new post
// 	post := archive.Post{Title: body.Title, Body: body.Body}

// 	// Return the post
// 	result := database.DB.Create(&post)
// 	if result.Error != nil {
// 		c.Status(400)
// 		return
// 	}
// 	c.JSON(200, gin.H{
// 		"post": post,
// 	})
// }

// func GetPosts(c *gin.Context) {
// 	// Get Posts
// 	var posts []archive.Post
// 	database.DB.Find(&posts)

// 	// Respond with Posts
// 	c.JSON(200, gin.H{
// 		"posts": posts,
// 	})
// }

// func GetPost(c *gin.Context) {
// 	// Get ID from url
// 	id := c.Param("id")
// 	// Get Posts
// 	var post archive.Post
// 	database.DB.Find(&post, id)

// 	// Respond with Posts
// 	c.JSON(200, gin.H{
// 		"post": post,
// 	})
// }

// func UpdatePost(c *gin.Context) {
// 	// Get ID from url
// 	id := c.Param("id")
// 	// get post with ID
// 	var post archive.Post
// 	database.DB.Find(&post, id)

// 	// Get Data from request body
// 	var body struct {
// 		Title string
// 		Body  string
// 	}

// 	c.Bind(&body)

// 	// update post
// 	database.DB.Model(&post).Updates(archive.Post{Title: body.Title, Body: body.Body})

// 	// Respond with Posts
// 	c.JSON(200, gin.H{
// 		"post": post,
// 	})
// }

// func DeletePost(c *gin.Context) {
// 	// Get ID off URL
// 	id := c.Param("id")

// 	// Delete Post
// 	database.DB.Delete(&archive.Post{}, id)

// 	// Respond with status
// 	c.Status(200)
// }
