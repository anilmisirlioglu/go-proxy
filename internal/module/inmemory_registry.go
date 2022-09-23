package module

import (
	"errors"
	"sync"
)

type InMemoryRegistry struct {
	sync.RWMutex

	m map[string]*Module
}

func NewInMemoryRegistry() Registry {
	return &InMemoryRegistry{
		m: map[string]*Module{},
	}
}

func (i *InMemoryRegistry) Register(m *Module) error {
	if m == nil {
		return errors.New("module is not be nil")
	}

	i.Lock()
	i.m[m.Name] = m
	i.Unlock()
	return nil
}

func (i *InMemoryRegistry) Unregister(name string) error {
	i.Lock()
	delete(i.m, name)
	i.Unlock()
	return nil
}

func (i *InMemoryRegistry) Search(name string) *Module {
	i.RLock()
	defer i.RUnlock()
	return i.m[name]
}

func (i *InMemoryRegistry) Index() []*Module {
	modules := make([]*Module, 0, len(i.m))
	for _, m := range i.m {
		modules = append(modules, m)
	}

	return modules
}
