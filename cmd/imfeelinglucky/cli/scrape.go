package cli

import (
	"context"
	"fmt"
	"log"

	"github.com/pau1ocampos/im-feeling-lucky/internal/app/lucky"
	"github.com/spf13/cobra"
)

var scrapeCmd = &cobra.Command{
	Use:     "scrape",
	Aliases: []string{},
	Short:   "Scrapes data from https://www.euro-millions.com website",
	Long: `Scrapes data from https://www.euro-millions.com website from a given year until current year.
If no year passed, 2004 will be the default.
You can choose to save the file on your local machine, otherwise results will be printed in an simple struct format`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		year, err := cmd.Flags().GetInt("year")

		if err != nil {
			log.Fatalf("An error occurred getting the year flag: %v\n", err)
		}

		fpath, err := cmd.Flags().GetString("file-path")

		if err != nil {
			log.Fatalf("An error occurred parsing the file-path flag: %v\n", err)
		}

		silent, err := cmd.Flags().GetBool("silent")
		if err != nil {
			log.Fatalf("An error occurred parsing the silent flag: %v\n", err)
		}

		shouldStore := false
		if fpath != "" {
			shouldStore = true
		}

		opts := lucky.Options{
			Client:      lucky.NewDefaultClient(),
			FileSystem:  lucky.HdFileSystem{},
			BaseUrl:     websiteUrl,
			FromYear:    year,
			UserAgent:   "im-feeling-lucky",
			ShouldStore: shouldStore,
			StoreFile:   fpath,
		}

		ctx, cancel := context.WithCancel(context.TODO())
		defer cancel()
		draws, err := opts.ScrapeFrom(ctx)
		if err != nil {
			log.Fatalf("An error occurred scraping %s: %v\n", websiteUrl, err)
		}

		if !silent {
			fmt.Println(draws)
		}
	},
}

func init() {
	f := scrapeCmd.Flags()
	f.IntP("year", "y", 2004, "The year the cli will scrape from until current year")
	f.StringP("file-path", "p", "", "The file path to store the data")
	f.BoolP("silent", "s", false, "Silent mode, will not print the structure with the scraped draws")
	rootCmd.AddCommand(scrapeCmd)
}
