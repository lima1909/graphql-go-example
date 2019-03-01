package list

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

const schema = `
schema {
	query: Query
}

type Query {
	users: [User]
}

type User {
	ID:     ID!
	Name:   String!
	Age:    Int
	IsRIch: Boolean
}
`

// Query schema-struct
type Query struct {
	users *[]*User // pointer slice of pointer user, because schema says slice and user can be null
}

// Schema (set Option: UseFieldResolvers)
func (*Query) Schema() *graphql.Schema {
	return graphql.MustParseSchema(schema, &Query{}, graphql.UseFieldResolvers())
}

// Users ...
func (Query) Users(ctx context.Context) (*[]*User, error) {

	return &[]*User{
		{
			IDField: "1",
			Name:    "Its me",
			Age:     toInt32Pointer(int32(42)),
		},
		{
			IDField: "2",
			Name:    "you",
			IsRich:  toBoolPointer(true),
		},
	}, nil
}

// User ...
type User struct {
	IDField string
	Name    string
	Age     *int32 // must be a pointer, because Schema allow Null-Values
	IsRich  *bool  // must be a pointer, because Schema allow Null-Values
}

// ID convert String to graphql-ID
func (u *User) ID() graphql.ID { return graphql.ID(u.IDField) }

// helper
func toInt32Pointer(v int32) *int32 { return &v }
func toBoolPointer(b bool) *bool    { return &b }
