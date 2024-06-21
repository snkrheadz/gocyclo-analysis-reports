package main

import (
	"fmt"
)

// Simple function
func simpleFunction(a int) int {
	if a > 0 {
		return a + 1
	}
	return a - 1
}

// Moderate function
func moderateFunction(a int) int {
	if a > 0 {
		if a < 10 {
			return a + 1
		} else if a < 20 {
			return a + 2
		} else if a < 30 {
			return a + 3
		} else if a < 40 {
			return a + 4
		} else if a < 50 {
			return a + 5
		} else if a < 60 {
			return a + 6
		} else if a < 70 {
			return a + 7
		} else if a < 80 {
			return a + 8
		} else if a < 90 {
			return a + 9
		} else {
			return a + 10
		}
	} else {
		return a - 1
	}
}

// Complex function
func complexFunction(a int) int {
	switch {
	case a > 0:
		if a < 10 {
			return a + 1
		} else if a < 20 {
			return a + 2
		} else if a < 30 {
			return a + 3
		} else if a < 40 {
			return a + 4
		} else if a < 50 {
			return a + 5
		} else if a < 60 {
			return a + 6
		} else if a < 70 {
			return a + 7
		} else if a < 80 {
			return a + 8
		} else if a < 90 {
			return a + 9
		} else if a < 100 {
			return a + 10
		} else if a < 110 {
			return a + 11
		} else if a < 120 {
			return a + 12
		} else if a < 130 {
			return a + 13
		} else if a < 140 {
			return a + 14
		} else if a < 150 {
			return a + 15
		} else if a < 160 {
			return a + 16
		} else if a < 170 {
			return a + 17
		} else if a < 180 {
			return a + 18
		} else if a < 190 {
			return a + 19
		} else if a < 200 {
			return a + 20
		} else {
			return a + 21
		}
	default:
		return a - 1
	}
}

// Very complex function
func veryComplexFunction(a int) int {
	switch {
	case a > 0:
		if a < 10 {
			if a%2 == 0 {
				return a + 1
			} else {
				return a + 2
			}
		} else if a < 20 {
			if a%2 == 0 {
				return a + 3
			} else {
				return a + 4
			}
		} else if a < 30 {
			if a%2 == 0 {
				return a + 5
			} else {
				return a + 6
			}
		} else if a < 40 {
			if a%2 == 0 {
				return a + 7
			} else {
				return a + 8
			}
		} else if a < 50 {
			if a%2 == 0 {
				return a + 9
			} else {
				return a + 10
			}
		} else if a < 60 {
			if a%2 == 0 {
				return a + 11
			} else {
				return a + 12
			}
		} else if a < 70 {
			if a%2 == 0 {
				return a + 13
			} else {
				return a + 14
			}
		} else if a < 80 {
			if a%2 == 0 {
				return a + 15
			} else {
				return a + 16
			}
		} else if a < 90 {
			if a%2 == 0 {
				return a + 17
			} else {
				return a + 18
			}
		} else if a < 100 {
			if a%2 == 0 {
				return a + 19
			} else {
				return a + 20
			}
		} else if a < 110 {
			if a%2 == 0 {
				return a + 21
			} else {
				return a + 22
			}
		} else if a < 120 {
			if a%2 == 0 {
				return a + 23
			} else {
				return a + 24
			}
		} else if a < 130 {
			if a%2 == 0 {
				return a + 25
			} else {
				return a + 26
			}
		} else if a < 140 {
			if a%2 == 0 {
				return a + 27
			} else {
				return a + 28
			}
		} else if a < 150 {
			if a%2 == 0 {
				return a + 29
			} else {
				return a + 30
			}
		} else if a < 160 {
			if a%2 == 0 {
				return a + 31
			} else {
				return a + 32
			}
		} else if a < 170 {
			if a%2 == 0 {
				return a + 33
			} else {
				return a + 34
			}
		} else if a < 180 {
			if a%2 == 0 {
				return a + 35
			} else {
				return a + 36
			}
		} else if a < 190 {
			if a%2 == 0 {
				return a + 37
			} else {
				return a + 38
			}
		} else if a < 200 {
			if a%2 == 0 {
				return a + 39
			} else {
				return a + 40
			}
		} else {
			return a + 41
		}
	default:
		return a - 1
	}
}

func main() {
	fmt.Println("Simple Function:", simpleFunction(5))
	fmt.Println("Moderate Function:", moderateFunction(15))
	fmt.Println("Complex Function:", complexFunction(25))
	fmt.Println("Very Complex Function:", veryComplexFunction(35))
}
