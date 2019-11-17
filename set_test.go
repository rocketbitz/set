package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SetTestSuite struct {
	suite.Suite
}

func (s *SetTestSuite) TestSafe() {
	testSet := New()

	assert.Equal(s.T(), 0, testSet.Len())

	testSet.Add("test0")
	assert.Equal(s.T(), 1, testSet.Len())

	testSet.Add("test1")
	assert.Equal(s.T(), 2, testSet.Len())

	testSet.Add("test1")
	assert.Equal(s.T(), 2, testSet.Len())
	assert.Equal(s.T(), "test1", testSet.At(1).(string))
	assert.Nil(s.T(), testSet.At(2))

	slc := testSet.Slice()

	assert.Len(s.T(), slc, 2)

	assert.Equal(s.T(), "test0", slc[0])
	assert.Equal(s.T(), "test1", slc[1])

	assert.True(s.T(), testSet.Contains("test0"))
	assert.True(s.T(), testSet.Contains("test1"))
	assert.False(s.T(), testSet.Contains("test2"))

	testSet.Remove("test1")

	assert.Equal(s.T(), 1, testSet.Len())
	assert.True(s.T(), testSet.Contains("test0"))
	assert.False(s.T(), testSet.Contains("test1"))

	testSet = NewFromSlice([]interface{}{"one", "two"})
	assert.Equal(s.T(), 2, testSet.Len())
	assert.True(s.T(), testSet.Replace("two", "three"))
	assert.Equal(s.T(), "three", testSet.Slice()[1])
	assert.Equal(s.T(), 1, testSet.Index("three"))
	assert.Equal(s.T(), -1, testSet.Index("zero"))
}

func (s *SetTestSuite) TestUnsafe() {
	testSet := NewUnsafe()

	assert.Equal(s.T(), 0, testSet.Len())

	testSet.Add("test0")
	assert.Equal(s.T(), 1, testSet.Len())

	testSet.Add("test1")
	assert.Equal(s.T(), 2, testSet.Len())

	testSet.Add("test1")
	assert.Equal(s.T(), 2, testSet.Len())
	assert.Equal(s.T(), "test1", testSet.At(1).(string))
	assert.Nil(s.T(), testSet.At(2))

	slc := testSet.Slice()

	assert.Len(s.T(), slc, 2)

	assert.Equal(s.T(), "test0", slc[0])
	assert.Equal(s.T(), "test1", slc[1])

	assert.True(s.T(), testSet.Contains("test0"))
	assert.True(s.T(), testSet.Contains("test1"))
	assert.False(s.T(), testSet.Contains("test2"))

	testSet.Remove("test1")

	assert.Equal(s.T(), 1, testSet.Len())
	assert.True(s.T(), testSet.Contains("test0"))
	assert.False(s.T(), testSet.Contains("test1"))

	testSet = NewUnsafeFromSlice([]interface{}{"one", "two"})
	assert.Equal(s.T(), 2, testSet.Len())
	assert.True(s.T(), testSet.Replace("two", "three"))
	assert.Equal(s.T(), "three", testSet.Slice()[1])
	assert.Equal(s.T(), 1, testSet.Index("three"))
	assert.Equal(s.T(), -1, testSet.Index("zero"))
}

func TestSetTestSuite(t *testing.T) {
	suite.Run(t, new(SetTestSuite))
}
