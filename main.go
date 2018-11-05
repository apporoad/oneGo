package main

import (
	"fmt"
	_ "os"
	_ "strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	sh := shell.NewShell("localhost:5001")

	fmt.Println(sh.Cat("/ipfs/QmWf5P56MCRMN1mmjACD3NBemqHDcxtueLjea5zM4CSgGo"))

	// cid , err := sh.Add(strings.NewReader("hello good day"))

	// if err !=nil {
	// 	fmt.Fprintf(os.Stderr, "error : %s",err)
	// 	os.Exit(1)
	// }
	// fmt.Println("added %s",cid)
	//cid ,err = sh.Add
}
