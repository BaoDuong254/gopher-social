package store

import (
	"net/http"
	"strconv"
)

type PaginatedFeedQuery struct {
	Limit  int    `json:"limit" validate:"gte=1,lte=20"`
	Offset int    `json:"offset" validate:"gte=0"`
	Sort   string `json:"sort" validate:"oneof=asc desc"`
}

func (fq PaginatedFeedQuery) Parse(r *http.Request) (PaginatedFeedQuery, error) {
	qs := r.URL.Query()
	limitStr := qs.Get("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			return fq, nil
		}
		fq.Limit = limit
	}
	offsetStr := qs.Get("offset")
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return fq, nil
		}
		fq.Offset = offset
	}
	sort := qs.Get("sort")
	if sort != "" {
		fq.Sort = sort
	}
	return fq, nil
}
