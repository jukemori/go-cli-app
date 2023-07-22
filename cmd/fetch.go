// cmd/fetch.go
package cmd

import (
	// other imports...
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/jukemori/go-cli-app/pkg"
)

var (
	entryIDs = []string{
			"6QRk7gQYmOyJ1eMG9H4jbB", // 蜂蜜豆乳クランベリー
			"41RUO5w4oIpNuwaqHuSwEc", // 黒ゴマポテロール
			"4Li6w5uVbJNVXYVxWjWVoZ", // 黒七味と岩塩のフォカッチャ
	}

	accessToken string
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch data from Contentful and save to the PostgreSQL database",
	Run:   fetchContentfulData,
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(fetchCmd)
}

func initConfig() {
	err := godotenv.Load()
	if err != nil {
			fmt.Println("Error loading .env file:", err)
			os.Exit(1)
	}

	accessToken = os.Getenv("CONTENTFUL_ACCESS_TOKEN")
	if accessToken == "" {
			fmt.Println("Error: CONTENTFUL_ACCESS_TOKEN environment variable is not set.")
			os.Exit(1)
	}
}

func fetchContentfulData(cmd *cobra.Command, args []string) {
	for _, entryID := range entryIDs {
			url := fmt.Sprintf("https://cdn.contentful.com/spaces/2vskphwbz4oc/entries/%s?access_token=%s", entryID, accessToken)
			response, err := pkg.FetchContentfulData(url)
			if err != nil {
					fmt.Printf("Error fetching data for entry ID %s: %v\n", entryID, err)
					continue
			}

			// Display the fetched data
			fmt.Println("ID:", response.Sys.ID)
			fmt.Println("Name:", response.Fields.Name)
			fmt.Println("CreatedAt:", response.Sys.CreatedAt)

			db, err := pkg.OpenDatabase()
			if err != nil {
					fmt.Println("Error opening database:", err)
					return
			}
			defer db.Close()

			err = pkg.CreateTable(db)
			if err != nil {
					fmt.Println("Error creating table:", err)
					return
			}

			err = pkg.SaveData(db, response.Sys.ID, response.Fields.Name, response.Sys.CreatedAt)
			if err != nil {
					fmt.Println("Error saving data to database:", err)
					return
			}

			fmt.Println("Data fetched from Contentful and saved to the PostgreSQL database.")
	}
}
