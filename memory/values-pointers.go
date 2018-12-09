package main

type point struct{ x, y float32 }

type circle struct {
	center point
	radius float32
}

func main() {
	c := circle{center: point{1, 1}, radius: 1}
	getRadius(c)
	updateRadius(&c) // No pointer arithmetic, but have "unsafe" package
}

func getRadius(v circle) float32 {
	return v.radius
}

func updateRadius(p *circle) {
	p.radius += 1 // No special syntax for pointers, just '.'
}

func valueNotEscape() {
	c := circle{center: point{1, 1}, radius: 1}
	getRadius(c)
}

func pointerNotEscape() {
	c := circle{center: point{1, 1}, radius: 1}
	updateRadius(&c)
}

func escapes() *circle {
	n := circle{center: point{3, 3}, radius: 3}
	return &n
}
