package file

type OnTaskClosure func(job *Job)

func InitWatchTaskCallable(t *WatchTask) {
	t.onCreateClosureList = make([]OnTaskClosure, 0)
	t.onRenameClosureList = make([]OnTaskClosure, 0)
	t.onDeleteClosureList = make([]OnTaskClosure, 0)
}

func (t *WatchTask) DoCreateCallable(job *Job) {
	for i := 0; i < len(t.onCreateClosureList); i++ {
		t.onCreateClosureList[i](job)
	}
}

func (t *WatchTask) DoRenameCallable(job *Job) {
	for i := 0; i < len(t.onRenameClosureList); i++ {
		t.onRenameClosureList[i](job)
	}
}

func (t *WatchTask) DoDeleteCallable(job *Job) {
	for i := 0; i < len(t.onDeleteClosureList); i++ {
		t.onDeleteClosureList[i](job)
	}
}

func (t *WatchTask) AddCreateCallable(closure OnTaskClosure) {
	t.onCreateClosureList = append(t.onCreateClosureList, closure)
}

func (t *WatchTask) AddRenameCallable(closure OnTaskClosure) {
	t.onRenameClosureList = append(t.onRenameClosureList, closure)
}

func (t *WatchTask) AddDeleteCallable(closure OnTaskClosure) {
	t.onDeleteClosureList = append(t.onDeleteClosureList, closure)
}
