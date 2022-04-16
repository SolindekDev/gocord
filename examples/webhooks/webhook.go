package main

import (
	gocord "github.com/solindekdev/gocord"
)

func main() {
	Client := gocord.NewClient("NzQwNjEyMTU0MDYyNzk4OTI4.Xyri0Q.EaEOyvNbmvXwZsxZa1SszuiMwCE")
	Client.AllWebhooksOfChannel("963859949505642516")
}
