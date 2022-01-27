package message

import (
	"RainbowRunner/pkg/byter"
	"sync"
)

type QueueType int

const (
	QueueTypeClientEntity QueueType = iota
)

//go:generate stringer -type=OpType
type OpType int

const (
	OpTypeAvatarMovement OpType = iota
	OpTypeCreateNPC
	OpTypeEquippedItemClickResponse
	OpTypeInventoryItemClickResponse
	OpTypeOther
)

type QueueItem struct {
	Type   QueueType
	Data   *byter.Byter
	OpType OpType
}

type Queue struct {
	sync.RWMutex

	queues map[QueueType][]*QueueItem
}

func (q *Queue) Enqueue(queueType QueueType, item *byter.Byter, opType OpType) {
	q.Lock()
	defer q.Unlock()

	if _, ok := q.queues[queueType]; !ok {
		q.queues[queueType] = make([]*QueueItem, 0, 1024)
	}

	q.queues[queueType] = append(q.queues[queueType], &QueueItem{
		Type:   queueType,
		Data:   item,
		OpType: opType,
	})
}

func (q *Queue) Dequeue(queueType QueueType) *QueueItem {
	q.Lock()
	defer q.Unlock()

	item := q.queues[queueType][0]

	q.queues[queueType] = q.queues[queueType][1:]

	return item
}

func (q *Queue) IsEmpty(queueType QueueType) bool {
	q.RLock()
	defer q.RUnlock()

	if _, ok := q.queues[queueType]; !ok {
		return true
	}

	return len(q.queues[queueType]) == 0
}

func NewQueue() *Queue {
	return &Queue{queues: map[QueueType][]*QueueItem{}}
}
