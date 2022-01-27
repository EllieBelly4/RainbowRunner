package global

import "sync"

type TickJobQueue struct {
	sync.RWMutex
	jobs []func()
}

func (q *TickJobQueue) Enqueue(job func()) {
	q.Lock()
	defer q.Unlock()

	q.jobs = append(q.jobs, job)
}

func (q *TickJobQueue) Dequeue() func() {
	q.Lock()
	defer q.Unlock()

	job := q.jobs[0]

	q.jobs = q.jobs[1:]

	return job
}

func (q *TickJobQueue) IsEmpty() bool {
	q.RLock()
	defer q.RUnlock()

	return len(q.jobs) == 0
}
