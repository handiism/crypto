package mod

// Mod computes the modulo x mod n in the interval [0, n). The result
// takes the sign of the divisor n whereas % takes the sign of the
// dividend x.
func Mod(x, n int) int {
	return (x%n + n) % n
}

// Inverse computes the multiplicative inverse of x (mod n). An inverse
// only exists when gcd(x, n) == 1.
func Inverse(x, n int) (inv int, ok bool) {
	t, t1 := 0, 1
	r, r1 := n, Mod(x, n)
	for r1 != 0 {
		quotient := r / r1
		t, t1 = t1, t-quotient*t1
		r, r1 = r1, r-quotient*r1
	}
	if r != 1 {
		return 0, false
	}
	if t < 0 {
		t += n
	}
	return t, true
}
