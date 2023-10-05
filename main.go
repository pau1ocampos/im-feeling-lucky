package main

import (
	"github.com/pau1ocampos/im-feeling-lucky/cmd/imfeelinglucky/cli"
)

func main() {
	// opts := lucky.Options{
	// 	Client:      lucky.NewDefaultClient(),
	// 	FileSystem:  lucky.HdFileSystem{},
	// 	BaseUrl:     "https://www.euro-millions.com/",
	// 	FromYear:    2004,
	// 	UserAgent:   "im-feeling-lucky",
	// 	ShouldStore: true,
	// 	StoreFile:   "./results.json",
	// }

	// draws, err := opts.ScrapeFrom(context.TODO())
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(draws)
	cli.Execute()
}
