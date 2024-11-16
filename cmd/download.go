package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	arena "go.zakaria.org/aredotna"
)

func downloadBlock(b arena.ApiChannelBlock) error {
	class := b.Class

	switch class {
	case "Image":
		url := b.Image.Original.Url
		fname := strings.Split(path.Base(url), "?")[0]
		log.Printf("downloading image %q...", fname)

		outputf, err := os.Create(path.Join(archiveOpts.Dir, fname))
		if err != nil {
			return err
		}

		r, err := http.Get(url)
		if err != nil {
			return err
		}
		defer r.Body.Close()

		_, err = outputf.ReadFrom(r.Body)
		if err != nil {
			return err
		}
		log.Printf("downloaded %q.", fname)
	case "Text":
		content := b.Content
		fname := fmt.Sprintf("%d-%s.txt", b.GeneratedTitle)
		log.Printf("downloading text %q...", fname)

		outputf, err := os.Create(path.Join(archiveOpts.Dir, fname))
		if err != nil {
			return err
		}
		fmt.Fprintf(outputf, "%s", content)
		outputf.Close()
		log.Printf("downloaded %q.", fname)
	default:
		log.Printf("skipping %d: unsupported class %s.", b.Id, class)
	}
	return nil
}

func downloadBlocks(blks []arena.ApiChannelBlock) error {
	for _, blk := range blks {
		err := downloadBlock(blk)
		if err != nil {
			return err
		}
	}
	return nil
}

/*
func channelBlocks(slug string) []arena.ApiChannelBlock {
	//	var blks []arena.ApiChannelBlock
	ch, err := a.GetChannel(slug)
	if err != nil {
		log.Fatal(err)
	}

	iterations :=


	return ch.Contents
}*/
