package main

var data interface{} = "great"

func main() {
	str, ok := data.(string)
	if ok {
		println(str)
	} else {
		println("not a string")
	}
}
