package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/jukemori/go-cli-app/pkg"
)

var (
	entryID     string
	accessToken string
)

func main() {
	var rootCmd = &cobra.Command{Use: "my-contentful-cli"}

	var fetchCmd = &cobra.Command{
		Use:   "fetch",
		Short: "Fetch data from Contentful and save to the PostgreSQL database",
		Run:   fetchContentfulData,
	}

	rootCmd.AddCommand(fetchCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	fetchCmd.Flags().StringVarP(&entryID, "entry_id", "e", "", "Contentful entry ID")
	fetchCmd.Flags().StringVarP(&accessToken, "access_token", "a", "", "Contentful access token")
	fetchCmd.MarkFlagRequired("entry_id")
	fetchCmd.MarkFlagRequired("access_token")
}


func fetchContentfulData(cmd *cobra.Command, args []string) {
	// Get the values of entry_id and access_token from flags
	// (These values are now stored in the entryID and accessToken variables.)
	// entryID, _ := cmd.Flags().GetString("entry_id")
	// accessToken, _ := cmd.Flags().GetString("access_token")

	// Fetch data from Contentful API
	url := fmt.Sprintf("https://cdn.contentful.com/spaces/2vskphwbz4oc/entries/%s?access_token=%s", entryID, accessToken)
	response, err := pkg.FetchContentfulData(url)
	if err != nil {
		fmt.Println("Error fetching data from Contentful:", err)
		return
	}

	// Parse the Contentful response into a struct
	var contentfulResponse struct {
		Sys struct {
			ID        string `json:"id"`
			CreatedAt string `json:"createdAt"`
		} `json:"sys"`
		Fields struct {
			Name string `json:"name"`
		} `json:"fields"`
	}
	err = json.Unmarshal([]byte(response), &contentfulResponse)
	if err != nil {
		fmt.Println("Error parsing Contentful response:", err)
		return
	}

	// Extract the required data from the parsed response
	id := contentfulResponse.Sys.ID
	name := contentfulResponse.Fields.Name
	createdAt := contentfulResponse.Sys.CreatedAt

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

	err = pkg.SaveData(db, id, name, createdAt)
	if err != nil {
		fmt.Println("Error saving data to database:", err)
		return
	}

	fmt.Println("Data fetched from Contentful and saved to the PostgreSQL database.")
}
