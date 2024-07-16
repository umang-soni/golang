package main

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// type CSVData struct {
// 	Data []map[string]string `json:"data"`
// }

type CSVDataValue struct {
	Box []map[string]string `json:"data"`
}

// func ApiCall() {

// 	url := "https://fantasy.premierleague.com/api//leagues-classic/8/standings"
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Printf("the error is %s:", err)
// 	}
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("the error in reading body data  of  response ")
// 	}
// 	defer resp.Body.Close()

// 	os.WriteFile("apis/Classic league standings/league_id=8.json", []byte(string(body)), 0644)
// }

func main() {

	// array := make([]int,2,5)
	// array1:=append(array,2,3,4)
	// fmt.Println(array)
	// fmt.Println(array1)

	// fmt.Println(array)
	// slice1:=array[1:3]
	// fmt.Println(slice1,len(slice1),cap(slice1))
	// slice1=slice1[1:]
	// fmt.Println(slice1)
	// fmt.Println(cap(slice1))

	r := gin.Default()
	r.Use(cors.Default())

	// Route handler for the POST request to handle CSV file upload
	r.POST("/upload", func(c *gin.Context) {
		// Retrieve the uploaded file from the request
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Open the uploaded file
		uploadedFile, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
			return
		}
		defer uploadedFile.Close()
		fmt.Println(uploadedFile)

		csvReader := csv.NewReader(uploadedFile)
		records, err := csvReader.ReadAll()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read CSV"})
			return
		}
		fmt.Println(records)

		// // Convert CSV data to JSON
		var csvdatavalue CSVDataValue
		var slice []string

		slice = append(slice, records[0]...)

		for _, record := range records[1:] {

			subBoxInstance := map[string]string{}

			for i, column := range record {
				subBoxInstance[slice[i]] = column

			}
			csvdatavalue.Box = append(csvdatavalue.Box, subBoxInstance)

		}

		// Return JSON response
		c.JSON(http.StatusOK, csvdatavalue.Box)
	})

	// Run the server on port 8080
	r.Run(":8080")
}
