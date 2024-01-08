package main

import (
	"path/filepath"
	"strings"
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

// can be more than one file ext eg. ".tmpl.html"
func stripFileExt(name string) string {
	r := name
	extCount := strings.Count(name, ".")
	for i := 0; i < extCount; i++ {
		r = strings.TrimSuffix(r, filepath.Ext(r))
	}
	return r
}
