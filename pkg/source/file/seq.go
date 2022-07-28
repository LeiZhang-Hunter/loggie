package file

type SeqHandle struct {
	watchUidTable map[string]uint64
}

func NewSeqHandle() *SeqHandle {
	return &SeqHandle{
		watchUidTable: make(map[string]uint64),
	}
}

func (h *SeqHandle) OnCreate(job *Job) {

}

func (h *SeqHandle) OnRename(job *Job) {

}

func (h *SeqHandle) OnDelete(job *Job) {

}
