package lib

import (
	"go-multipackageproject/assertions"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 25
	actual := Add(10, 15)

	assertions.AssertEqual(t, expected, actual)
}
