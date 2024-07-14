package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func ApiCall() {

	url := "https://fantasy.premierleague.com/api/element-summary/57/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("the error is %s:", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("the error in reading body data  of  response ")
	}
	defer resp.Body.Close()
	fmt.Println(string(body))
	os.WriteFile("apis/element-summary-57.txt", []byte(string(body)), 0644)

	// 	// http.HandleFunc("/",handler)
	// 	// http.ListenAndServe(":8080",nil)
	// 	// r := gin.Default()
	// 	// r.GET("/", func(c *gin.Context) {
	// 	// 	c.JSON(http.StatusOK, gin.H{
	// 	// 		"message": "Hello, Gin!",
	// 	// 	})
	// 	// })

	// 	// r.Run(":8080")

	// 	r := gin.Default();
	// 	r.POST("/upload", func(c *gin.Context) {

	//         if err != nil {
	//             c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//             return
	//         }

	//         uploadedFile, err := file.Open()
	//         if err != nil {
	//             c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
	//             return
	//         }
	//         defer uploadedFile.Close()

	//         csvReader := csv.NewReader(uploadedFile)
	//         records, err := csvReader.ReadAll()
	//         if err != nil {
	//             c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read CSV"})
	//             return
	//         }

	//         var jsonData CSVData
	//         jsonData.Data = make([]map[string]string, 0)

	//         for _, record := range records {
	//             data := make(map[string]string)
	//             for i, column := range record {
	//                 data[csvReader.Read()[i]] = column
	//             }
	//             jsonData.Data = append(jsonData.Data, data)
	//         }

	//         c.JSON(http.StatusOK, jsonData)
	//     })

	//         // Return JSON response

	// r.Run(":8080")
}

// input
// for loop
//while loop
