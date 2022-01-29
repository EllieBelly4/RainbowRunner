package pkg

import "RainbowRunner/pkg/datatypes"

type Queue struct {
	queue []interface{}
}

func NewQueue(size int) *Queue {
	return &Queue{
		queue: make([]interface{}, size),
	}
}

type Vector2Queue struct {
	queue []datatypes.Vector2
}

func (q Vector2Queue) Add(pos datatypes.Vector2) {

}

func NewVector2Queue(size int) *Vector2Queue {
	return &Vector2Queue{
		queue: make([]datatypes.Vector2, size),
	}
}
