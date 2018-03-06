package httpx

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseQueryParametersEmptyQuery(t *testing.T) {
	url := "http://example.com/test"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	cut := &Request{Request: request}

	assert.True(t, cut.ParseQueryParameters())

	assert.Len(t, cut.QueryParams, 0)
}

func TestParseQueryParametersIgnoreEmptyQueryPair(t *testing.T) {
	url := "http://example.com/test?zalgo=he:comes&&waits=behind%20the%20walls&zalgo=slithy"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	cut := &Request{Request: request}

	assert.True(t, cut.ParseQueryParameters())

	assert.Len(t, cut.QueryParams, 2)
	for i, p := range cut.QueryParams {
		switch i {
		case 0:
			assert.Equal(t, i, p.Ordinal)
			assert.Len(t, p.Values, 2)
			assert.Equal(t, "zalgo", p.Name)
			assert.Equal(t, "he:comes", p.Values[0])
			assert.Equal(t, "slithy", p.Values[1])
			assert.NoError(t, p.Err)
		case 1:
			assert.Equal(t, i, p.Ordinal)
			assert.Len(t, p.Values, 1)
			assert.Equal(t, "waits", p.Name)
			assert.Equal(t, "behind the walls", p.Values[0])
			assert.NoError(t, p.Err)
		}
	}
}

func TestParseQueryParametersBadEscapeCodes(t *testing.T) {
	url := "http://example.com/test?name=foo%zzbar&sec%zzret=42"
	request := httptest.NewRequest(http.MethodGet, url, nil)
	cut := &Request{Request: request}

	assert.False(t, cut.ParseQueryParameters())

	assert.Len(t, cut.QueryParams, 2)
	for i, p := range cut.QueryParams {
		switch i {
		case 0:
			assert.Equal(t, i, p.Ordinal)
			assert.Len(t, p.Values, 1)
			assert.Equal(t, "name", p.Name)
			assert.Equal(t, "foo%zzbar", p.Values[0])
			assert.Error(t, p.Err)
		case 1:
			assert.Equal(t, i, p.Ordinal)
			assert.Len(t, p.Values, 1)
			assert.Equal(t, "sec%zzret", p.Name)
			assert.Equal(t, "42", p.Values[0])
			assert.Error(t, p.Err)
		}
	}
}
