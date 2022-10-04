package mod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInverse(t *testing.T) {
	for i := -50; i < 50; i++ {
		inv1, ok1 := Inverse(i, 26)
		inv2, ok2 := inverseNaive(i, 26)
		if ok1 != ok2 || inv1 != inv2 {
			t.Errorf("%d: inverse is %d %t, but naive is %d %t", i, inv1, ok1, inv2, ok2)
		}
	}
}

func inverseNaive(a, n int) (int, bool) {
	for i := 1; i < n; i++ {
		if Mod(a*i, n) == 1 {
			return i, true
		}
	}
	return 0, false
}

func TestMod(t *testing.T) {
	assert.Equal(t, 27%10, Mod(27, 10))
	assert.Equal(t, 13%10, Mod(13, 5))
	assert.Equal(t, 1, Mod(10, 3))
}
