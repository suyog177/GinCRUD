package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

// Employee struct containing only 2 fields for now
type Employee struct {
	EmpCode int    `json:"empCode" pg:"empCode,pk"`
	EmpName string `json:"empName" pg:"empName"`
}
//global variable for accessing the postgre
var dbConnect *pg.DB

//db configs and starting connection 
func DBConnect() *pg.DB {
	conf := &pg.Options{
		User:     "admin",
		Password: "123456",
		Addr:     "localhost:3352",
		Database: "local",
	}

	var db *pg.DB = pg.Connect(conf)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	InitiateDB(db)
	return db
}

//assigning the connected db to the globally declared db variable
func InitiateDB(db *pg.DB) {
	dbConnect = db
}

//main function handling all the routes
func main() {
	DBConnect()
	r := gin.Default()
	r.GET("/", GetString)
	r.GET("/:name", GetName)
	r.GET("/getAllEmployee", GetEmployees)
	r.POST("/registerEmployee", AddEmployee)
	r.PUT("/updateName/:empCode", UpdateDetails)
	r.DELETE("/removeEmployee/:empCode", DeleteDetails)

	r.Run()
}

//function to return hello world with the string received from URL
func GetName(c *gin.Context) {
	name := c.Param("name")
	msg := fmt.Sprintf("Hello world, %s", name)
	c.String(200, msg)
}

//generic hello world function
func GetString(c *gin.Context) {

	c.String(200, "Hello world")
}

//all below listed are CRUD functions 
func GetEmployees(c *gin.Context) {
	var emp []Employee
	err := dbConnect.Model(&emp).Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   emp,
	})
}

func AddEmployee(c *gin.Context) {
	user := Employee{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	code := user.EmpCode
	name := user.EmpName

	err := dbConnect.Insert(&Employee{
		EmpCode: code,
		EmpName: name,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   user,
	})

}

func UpdateDetails(c *gin.Context) {
	code, _ := strconv.Atoi(c.Param("empCode"))
	user := Employee{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	name := user.EmpName
	_, err := dbConnect.Model(&Employee{}).Set("empName = ?", name).Where("empCode = ?", code).Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   user,
	})
}

func DeleteDetails(c *gin.Context) {
	code, _ := strconv.Atoi(c.Param("empCode"))
	user := &Employee{EmpCode: code}

	err := dbConnect.Delete(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   user,
	})
}
