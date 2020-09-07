package tracker

import (
	"github.com/AlexsJones/vinculum/pkg/proto"
	"sync"
)

var once sync.Once

type Tracker struct {
	Nodes map[string]*proto.NodeConfig
	mutex *sync.Mutex
}

var instance *Tracker

func Instance() *Tracker {

	once.Do(func() {
		instance = &Tracker{
			Nodes: make(map[string]*proto.NodeConfig),
			mutex: new(sync.Mutex),
		}
	})
	return instance
}

func (t *Tracker) Get(guid string) *proto.NodeConfig {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if val, ok := t.Nodes[guid]; ok {
		return val
	}
	return nil
}

func (t *Tracker) Count() int {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return len(t.Nodes)
}

func (t *Tracker) Add(node *proto.NodeConfig) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.Nodes[node.Guid] = node
}

