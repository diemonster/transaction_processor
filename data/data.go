package data

type Entry struct {
	Datetime  string `json:"datetime"`
	Value     string `json:"value"`
	Partition string `json:"partition"`
}
