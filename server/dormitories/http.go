package dormitories

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FictProger/architecture-lab-3/server/tools"
)

// Channels HTTP handler.
type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of channels HTTP handler.
func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleRecomendDormitory(r, store, rw)
		} else if r.Method == "POST" {
			handleStudentCreate(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleStudentCreate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var stud Student
	if err := json.NewDecoder(r.Body).Decode(&stud); err != nil {
		log.Printf("Error decoding channel input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.AddStudent(stud.DormId, stud.Specialty)
	if err == nil {
		tools.WriteJsonOk(rw, &stud)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleRecomendDormitory(r *http.Request, store *Store, rw http.ResponseWriter) {
	specity, ok := r.URL.Query()["specity"]
	if !ok || len(specity) > 1 {
		log.Println("Error query is wrong")
		tools.WriteJsonBadRequest(rw, "wrong query")
		return
	}

	res, err := store.RecomendDormitory(specity[0])
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
