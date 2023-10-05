package lucky

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type httpStub struct {
	Status     string
	StatusCode int
}

func (s *httpStub) Do(req *http.Request) (*http.Response, error) {
	responseBody := io.NopCloser(bytes.NewReader([]byte(`{"value":"fixed"}`)))
	return &http.Response{
		Status:     s.Status,
		StatusCode: s.StatusCode,
		Body:       responseBody,
	}, nil
}

func TestGetFullUrl_WithYear2020_ReturnsUrlForThatYear(t *testing.T) {
	// Arrange
	year := 2020
	opt := Options{
		BaseUrl: "abc/",
	}
	expected := "abc/" + resultResource + strconv.Itoa(year)
	// Act
	r := opt.getFullUrl(year)
	// Assert
	assert.Equal(t, expected, r)
}

func TestMakeRange_WithTwoConsecutiveYears_ReturnsTwoYears(t *testing.T) {
	// Arrange
	fromYear := 2001
	toYear := 2002
	expected := []int{2001, 2002}
	// Act
	r := makeRange(fromYear, toYear)
	// Assert
	assert.Len(t, r, 2)
	assert.Equal(t, expected, r)
}

func TestGetHtml_Response200OK_ReturnsResponse(t *testing.T) {
	// Arrage
	httpStub := &httpStub{
		Status:     "200 OK",
		StatusCode: 200,
	}
	opt := Options{
		Client:    httpStub,
		BaseUrl:   "abc",
		FromYear:  2004,
		UserAgent: "abc",
	}
	// Act
	r, err := opt.getHmtl(context.Background(), 2004)
	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestGetHtml_Response404OK_ReturnsErrorAndNilResponse(t *testing.T) {
	// Arrage
	httpStub := &httpStub{
		Status:     "404 NOT FOUND",
		StatusCode: 404,
	}
	opt := Options{
		Client:    httpStub,
		BaseUrl:   "abc",
		FromYear:  2004,
		UserAgent: "abc",
	}
	// Act
	r, err := opt.getHmtl(context.Background(), 2004)
	// Assert
	assert.Error(t, err)
	assert.Nil(t, r)
}
