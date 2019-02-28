package hello

import graphql "github.com/graph-gophers/graphql-go"

const schema = `
schema {
    query: Query
}

type Query {
    hello: String!
}
`

// Query ...
type Query struct{}

// Hello ...
func (*Query) Hello() string { return "Hello" }

// Schema ...
func (*Query) Schema() *graphql.Schema {
	return graphql.MustParseSchema(schema, &Query{})
}
