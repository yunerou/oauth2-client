package utils

import (
	"context"
	"net/url"

	"github.com/yunerou/oauth2-client/singleton"
)

func AddQueriesIntoURL(
	ctx context.Context,
	rawURL string,
	queries map[string]string,
) (newURLwithQueries string) {

	defer func() {
		singleton.GetLogger(ctx).
			Trace().
			Str("rawURL", rawURL).
			Interface("queries", queries).
			Str("newURLwithQueries", newURLwithQueries).
			Send()
	}()

	uurl, err := url.Parse(rawURL)
	if err != nil {
		singleton.GetLogger(ctx).Error().Err(err).Send()
		return rawURL
	}
	query := uurl.Query()
	for k, v := range queries {
		query.Add(k, v)
	}

	uurl.RawQuery = query.Encode()
	newURLwithQueries = uurl.String()
	return
}
