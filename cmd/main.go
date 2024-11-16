package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	arena "go.zakaria.org/aredotna"
	"go.zakaria.org/subc"
)

var a *arena.Arena
var archiveOpts struct {
	Slug string
	Dir  string
}

func main() {
	a = arena.New(os.Getenv("ARENA_TOKEN"))

	subc.Sub("archive").StringVar(&archiveOpts.Slug, "slug", "", "slug of channel to download")
	subc.Sub("archive").StringVar(&archiveOpts.Dir, "d", ".", "output directory")
	subcommand, err := subc.Parse()
	if err != nil {
		log.Fatal(err)
	}

	switch subcommand {
	case "archive":
		log.Printf("slug:%q\n", archiveOpts.Slug)
		//blks := channelBlocks(archiveOpts.Slug)
		blks := GetChannelContents(archiveOpts.Slug)
		//err := downloadBlocks(blks)
		//if err != nil {
		//	log.Fatal(err)
		//}
		for _, b := range blks {
			u, _ := url.Parse(b.Image.Original.Url)
			u.RawQuery = ""
			fmt.Printf("%s\n", u.String())
		}
		////for _, b := range blks {
		////	u, _ := url.Parse(b.Image.Original.Url)
		////	u.RawQuery = ""
		////	fmt.Printf("%s\n", u.String())
		////}
	default:
		log.Fatalf("unknown subcommand %s", os.Args[1])
	}

}
