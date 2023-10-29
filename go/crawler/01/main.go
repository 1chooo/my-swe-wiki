package main

import (
    "encoding/csv"
    "fmt"
    "net/http"
    "os"
    "strconv"
    "strings"
)

func main() {
    csvFile, err := os.Create("../csv/number_of_films.csv")
    if err != nil {
        fmt.Println("Error creating CSV file:", err)
        return
    }
    defer csvFile.Close()

    writer := csv.NewWriter(csvFile)
    defer writer.Flush()

    data := map[string]string{
        "realm":            "title",
        "title_type":       "feature",
        "release_date-min": "2015-02",
        "release_date-max": "2015-02",
        "view":             "detailed",
        "count":            "50",
        "sort":             "moviemeter,asc",
    }

    for i := 2015; i <= 2020; i++ {
        for j := 1; j <= 12; j++ {
            date := fmt.Sprintf("%d-%02d", i, j)
            data["release_date-min"] = date
            data["release_date-max"] = date

            res, err := http.PostForm("https://www.imdb.com/search/title/", data)
            if err != nil {
                fmt.Println("HTTP POST error:", err)
                return
            }
            defer res.Body.Close()

            // Read response body
            buf := make([]byte, 8192)
            var builder strings.Builder
            for {
                n, err := res.Body.Read(buf)
                if n == 0 || err != nil {
                    break
                }
                builder.Write(buf[:n])
            }

            // Extract number of films
            responseText := builder.String()
            numStart := strings.Index(responseText, `<div class="desc">`)
            if numStart == -1 {
                fmt.Println("Number not found in response.")
                return
            }
            numEnd := strings.Index(responseText[numStart:], " titles.</span>")
            if numEnd == -1 {
                fmt.Println("Number not found in response.")
                return
            }
            numText := responseText[numStart:numStart+numEnd]
            numWords := strings.Fields(numText)
            if len(numWords) > 0 {
                numStr := numWords[len(numWords)-1]
                numStr = strings.ReplaceAll(numStr, ",", "")
                num, err := strconv.Atoi(numStr)
                if err != nil {
                    fmt.Println("Error converting number:", err)
                    return
                }
                writer.Write([]string{date, strconv.Itoa(num)})
            }
        }
    }
}
