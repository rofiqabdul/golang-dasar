package helper

var version = "1.0.0"      // tidak bisa diakses dari luar package
var Application = "golang" // bisa diakses dari luar package

func sayGoodbye(name string) string { // tidak bisa diakses dari luar package
	return "Good bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}
