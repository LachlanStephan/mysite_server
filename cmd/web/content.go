package main

import (
	"path/filepath"
)

type ContentLinks struct {
	Href string
	Name string
}

func getContentLinks(contentParentPath, hrefPrefix string, isHomePage bool) ([]*ContentLinks, error) {
	content, err := getFilesFromPath(contentParentPath)
	if err != nil {
		return nil, err
	}

	cl := []*ContentLinks{}
	if !isHomePage {
		cl = append(cl, getHomeLink())
	}

	// reverse order for newest fist
	for i := len(content) - 1; i >= 0; i-- {
		name := stripFileExt(filepath.Base(content[i]))
		data := &ContentLinks{
			Href: hrefPrefix + name,
			Name: stringReplace(name, "-", ": ", 1),
		}
		cl = append(cl, data)
	}

	return cl, nil
}

func getHomeLink() *ContentLinks {
	link := ContentLinks{
		Href: "",
		Name: "Home",
	}

	return &link
}
