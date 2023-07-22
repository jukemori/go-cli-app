package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/jukemori/go-cli-app/pkg"
)

var (
	entryID     string
	accessToken string
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch data from Contentful and save to the PostgreSQL database",
	Run:   fetchContentfulData,
}

func init() {
	fetchCmd.Flags().StringVar(&entryID, "entry_id", "", "Contentful entry ID")
	fetchCmd.Flags().StringVar(&accessToken, "access_token", "", "Contentful access token")
	fetchCmd.MarkFlagRequired("entry_id")
	fetchCmd.MarkFlagRequired("access_token")

	rootCmd.AddCommand(fetchCmd)
}

// fetchContentfulData is the function to fetch data from Contentful and save it to the PostgreSQL database.
// cmd/fetch.go

// fetchContentfulData is the function to fetch data from Contentful and save it to the PostgreSQL database.
func fetchContentfulData(cmd *cobra.Command, args []string) {
	// Fetch data from Contentful API
	url := fmt.Sprintf("https://cdn.contentful.com/spaces/2vskphwbz4oc/entries/%s?access_token=%s", entryID, accessToken)
	response, err := pkg.FetchContentfulData(url)
	if err != nil {
		fmt.Println("Error fetching data from Contentful:", err)
		return
	}

	// Display the fetched data
	fmt.Println("ID:", response.Sys.ID)
	fmt.Println("Name:", response.Fields.Name)
	fmt.Println("CreatedAt:", response.Sys.CreatedAt)

	// Save the fetched data to the PostgreSQL database.
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
