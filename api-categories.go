package kickkit

import (
	"context"
	"fmt"
	optionalvalues "github.com/glichtv/kick-kit/internal/optional-values"
	"net/http"
)

type Category struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
}

type Categories struct {
	client *Client
}

func (c *Client) Categories() Categories {
	return Categories{client: c}
}

type SearchCategoriesInput struct {
	Query string
}

// Search searches for Categories based on the search input.
//
// Reference: https://docs.kick.com/apis/categories#categories
func (c Categories) Search(ctx context.Context, input SearchCategoriesInput) (Response[[]Category], error) {
	const resource = "public/v1/categories"

	apiRequest := newAPIRequest[[]Category](
		ctx,
		c.client,
		requestOptions{
			resource: resource,
			method:   http.MethodGet,
			authType: AuthTypeUserToken,
			urlValues: optionalvalues.Values{
				"q": optionalvalues.Single(input.Query),
			},
		},
	)

	return apiRequest.execute()
}

type GetCategoryByIDInput struct {
	CategoryID int
}

// ByID retrieves Category based on it's ID.
//
// Reference: https://docs.kick.com/apis/categories#categories-category_id
func (c Categories) ByID(ctx context.Context, input GetCategoryByIDInput) (Response[Category], error) {
	const resource = "public/v1/categories"

	apiRequest := newAPIRequest[Category](
		ctx,
		c.client,
		requestOptions{
			resource: fmt.Sprintf("%s/%d", resource, input.CategoryID),
			method:   http.MethodGet,
			authType: AuthTypeUserToken,
		},
	)

	return apiRequest.execute()
}
