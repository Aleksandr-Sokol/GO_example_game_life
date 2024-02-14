package main

type Position struct {
	// позиция на доске
	x int
	y int
}

type Border struct {
	// позиция на доске
	top    int
	bottom int
	right  int
	left   int
}

type Object struct {
	// организм
	id              string
	present         int
	future          int
	neighbors_count int
}

func (object *Object) update() {
	object.present = object.future
}
