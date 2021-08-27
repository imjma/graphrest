package graphrest

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	v := url.Values{}
	v.Set("field", "title,text")
	r := &Resource{
		Path:   "/foo/123?field=title,text",
		Name:   "foo",
		ID:     "123",
		Params: v,
	}

	q := Parse(r)
	assert.Equal(t, "query{foo(id:\"123\"){title,text}}", q)
}

func TestParseWithoutID(t *testing.T) {
	v := url.Values{}
	v.Set("field", "title,text")
	r := &Resource{
		Path:   "/foo?field=title,text",
		Name:   "foo",
		Params: v,
	}

	q := Parse(r)
	assert.Equal(t, "query{foo{title,text}}", q)
}
