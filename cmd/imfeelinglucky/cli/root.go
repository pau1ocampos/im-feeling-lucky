package cli

import (
	"os"

	"github.com/spf13/cobra"
)

const websiteUrl string = "https://www.euro-millions.com"

var rootCmd = &cobra.Command{
	Use:   "lucky",
	Short: "lucky generates draws for Euromillions",
	Long:  "lucky generates draws for Euromillions, it scrapes data from https://www.euro-millions.com to scrape all draws and has the capability of not generate a draw that was already prized",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zero.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
