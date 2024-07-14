package main

import (
	"fmt"
	"io"
	"os"

	"net/http"
)

// type CSVData struct {
// 	Data []map[string]string `json:"data"`
// }

type CSVDataValue struct {
	Box [][]SubBox `json:"box"`
}

type SubBox struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

func ApiCall() {

	url := "https://fantasy.premierleague.com/api//leagues-classic/8/standings"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("the error is %s:", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("the error in reading body data  of  response ")
	}
	defer resp.Body.Close()
	
	os.WriteFile("apis/Classic league standings/league_id=8.json", []byte(string(body)), 0644)
}

func main() {
	ApiCall()

	// Initialize Gin router with default middleware
	// r := gin.Default()

	// // Route handler for the POST request to handle CSV file upload
	// r.POST("/upload", func(c *gin.Context) {
	// 	// Retrieve the uploaded file from the request
	// 	file, err := c.FormFile("file")
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	// Open the uploaded file
	// 	uploadedFile, err := file.Open()
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
	// 		return
	// 	}
	// 	defer uploadedFile.Close()
	// 	fmt.Println(uploadedFile)

	// 	csvReader := csv.NewReader(uploadedFile)
	// 	records, err := csvReader.ReadAll()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read CSV"})
	// 		return
	// 	}
	// 	fmt.Println(records)

	// 	// // Convert CSV data to JSON
	// 	var csvdatavalue CSVDataValue
	// 	var slice []string
	// 	for _, value := range records[0] {
	// 		slice = append(slice, value)
	// 	}

	// 	for _, record := range records[1:] {

	// 		var subBoxInstance []SubBox

	// 		for i, column := range record {
	// 			subBoxInstance = append(subBoxInstance, SubBox{Key: slice[i], Value: column})

	// 		}
	// 		csvdatavalue.Box = append(csvdatavalue.Box, subBoxInstance)

	// 	}

	// 	// Return JSON response
	// 	c.JSON(http.StatusOK, csvdatavalue.Box)
	// })

	// // Run the server on port 8080
	// r.Run(":8080")
}
