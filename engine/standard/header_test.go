package standard

import (
	"github.com/Hunter-Dolan/echo/engine/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHeader(t *testing.T) {
	header := &Header{http.Header{}}
	test.HeaderTest(t, header)

	header.reset(http.Header{})
	assert.Len(t, header.Keys(), 0)
}
