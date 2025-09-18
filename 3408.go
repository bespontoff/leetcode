package main

import "container/heap"

type TaskManager struct {
	tasks Tasks
}

type Task struct {
	userId   int
	taskId   int
	priority int
	index    int
}

type Tasks []*Task

func (t *Tasks) Len() int {
	return len(*t)
}

func (t *Tasks) Less(i, j int) bool {
	return (*t)[i].priority < (*t)[j].priority
}

func (t *Tasks) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
	(*t)[i].index = i
	(*t)[j].index = j
}

func (t *Tasks) Push(x any) {
	n := len(*t)
	task := x.(*Task)
	task.index = n
	*t = append(*t, task)
}

func (t *Tasks) Pop() any {
	old := *t
	n := len(old)
	task := old[n-1]
	old[n-1] = nil
	task.index = -1
	*t = old[0 : n-1]
	return task
}

func Constructor(tasks [][]int) TaskManager {
	taskManager := TaskManager{
		tasks: make([]*Task, 0),
	}
	for i, task := range tasks {
		taskManager.tasks = append(taskManager.tasks, &Task{
			userId:   task[0],
			taskId:   task[1],
			priority: task[2],
			index:    i,
		})
	}
	heap.Init(&taskManager.tasks)
	return taskManager
}

func (tm *TaskManager) Add(userId int, taskId int, priority int) {
	t := &Task{
		userId:   userId,
		taskId:   taskId,
		priority: priority,
	}
	heap.Push(&tm.tasks, t)
}

func (tm *TaskManager) Edit(taskId int, newPriority int) {

}

func (tm *TaskManager) Rmv(taskId int) {

}

func (tm *TaskManager) ExecTop() int {

	return 0
}
