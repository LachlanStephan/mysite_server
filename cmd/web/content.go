package main

import "path/filepath"

type ContentLinks struct {
	Href string
	Name string
}

func getContentLinks(contentParentPath, hrefPrefix string) ([]*ContentLinks, error) {
	content, err := getFilesFromPath(contentParentPath)
	if err != nil {
		return nil, err
	}

	cl := []*ContentLinks{}

	for _, page := range content {
		name := stripFileExt(filepath.Base(page))
		data := &ContentLinks{
			Href: hrefPrefix + name,
			Name: name,
		}
		cl = append(cl, data)
	}
	return cl, nil
}
