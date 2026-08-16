package todo

type List struct {
	tasks map[string]Task
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (l *List) AddTask(task Task) error {
	if _, ok := l.tasks[task.Title]; ok {
		return ErrTaskAlreadyExist
	}
	l.tasks[task.Title] = task
	return nil
}

func (l *List) GetTask(title string) (Task, error) {
	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	return task, nil
}

func (l *List) ListTasks() map[string]Task {
	tmp := make(map[string]Task, len(l.tasks))
	for k, v := range l.tasks {
		tmp[k] = v
	}
	return tmp
}

func (l *List) ListUncompletedTasks() map[string]Task {
	uncompletedTasks := make(map[string]Task)
	for title, task := range l.tasks {
		if !task.Completed {
			uncompletedTasks[title] = task
		}
	}
	return uncompletedTasks
}

func (l *List) CompleteTask(title string) error {
	task, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}
	task.Complete()

	l.tasks[title] = task
	return nil
}

func (l *List) UncompleteTask(title string) error {
	task, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}
	task.Complete()

	l.tasks[title] = task
	return nil
}

func (l *List) DeleteTask(title string) error {
	_, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}

	delete(l.tasks, title)

	return nil
}
