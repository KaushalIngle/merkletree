package merkletree

// indent returns a string of "-" characters of length depth.
func indent(depth int) string {
	var s string
	for i := 0; i < depth; i++ {
		s += "-"
	}
	return s
}
