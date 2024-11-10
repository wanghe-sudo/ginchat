package main

import "ginchat/rotuer"

func main() {
	r := rotuer.Router()
	r.Run(":8080")
}
