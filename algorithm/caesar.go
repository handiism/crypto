package algorithm

// NewCaesar constructs a Caesar cipher.
func NewCaesar(shift int) *Affine {
	return &Affine{1, shift, 1}
}
