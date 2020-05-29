package job

import (
	"database/sql"
	"log"
	"fmt"
)

var current *sql.Stmt 

func Setup(db *sql.DB) {
	var err error
	current, err = db.Prepare("select homepage_order.id, firstname, lastname, homepage_department.name," +
			" homepage_order.dateOrdered, homepage_order.customerPO, homepage_order.teamLab," +
			" homepage_order.description, homepage_jobstatus.name" +
		" from homepage_order inner join homepage_person on homepage_order.customer_id = homepage_person.id" +
			" inner join homepage_department on homepage_order.department_id = homepage_department.id" +
			" inner join homepage_joborder on homepage_order.id = homepage_joborder.order_ptr_id" +
			" inner join homepage_jobstatus on homepage_joborder.jobStatus_id = homepage_jobstatus.id" +
		" where dateCompleted = 0" +
			" and homepage_jobstatus.name != 'Cancelled'" + 
		" order by id asc")
	if err != nil {
		log.Fatal(err)
	}
}

func Current() []Row {
	rs, err := current.Query()
	if err != nil {
		log.Print(err)
	}
	res := []Row{}
	cursor := Row{}
	var fname, lname string
	for rs.Next() {
		rs.Scan(&cursor.Id, &fname, &lname, &cursor.Dept,
				&cursor.Date, &cursor.PO, &cursor.Team,
				&cursor.Desc, &cursor.Status)
		cursor.Name = lname + ", " + fname
fmt.Printf("scanning: %+v\n", cursor);
		res = append(res, cursor)
	}
	return res
}

