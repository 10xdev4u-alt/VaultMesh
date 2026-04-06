package healing

import (
	"container/heap"
	"sync"

	"github.com/10xdev4u-alt/VaultMesh/internal/storage"
)

// RepairTask represents a pending repair operation for a file.
type RepairTask struct {
	Manifest *storage.Manifest
	Priority int // Lower number = Higher priority
	index    int 
}

// TaskQueue implements heap.Interface and holds RepairTasks.
type TaskQueue []*RepairTask

func (tq TaskQueue) Len() int           { return len(tq) }
func (tq TaskQueue) Less(i, j int) bool { return tq[i].Priority < tq[j].Priority }
func (tq TaskQueue) Swap(i, j int)      { tq[i], tq[j] = tq[j], tq[i]; tq[i].index = i; tq[j].index = j }

func (tq *TaskQueue) Push(x any) {
	n := len(*tq)
	item := x.(*RepairTask)
	item.index = n
	*tq = append(*tq, item)
}

func (tq *TaskQueue) Pop() any {
	old := *tq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*tq = old[0 : n-1]
	return item
}

// RepairScheduler manages the execution of repair tasks based on priority.
type RepairScheduler struct {
	mu    sync.Mutex
	queue TaskQueue
}

// NewRepairScheduler creates a new RepairScheduler.
func NewRepairScheduler() *RepairScheduler {
	rs := &RepairScheduler{
		queue: make(TaskQueue, 0),
	}
	heap.Init(&rs.queue)
	return rs
}

// AddTask adds a manifest to the repair queue with a specific priority.
func (s *RepairScheduler) AddTask(m *storage.Manifest, priority int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	heap.Push(&s.queue, &RepairTask{Manifest: m, Priority: priority})
}

// GetNextTask retrieves the highest priority task from the queue.
func (s *RepairScheduler) GetNextTask() *RepairTask {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.queue.Len() == 0 {
		return nil
	}
	return heap.Pop(&s.queue).(*RepairTask)
}
