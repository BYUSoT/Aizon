package job

import "time"

type State struct {
	User string
	Rows []Row
}

type Row struct {
	Id int
	Name, Dept, PO, Team, Desc, Status string
	Date time.Time
}
