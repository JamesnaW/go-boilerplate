package model

import "time"

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func (t *Todo) Complete() {
	t.Name = "Already change!"
	t.Completed = true
	t.Due = time.Now()
}
