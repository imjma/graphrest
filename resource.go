package graphrest

import (
	"net/url"
	"strings"
)

type Resource struct {
	Path   string
	Name   string
	ID     string
	Params url.Values
}

func NewResource(path string) *Resource {
	if len(path) == 0 {
		return nil
	}

	u, err := url.Parse(path)
	if err != nil {
		return nil
	}

	var r = &Resource{Path: path}

	// get resource name
	name, path := pathComponent(u.Path)
	r.Name = name

	if len(path) > 0 {
		// get resource id
		id, _ := pathComponent(path)
		r.ID = id
	}

	r.Params = u.Query()

	return r
}

// pathComponent returns the first part of path tand the remaining path
//
// Input: /comp1/comp2/comp3
// Output: comp1, comp2/comp3
func pathComponent(path string) (string, string) {
	// remove leading slash
	for len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	// get first part of path
	i := strings.IndexByte(path, '/')
	if i != -1 {
		return path[:i], path[i+1:]
	}

	return path, ""
}
