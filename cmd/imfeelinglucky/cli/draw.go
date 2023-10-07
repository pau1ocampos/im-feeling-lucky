package cli

import (
	"context"
	"fmt"
	"log"

	"github.com/pau1ocampos/im-feeling-lucky/internal/app/lucky"
	"github.com/spf13/cobra"
)

var drawCmd = &cobra.Command{
	Use:     "draw",
	Aliases: []string{},
	Short:   "Draws a valid Euromillions key",
	Long: `Draws a valid Euromillions key. This draw is purely based on computer randoomization, it has no special algorithms supporting it.
Only smarteness is that it checks if the generated key was a past winning key. If it is, it gerenates a new on.
Since the odds of the generated key is one of the past winning keys is so low, it's possible to disable this behaviour. Check the flags documentation.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fpath, err := cmd.Flags().GetString("file-path")

		if err != nil {
			log.Fatalf("An error occurred parsing the file-path flag: %v\n", err)
		}

		dc, err := cmd.Flags().GetBool("disable-check")
		if err != nil {
			log.Fatalf("An error occurred parsing the silent flag: %v\n", err)
		}

		opts := lucky.Options{
			Client:      lucky.NewDefaultClient(),
			FileSystem:  lucky.HdFileSystem{},
			BaseUrl:     websiteUrl,
			UserAgent:   "im-feeling-lucky",
			ShouldStore: false,
			StoreFile:   fpath,
		}

		if dc {
			draws := lucky.Draws{}
			draw := draws.Generate(dc)
			printKey(draw)
			return
		}

		if fpath != "" {
			draws, err := opts.ParseFromFile(fpath)
			if err != nil {
				log.Fatalf("%v", err)
			}

			gDraw := draws.Generate(dc)
			printKey(gDraw)
			return
		}

		ctx, cancel := context.WithCancel(context.TODO())
		defer cancel()
		draws, err := opts.ScrapeFrom(ctx)
		if err != nil {
			log.Fatalf("%v", err)
		}

		gDraw := draws.Generate(dc)
		printKey(gDraw)
	},
}

func printKey(draw *lucky.Draw) {
	fmt.Print("🔢 ")
	for i, nr := range draw.Numbers {
		fmt.Printf("%d", nr)
		if i != len(draw.Numbers)-1 {
			fmt.Print("-")
		} else {
			fmt.Println()
		}
	}

	fmt.Print("🌟")
	for i, st := range draw.Starts {
		fmt.Printf("%d", st)
		if i != len(draw.Starts)-1 {
			fmt.Printf("-")
		} else {
			fmt.Println()
		}
	}
}

func init() {
	f := drawCmd.Flags()
	f.StringP("file-path", "p", "", "File path of the past draws")
	f.BoolP("disable-check", "", false, "Disables the ability of checking if the key already been drawn in the past. This check takes precedence over file-path")
	rootCmd.AddCommand(drawCmd)
}
