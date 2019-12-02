package boomer

import "sync"

type TaskSetSequence struct {
	tasks       []*Task
	lock        sync.RWMutex
	totalTask   int
	currentTask int
}

// NewWeighingTaskSet returns a new WeighingTaskSet.
func NewTaskSetSequence() *TaskSetSequence {
	return &TaskSetSequence{
		totalTask:   0,
		currentTask: 0,
		tasks:       make([]*Task, 0),
	}
}

func (ts *TaskSetSequence) AddTask(task *Task) {
	ts.lock.Lock()
	ts.totalTask++
	ts.tasks = append(ts.tasks, task)
	ts.lock.Unlock()
}

func (ts *TaskSetSequence) GetTask() (task *Task) {
	ts.lock.Lock()
	defer func() {
		ts.currentTask++
		ts.lock.Unlock()
	}()

	return ts.tasks[ts.currentTask%ts.totalTask]
}
func (ts *TaskSetSequence) SetWeight(weight int) {}

func (ts *TaskSetSequence) GetWeight() (weight int) {
	return 0
}

func (ts *TaskSetSequence) Run() {
	task := ts.GetTask()
	task.Fn()
}
