package handlers

import (
	"be-library-go/db"
	"be-library-go/models"
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

func GetStaff(w http.ResponseWriter, r*http.Request){
	rows, err := db.DB.Query("SELECT id,name,email,phone FROM staff")
	if err !=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
return
	}
	defer rows.Close()
	var staff []models.Staff
	for rows.Next(){
		var c models.Staff
		err := rows.Scan(&c.ID, &c.NAME, &c.EMAIL,&c.PHONE)
		if err !=nil {
			http.Error(w,err.Error(),  http.StatusInternalServerError)
			return
		}


		staff = append(staff, c)
	}
	json.NewEncoder(w).Encode(staff)
}
func CreateStaff(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var c models.Staff
    if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }

    res, err := db.DB.Exec(
        "INSERT INTO staff(name,email,phone) VALUES(?,?,?)",
        c.NAME, c.EMAIL, c.PHONE,
    )
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }

    id, _ := res.LastInsertId()
    c.ID = int(id)

    json.NewEncoder(w).Encode(c) // trả JSON thành công
}
func DeleteStaff(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	_, err := db.DB.Exec("DELETE FROM staff WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Deleted",
	})
}

func UpdateStaff(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	var c models.Staff

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec(
		"UPDATE staff SET ,name=?,email=?,phone=?,membership_date=?,status=?,memberscol=? WHERE id=?",	
		c.NAME,
		c.EMAIL,
		c.PHONE,
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.ID, _ = strconv.Atoi(id)

	json.NewEncoder(w).Encode(c)
}