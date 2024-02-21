package main

import (
	"path/filepath"
)

type BlogLinks struct {
	Href     string
	BlogName string
}

func getBlogLinks() ([]*BlogLinks, error) {
	blogs, err := getFilesFromPath(blogsPath)
	if err != nil {
		return nil, err
	}

	bl := []*BlogLinks{}

	for _, page := range blogs {
		name := stripFileExt(filepath.Base(page))
		data := &BlogLinks{
			Href:     "blog/view/" + name,
			BlogName: name,
		}
		bl = append(bl, data)
	}
	return bl, nil
}
