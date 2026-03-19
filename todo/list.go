package todo

type List struct {
	tasks map[string]Task
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (l *List) AddTask(task Task) {
	l.tasks[task.Title] = task
}

func (l *List) ListTasks() map[string]Task {
	copiedTasks := make(map[string]Task, len(l.tasks))

	for title, task := range l.tasks {
		copiedTasks[title] = task
	}

	return copiedTasks
}

func (l *List) DoneTask(title string) string {
	task, ok := l.tasks[title]
	if !ok {
		return taskNotFound
	}

	if task.IsDone {
		return taskAlreadyDone
	}

	task.Done()

	l.tasks[title] = task

	return ""
}

func (l *List) DeleteTask(title string) string {
	_, ok := l.tasks[title]
	if !ok {
		return taskNotFound
	}

	delete(l.tasks, title)

	return ""
}
