package api

var temps = []int{5, 10, 3, 25, 45, 80, 90}

// Read returns a slice of elements from the temps slice.
func Read(start, stop int) []int {
	// ----------------------------------------
	// RESTRICTIONS — ONLY ADD YOUR CODE IN THIS AREA
	// clone := make([]int, len(temps))
	// copy(clone, temps)

	portion := temps[start:stop:stop]

	// RESTRICTIONS — ONLY ADD YOUR CODE IN THIS AREA
	// ----------------------------------------

	return portion
}

// All returns the temps slice
func All() []int {
	return temps
}
