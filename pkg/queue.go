package pkg

type Queue struct {
	queue []interface{}
}

func NewQueue(size int) *Queue {
	return &Queue{
		queue: make([]interface{}, size),
	}
}

type Vector2Queue struct {
	queue []Vector2
}

func (q Vector2Queue) Add(pos Vector2) {

}

func NewVector2Queue(size int) *Vector2Queue {
	return &Vector2Queue{
		queue: make([]Vector2, size),
	}
}
