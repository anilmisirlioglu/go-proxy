package module

type Registry interface {
	Register(m *Module) error
	Unregister(name string) error
	Search(name string) *Module
	Index() []*Module
}
