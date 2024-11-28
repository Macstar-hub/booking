package makearray

func makeArray(firstName string, lastName string) []string {
	booking := make([]string, 1)
	booking[0] = firstName + " " + lastName
	return booking
}
