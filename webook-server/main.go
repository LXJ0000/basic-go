package main

func main() {
	r := InitWebServer()
	_ = r.Run(":8080")
}
