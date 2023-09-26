package lucky

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicates_WithDuplicates_ReturnsTrue(t *testing.T) {
	// Arrange
	dup := []int{1, 1}
	// Act
	r := duplicates(dup)
	// Assert
	assert.True(t, r)
}

func TestDuplicates_WithoutDuplicates_ReturnsFalse(t *testing.T) {
	// Arrange
	noDup := []int{1, 2, 3, 4, 5}
	// Act
	r := duplicates(noDup)
	// Assert
	assert.False(t, r)
}

func TestDealWithZeros_GivenSlice_ReturnsSliceWithAllElementsIncrementedOne(t *testing.T) {
	// Arrange
	slice := []int{1, 2, 3, 4}
	expected := []int{2, 3, 4, 5}
	// Act
	dealWithZeros(slice)
	// Assert
	assert.Equal(t, expected, slice)
}

func TestIsARepeatedDraw_WithRepeatedDraw_ReturnsTrue(t *testing.T) {
	// Arrange
	draws := Draws{
		Draw{
			Numbers: []int{1, 2, 3, 4, 5},
			Starts:  []int{1, 3},
		},
		Draw{
			Numbers: []int{10, 20, 30, 40, 41},
			Starts:  []int{10, 11},
		},
	}
	// Act
	r := draws.isARepeatedDraw(&Draw{
		Numbers: []int{10, 20, 30, 40, 41},
		Starts:  []int{10, 11},
	})

	// Assert
	assert.True(t, r)
}

func TestIsARepeatedDraw_WithDraw_ReturnsFalse(t *testing.T) {
	// Arrange
	draws := Draws{
		Draw{
			Numbers: []int{1, 2, 3, 4, 5},
			Starts:  []int{1, 3},
		},
		Draw{
			Numbers: []int{10, 20, 30, 40, 41},
			Starts:  []int{10, 11},
		},
	}
	// Act
	r := draws.isARepeatedDraw(&Draw{
		Numbers: []int{11, 20, 30, 40, 41},
		Starts:  []int{10, 11},
	})

	// Assert
	assert.False(t, r)
}

func TestGenerator_ResultHasFiveNumbersAndTwoStars(t *testing.T) {
	// Arrange
	draws := Draws{
		Draw{
			Numbers: []int{1, 2, 3, 4, 5},
			Starts:  []int{1, 3},
		},
		Draw{
			Numbers: []int{10, 20, 30, 40, 41},
			Starts:  []int{10, 11},
		},
	}
	// Act
	r := draws.Generate(false)
	// Assert
	assert.Len(t, r.Numbers, 5)
	assert.Len(t, r.Starts, 2)
}
