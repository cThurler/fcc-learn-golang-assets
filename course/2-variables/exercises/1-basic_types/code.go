package main

import "fmt"

func main() {
	// initialize variables here

	/* var smsSendingLimit int = 40;
	var costPerSMS float32 = 0.8;
	var hasPermission bool = true;
	var username string = "cthurler;" */

	smsSendingLimit, costPerSMS, hasPermission, username := 40, 0.8, true, "cthurler"

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
