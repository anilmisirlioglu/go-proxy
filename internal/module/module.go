package module

type Module struct {
	Name   string   `json:"name"`
	Git    string   `json:"git"`
	Source []string `json:"source"`
}
