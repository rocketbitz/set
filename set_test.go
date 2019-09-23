package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SetTestSuite struct {
	suite.Suite
}

func (s *SetTestSuite) TestSet() {
	testSet := New()

	assert.Equal(s.T(), 0, int(testSet.Len()))

	testSet.Add("test0")
	assert.Equal(s.T(), 1, int(testSet.Len()))

	testSet.Add("test1")
	assert.Equal(s.T(), 2, int(testSet.Len()))

	testSet.Add("test1")
	assert.Equal(s.T(), 2, int(testSet.Len()))

	slc := testSet.Slice()

	assert.Len(s.T(), slc, 2)

	assert.Equal(s.T(), "test0", slc[0])
	assert.Equal(s.T(), "test1", slc[1])

	assert.True(s.T(), testSet.Contains("test0"))
	assert.True(s.T(), testSet.Contains("test1"))
	assert.False(s.T(), testSet.Contains("test2"))

	testSet.Remove("test1")

	assert.Equal(s.T(), 1, int(testSet.Len()))
	assert.True(s.T(), testSet.Contains("test0"))
	assert.False(s.T(), testSet.Contains("test1"))
}

func TestSetTestSuite(t *testing.T) {
	suite.Run(t, new(SetTestSuite))
}
