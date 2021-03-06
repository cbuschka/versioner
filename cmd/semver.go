package main

import (
	"fmt"
	"github.com/cbuschka/semver/internal/git"
	"github.com/cbuschka/semver/internal/version"
	"log"
)

func main() {
	tags, err := git.ListTags()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Tags found: %v\n", tags)

	versions, removedTags, err := version.FilterConvertAndSortAsc(tags)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Invalid tags skipped: %v\n", removedTags)
	fmt.Printf("Sorted versions: %v\n", versions)
}
