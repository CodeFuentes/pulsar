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

func (p *Pulsable) SetPosition(pos Position) {
	p.position = pos
}

func (p Pulsable) GetPosition() Position {
	return p.position
}

func (p Pulsable) GetSpeed() int {
	return p.speed
}

func (p Pulsable) GetBody() string {
	return p.body
}

func (p Pulsable) Len() int {
	return len(p.body)
}
