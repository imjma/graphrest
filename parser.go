package graphrest

import (
	"strings"

	"github.com/imjma/graphqb"
)

// Parse converts Resource to GraphQL query
func Parse(r *Resource) string {
	query := graphqb.NewQuery("query")

	f := graphqb.NewField(r.Name)

	if len(r.ID) > 0 {
		f.SetArguments(graphqb.NewArgument("id", r.ID))
	}

	// fields
	fields := strings.Split(r.Params.Get("field"), ",")
	for _, v := range fields {
		f.AddField(graphqb.NewField(v))
	}

	query.SetFields(f)
	// parse Params
	return query.Stringify()
}
