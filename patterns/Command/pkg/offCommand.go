// Concrete Command 1
package pkg

type OffCommand struct {
	device Device
}

func (c *OffCommand) Execute() {
	c.device.Off()
}