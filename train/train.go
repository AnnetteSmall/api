package train

import "fmt"

type Engine struct {
	speed int
}

func NewEngine() *Engine {
	fmt.Println("New Engine in Func is ", Engine{})
	return &Engine{}
}

func (e *Engine) Speed() int {
	fmt.Println("speed in Func is ", e.speed)
	return e.speed
}

func (e *Engine) Accel() {
	e.speed += 10
	fmt.Println("Accel by 10, speed is now ",e.speed)
}

func (e *Engine) Decel() {
	e.speed -= 10
	fmt.Println("Decel by 10, speed is now ",e.speed)
}
