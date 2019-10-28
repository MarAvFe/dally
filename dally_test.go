package dally

import (
	"testing"
	"time"

	mocks "github.com/MarAvFe/dally/mocks"
	"github.com/stretchr/testify/assert"
)

// https://medium.com/@harrygogonis/testing-go-mocking-third-party-dependancies-4ab4e1c9bd3f
func TestInsertTime(t *testing.T) {
	mockThing := Thing{time.Now().Unix()}

	mockDAL := &mocks.DataAccessLayer{}
	mockDAL.On("Insert", "foo", mockThing).Return(nil)

	actual := InsertTime(mockDAL)

	mockDAL.AssertExpectations(t)

	assert.Equal(t, mockThing, actual, "should return a Thing")
}
