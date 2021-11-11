package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"github.com/iomesh/genqlient/graphql"
	"github.com/iomesh/genqlient/internal/testutil"
)

// ListInputQueryResponse is returned by ListInputQuery on success.
type ListInputQueryResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User ListInputQueryUser `json:"user"`
}

// GetUser returns ListInputQueryResponse.User, and is useful for accessing the field via an interface.
func (v *ListInputQueryResponse) GetUser() ListInputQueryUser { return v.User }

// ListInputQueryUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type ListInputQueryUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id testutil.ID `json:"id"`
}

// GetId returns ListInputQueryUser.Id, and is useful for accessing the field via an interface.
func (v *ListInputQueryUser) GetId() testutil.ID { return v.Id }

// __ListInputQueryInput is used internally by genqlient
type __ListInputQueryInput struct {
	Names []string `json:"names"`
}

// GetNames returns __ListInputQueryInput.Names, and is useful for accessing the field via an interface.
func (v *__ListInputQueryInput) GetNames() []string { return v.Names }

func ListInputQuery(
	client graphql.Client,
	names []string,
) (*ListInputQueryResponse, error) {
	__input := __ListInputQueryInput{
		Names: names,
	}
	var err error

	var retval ListInputQueryResponse
	err = client.MakeRequest(
		nil,
		"ListInputQuery",
		`
query ListInputQuery ($names: [String]) {
	user(query: {names:$names}) {
		id
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}
