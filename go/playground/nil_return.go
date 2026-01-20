package main

func GetNilMap() map[string]string {
	return nil
}

func isNil(thing any) bool {
	return thing == nil
}

func main() {
	result := GetNilMap()
	println(result == nil) // prints true
	println(isNil(result)) // prints false
}
