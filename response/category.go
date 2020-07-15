package response

import (
	"encoding/json"
	"net/url"
)

type CategoryResponse struct {
	HREF  string          `json:"href"`
	Icons []ImageResponse `json:"icons"`
	ID    string          `json:"id"`
	Name  string          `json:"name"`
}

func NewCategoryResponse(data []byte) (*CategoryResponse, error) {
	category := &CategoryResponse{}
	err := json.Unmarshal(data, &category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

type Category struct {
	href  url.URL
	icons []Image
	id    string
	name  string
}

type CategoryPageResponse struct{}

type CategoryPage struct{}
