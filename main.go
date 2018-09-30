package main

import (
	"fmt"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	// for {
	// text, _ := reader.ReadString('\n')
	// atreus62: col=11, row=02, pressed=1, tap_count=0, tap_interrupted=0, keycode=26642, layer=0
	text := "atreus62:11,02,1,0,0,26642,0"

	fmt.Printf("From keylogger: %s", text)
	// }
}
