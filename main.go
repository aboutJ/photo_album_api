package main

import (
	_ "photo_album/config"
	"photo_album/route"
)

func main() {
	route.InitRoute()
}
