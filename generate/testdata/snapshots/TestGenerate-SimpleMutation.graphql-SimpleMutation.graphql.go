package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"github.com/iomesh/genqlient/graphql"
	"github.com/iomesh/genqlient/internal/testutil"
)

// SimpleMutationCreateUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type SimpleMutationCreateUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// GetId returns SimpleMutationCreateUser.Id, and is useful for accessing the field via an interface.
func (v *SimpleMutationCreateUser) GetId() testutil.ID { return v.Id }

// GetName returns SimpleMutationCreateUser.Name, and is useful for accessing the field via an interface.
func (v *SimpleMutationCreateUser) GetName() string { return v.Name }

// SimpleMutationResponse is returned by SimpleMutation on success.
type SimpleMutationResponse struct {
	CreateUser SimpleMutationCreateUser `json:"createUser"`
}

// GetCreateUser returns SimpleMutationResponse.CreateUser, and is useful for accessing the field via an interface.
func (v *SimpleMutationResponse) GetCreateUser() SimpleMutationCreateUser { return v.CreateUser }

// __SimpleMutationInput is used internally by genqlient
type __SimpleMutationInput struct {
	Name string `json:"name"`
}

// GetName returns __SimpleMutationInput.Name, and is useful for accessing the field via an interface.
func (v *__SimpleMutationInput) GetName() string { return v.Name }

// SimpleMutation creates a user.
//
// It has a long doc-comment, to test that we handle that correctly.
// What a long comment indeed.
func SimpleMutation(
	client graphql.Client,
	name string,
) (*SimpleMutationResponse, error) {
	__input := __SimpleMutationInput{
		Name: name,
	}
	var err error

	var retval SimpleMutationResponse
	err = client.MakeRequest(
		nil,
		"SimpleMutation",
		`
mutation SimpleMutation ($name: String!) {
	createUser(name: $name) {
		id
		name
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}
