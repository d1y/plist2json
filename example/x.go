package main

import (
	"fmt"
	"log"

	"github.com/d1y/plist2json"
)

func main() {
	var p = `data/SoundsSettings.plist`

	// check plist file
	x := plist2json.EasyCheck(p)
	xText := fmt.Sprintf("the file is %v", x)
	log.Println(xText)

	// format the `plist` file
	formatFlag := plist2json.Easy("SoundsSettings.plistx", "cao.json")
	if formatFlag != nil {
		log.Fatalln(formatFlag)
	}
	fmt.Println("format success")
}
