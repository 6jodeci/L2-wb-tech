// Concrete Command 1
package pkg

type OnCommand struct {
	device Device
}

func(c *OnCommand) Execute() {
	c.device.On()
}