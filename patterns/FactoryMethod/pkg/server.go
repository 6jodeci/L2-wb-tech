package pkg

import "fmt"

type Server struct {
	Type   string
	Core   int
	Memory int
}

// базовый конструктор сервиса
func NewServer() Computer {
	return Server{
		Type:   ServerType,
		Core:   16,
		Memory: 256,
	}
}

// реализация сервера
func (pc Server) GetType() string {
	return pc.Type
}

// реализация сервера
func (pc Server) PrintDetails() {
	fmt.Printf("Type:||%s|| == Core:[%d] Mem:[%d]\n", pc.Type, pc.Core, pc.Memory)
}
