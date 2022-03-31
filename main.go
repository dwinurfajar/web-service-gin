//simple restful API golang + gin gonic

package main

import (
    "net/http"

    "github.com/gin-gonic/gin"//library gin gonic
)

//struct untuk album
type album struct {
    ID     string  `json:"id"`
    Judul  string  `json:"judul"`
    Artis string  `json:"artis"`
    Harga  float64 `json:"harga"`
}

//data awal album
var albums = []album{
    {ID: "1", Judul: "Palui", Artis: "iyaa", Harga: 56.99},
    {ID: "2", Judul: "Amang", Artis: "Iyaaa", Harga: 17.99},
    {ID: "3", Judul: "Acil", Artis: "Iyaaaa", Harga: 39.99},
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)//route untuk mengambil semua data
	router.GET("/albums/:id", getAlbumByID)//route untuk mengambil data spesifik
	router.POST("/albums", postAlbums)//route untuk menyimpan/menambah data
    router.Run("localhost:8080")
}

//fungsi get album
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

//fungsi post album
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Camemanggil BindJSON untuk bind JSON yang diterima ke newAlbum
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // tambah newAlbum ke albums.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

//fungsi get album dengan spesifik data dengan selektor ID
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // proses pencocokan data yang di dari dengan di album
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}