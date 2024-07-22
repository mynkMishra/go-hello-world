package main

import (
    "net/http"
	"github.com/google/uuid"
    "github.com/gin-gonic/gin"
	"strings"
	"fmt"
)

type Record struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Age int64 `json:"age"`
}

var records []Record

func sayHello(c *gin.Context){
	c.String(http.StatusOK, "Hello Mayank !")
}

func saveRecord(c *gin.Context){
	var record Record
        if err := c.BindJSON(&record); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        id := uuid.New().String()
        record.ID = id

        records = append(records, record)

        c.JSON(http.StatusOK, gin.H{"id": id})
}

func getRecord(c *gin.Context){
	id := c.Param("id")
	fmt.Println(id);
        for _, p := range records {
            if strings.EqualFold(p.ID, id) {
                c.JSON(http.StatusOK, p)
                return
            }
        }

        c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
}

func getRecords(c *gin.Context){
	c.JSON(http.StatusOK, records);
}

func main(){
	router := gin.Default()
	router.GET("/", sayHello)
	router.GET("/record/:id", getRecord)
	router.POST("/record", saveRecord)
	router.GET("/records", getRecords)

	router.Run("localhost:8080")
}