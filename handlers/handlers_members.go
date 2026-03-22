package handlers

import (
	"be-library-go/db"
	"be-library-go/models"
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)
func GetMembers(w http.ResponseWriter, r*http.Request){
	rows, err := db.DB.Query("SELECT id,name,phone,address,membership_date,status  FROM members")
	if err !=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
return
	}
	defer rows.Close()
	var members []models.Members
	for rows.Next(){
		var c models.Members
		err := rows.Scan(&c.ID, &c.NAME, &c.PHONE,&c.ADDRESS, &c.MEMBERSHIP_DATE, &c.STATUS,)
		if err !=nil {
			http.Error(w,err.Error(),  http.StatusInternalServerError)
			return
		}


		members = append(members, c)
	}
	json.NewEncoder(w).Encode(members)
}
func CreateMembers(w http.ResponseWriter, r *http.Request) {

	var c models.Members

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := db.DB.Exec(
		"INSERT INTO members( name,phone,address,membership_date,status) VALUES(?,?,?,?,?)",
		c.NAME,
		c.PHONE,
		c.ADDRESS,
		c.MEMBERSHIP_DATE,
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

func DeleteMembers(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	_, err := db.DB.Exec("DELETE FROM members WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Deleted",
	})
}

func UpdateMembers(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	var c models.Members

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec(
		"UPDATE members SET ,name=?,phone=?,address=?,membership_date=?,status=?, WHERE id=?",	
		c.NAME,
		c.PHONE,
		c.ADDRESS,
		c.MEMBERSHIP_DATE,
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