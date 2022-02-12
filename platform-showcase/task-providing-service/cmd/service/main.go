package main

func main() {
	println("Hello, I am task-providing-service!")

	// ineffectual assignment
	a := 1
	a = 2
	println(a)

	// dogsled
	_, _, _, _, err := dogsled()
	if err != nil {
		return
	}
}

func dogsled() (int, int, int, int, error) {
	return 0, 0, 0, 0, nil
}
