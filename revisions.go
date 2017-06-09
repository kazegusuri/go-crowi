package crowi

import (
	"context"
	"net/http"
	"net/url"
)

// RevisionsService handles communication with the Revisions related
// methods of the Crowi API.
type RevisionsService service

// RevisionsListOptions specifies the optional parameters to the
// RevisionsService.List methods.
type RevisionsListOptions struct {
	ListOptions
}

type RevisionsListResponse struct {
	Revisions []PageRevision `json:"-"`
}

// List shows the list of pages. A parameter of path or user is required.
func (s *RevisionsService) List(ctx context.Context, pageID string) (*RevisionsListResponse, error) {
	var resp RevisionsListResponse
	params := url.Values{}
	params.Set("access_token", s.client.config.Token)
	params.Set("page_id", pageID)
	err := s.client.newRequest(ctx, http.MethodGet, "/_api/revisions.list", params, &resp.Revisions)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
