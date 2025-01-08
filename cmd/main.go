package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func getBody(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make GET request: %v", err)
	}
	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	// Read the body into a byte slice
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	return body
}

type Article struct {
	Area    string   `json:"area"`
	Source  string   `json:"source"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Journal string   `json:"journal"`
	Vip     string   `json:"vip"`
	Doi     string   `json:"doi"`
	Jif     float32  `json:"jif"`
	Ahi     float32  `json:"ahi"`
	Agi     float32  `json:"agi"`
	Soc     float32  `json:"soc"`
	Alt     float32  `json:"alt"`
	Sw      float32  `json:"sw"`
	Oi      float32  `json:"oi"`
}

type Articles struct {
	Articles []Article `json:"articles"`
}

func main() {

	rootDir, err := filepath.Abs(".env") // Это путь относительно текущей директории
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}
	fmt.Println(rootDir)

	err = godotenv.Load(rootDir)
	if err != nil {
		log.Fatalf("Error loading .env file: %e", err)
	}

	var apiUrl string = os.Getenv("API_URL")

	var rootCmd = &cobra.Command{
		Use:   "paperspls",
		Short: "CLI interface for retrieving best scientific papers of the week.",
		Long: `
		This command-line interface (CLI) helps you stay up-to-date with cutting-edge scientific research in your area of interest with minimal effort.
		Powered by an algorithm that analyzes tens of thousands of papers each week, PapersPlease helps you 
		discover the best research based on a variety of metrics. You'll always be in the know, effortlessly.
		Repo: https://github.com/ZeusKronid/PapersPlease
		`,
		Run: func(cmd *cobra.Command, args []string) {
			var getAll []byte = getBody(apiUrl + "/articles")

			articles := new(Articles)

			err := json.Unmarshal(getAll, &articles)
			if err != nil {
				log.Fatalf("Error unmarshalling JSON: %v", err)
			}
			for _, article := range articles.Articles {
				fmt.Printf("Title: %s; oi: %f, source: %s \n", article.Title, article.Oi, article.Source)

			}
		},
	}

	//rootCmd.AddCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
