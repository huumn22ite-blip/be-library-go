package handlers

import (
	"be-library-go/db"
	"be-library-go/models"
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)
	
func GetLoans(w http.ResponseWriter, r*http.Request){
	rows, err := db.DB.Query("SELECT book_id,member_id,staff_id,borrow_date,due_date,return_date,status FROM loans")
	if err !=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
return
	}
	defer rows.Close()
	var loans []models.Loans
	for rows.Next(){
		var c models.Loans
		err := rows.Scan(&c.ID, &c.BOOK_ID, &c.MEMBER_ID, &c.STAFF_ID,&c.BORROW_DATE,&c.DUE_DATE, &c.RETURN_DATE, &c.STATUS,)
		if err !=nil {
			http.Error(w,err.Error(),  http.StatusInternalServerError)
			return
		}
	

		loans = append(loans,c )
	}
	json.NewEncoder(w).Encode(loans)
}
func CreateLoans(w http.ResponseWriter, r *http.Request) {

	var c models.Loans

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := db.DB.Exec(
		"INSERT INTO loans( book_id,member_id,staff_id,borrow_date,due_date,return_date,status) VALUES(?,?,?,?,?,?,?)",
		c.BOOK_ID,
		c.MEMBER_ID,
		c.STAFF_ID,
		c.BORROW_DATE,
		c.DUE_DATE,
		c.RETURN_DATE,
		c.STATUS,
		
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.ID = int(id)

	json.NewEncoder(w).Encode(c)
}

func DeleteLoans(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	_, err := db.DB.Exec("DELETE FROM loans WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Deleted",
	})
}

func UpdateLoans(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	var c models.Loans

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec(
		"UPDATE loans SET ,name=?,email=?,phone=?,membership_date=?,status=?,memberscol=? WHERE id=?",	
			c.BOOK_ID,
		c.MEMBER_ID,
		c.STAFF_ID,
		c.BORROW_DATE,
		c.DUE_DATE,
		c.RETURN_DATE,
		c.STATUS,
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.ID, _ = strconv.Atoi(id)

	json.NewEncoder(w).Encode(c)
}