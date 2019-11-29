package boomer

import "sync"

type TaskSetSequence struct {
	tasks       []*Task
	lock        sync.RWMutex
	totalTask   int
	currentTask int
}

func (ts TaskSetSequence) AddTask(task *Task) {
	ts.lock.Lock()
	ts.totalTask++
	ts.tasks = append(ts.tasks, task)
	ts.lock.Unlock()
}

func (ts *TaskSetSequence) GetTask() (task *Task) {
	ts.lock.RLock()
	defer func() {
		ts.currentTask++
		ts.lock.Unlock()
	}()

	return ts.tasks[ts.currentTask]
}
func (ts *TaskSetSequence) SetWeight(weight int) {}

func (ts *TaskSetSequence) GetWeight() (weight int){
	return 0
}


func (ts TaskSetSequence) Run() {
	task := ts.GetTask()
	task.Fn()
}
