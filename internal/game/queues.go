package game

import "RainbowRunner/pkg"

type MovementUpdate struct {
	Position pkg.Vector2
	Tick     uint
	Rotation int32
}

type MovementUpdateQueue struct {
	queue []MovementUpdate
}

func (q MovementUpdateQueue) Add(pos MovementUpdate) {

}

func (q *MovementUpdateQueue) Peek() *MovementUpdate {
	if len(q.queue) == 0 {
		return nil
	}

	return &q.queue[0]
}

func (q *MovementUpdateQueue) Dequeue() (move *MovementUpdate) {
	move = q.Peek()
	q.queue = q.queue[1:]
	return
}

func NewMovementUpdateQueue(size int) *MovementUpdateQueue {
	return &MovementUpdateQueue{
		queue: make([]MovementUpdate, size),
	}
}
