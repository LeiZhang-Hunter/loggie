package file

import (
	"github.com/loggie-io/loggie/pkg/core/log"
	"sync"
	"time"
)

type SeqDBConfig struct {
	File                 string        `yaml:"file,omitempty" default:"./data/loggie.db"`
	FlushTimeout         time.Duration `yaml:"flushTimeout,omitempty" default:"2s"`
	BufferSize           int           `yaml:"bufferSize,omitempty" default:"2048"`
	TableName            string        `yaml:"tableName,omitempty" default:"registry"`
	CleanInactiveTimeout time.Duration `yaml:"cleanInactiveTimeout,omitempty" default:"504h"` // default records not updated in 21 days will be deleted
	CleanScanInterval    time.Duration `yaml:"cleanScanInterval,omitempty" default:"1h"`
}

type SeqConfig struct {
	Enable *bool `yaml:"active,omitempty" default:"false"`
}

type SeqHandle struct {
	watchUidTable map[string]uint64
	mutex         sync.RWMutex
}

func InitSeqHandle(s *Source) {
	if !(*s.config.Seq.Enable) {
		return
	}
	s.seq = NewSeqHandle()
	s.watchTask.AddCreateCallable(s.seq.OnCreate)
	s.watchTask.AddRenameCallable(s.seq.OnRename)
	s.watchTask.AddDeleteCallable(s.seq.OnDelete)
}

func NewSeqHandle() *SeqHandle {
	return &SeqHandle{
		watchUidTable: make(map[string]uint64),
	}
}

func (h *SeqHandle) OnCreate(job *Job) {
	log.Info("OnCreate")
}

func (h *SeqHandle) OnRename(job *Job) {
	log.Info("OnRename")
}

func (h *SeqHandle) OnDelete(job *Job) {
	log.Info("OnDelete")
}

func (h *SeqHandle) OnProduct(job *Job) {
	log.Info("OnDelete")
}
