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

	for i := len(content) - 1; i >= 0; i-- {
		name := stripFileExt(filepath.Base(content[i]))
		data := &ContentLinks{
			Href: hrefPrefix + name,
			Name: replaceDashWithColon(name),
		}
		cl = append(cl, data)
	}

	return cl, nil
}
