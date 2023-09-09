package main

import (
	"os"

	r "github.com/api-rest-go/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	router := gin.Default()

	app := r.Router(router)

	app.Run(port)
}

// type album struct {
// 	ID     string `json:"id"`
// 	Title  string `json:"title"`
// 	Artist string `json:"artist"`
// 	Year   int    `json:"year"`
// }

// var albums = []album{
// 	{ID: "1", Title: "Familia", Artist: "Camila Cabello", Year: 2022},
// 	{ID: "2", Title: "21", Artist: "Adele", Year: 2011},
// 	{ID: "3", Title: "The Eminem Show", Artist: "Eminem", Year: 2022},
// }

// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// func postAlbum(c *gin.Context) {
// 	var newAlbum album
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	newAlbum.ID = strconv.Itoa(len(albums) + 1)
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, albums)
// }

// func getById(c *gin.Context) {
// 	ID := c.Param("id")

// 	for _, a := range albums {
// 		if a.ID == ID {
// 			c.IndentedJSON(http.StatusAccepted, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found."})
// }

// func RemoveIndex(s []album, index int) []album {
// 	s[index] = s[len(s)-1]
// 	return s[:len(s)-1]
// }

// func deleteById(c *gin.Context) {
// 	ID := c.Param("id")

// 	for b, a := range albums {
// 		if a.ID == ID {
// 			albums = RemoveIndex(albums, b)
// 		}
// 	}
// 	c.IndentedJSON(http.StatusAccepted, albums)
// }
