package LIS_test


import (
	"testing"

	"github.com/Ramshackle-Jamathon/go-LIS"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    input := []int{1, 2, 3, 4, 5}
    output := LIS.LIS(input)
    assert.Equal(t, output, []int{1, 2, 3, 4, 5})


    input = []int{5, 4, 3, 2, 1}
    output = LIS.LIS(input)
    assert.Equal(t, output, []int{1})


    input = []int{5, 4, 3, 2, 1}
    output = LIS.LIS(input)
    assert.Equal(t, output, []int{1})
}