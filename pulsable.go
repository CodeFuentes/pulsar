package main

type Pulsable struct {
	body     string
	position Position
	speed    int
}

func NewPulsable(body string, position Position) *Pulsable {
	return &Pulsable{
		body:     body,
		position: position,
	}
}

func (p Pulsable) Position() Position {
	return p.position
}

func (p Pulsable) Speed() int {
	return p.speed
}

func (p Pulsable) Body() string {
	return p.body
}

func (p Pulsable) Len() int {
	return len(p.body)
}
