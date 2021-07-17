package models

type Job struct {
	Name      string   `json:"name"`
	Partition string   `json:"partition"`
	GPU       int      `json:"gpu"`
	Modules   []string `json:"modules"`
	Command   string   `json:"command"`
}
