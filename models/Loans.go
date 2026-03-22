package models



type Loans struct {
	ID          int    `json:"id"`
	BOOK_ID     int `json:"book_id"`
	MEMBER_ID   int `json:"member_id"`
	STAFF_ID    int `json:"staff_id"`
	BORROW_DATE string  `json:"borrow_date"`
	DUE_DATE     string `json:"due_date"`
	RETURN_DATE  string `json:"return_date"`
	STATUS   string `json:"status"`
}