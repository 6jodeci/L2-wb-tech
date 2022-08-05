// client
package main

import "command-pattern/pkg"

func main() {
	var d1 pkg.Device
	d1 = &pkg.Tv{}
	d1.On()
	d1.Off()
}
