package algorithm

// ROT13 is the ROT-13 symmetric cipher.
// For example, "ABCDEFGHIJKLM" becomes "NOPQRSTUVWXYZ".
var ROT13 = Affine{1, 13, 1}
