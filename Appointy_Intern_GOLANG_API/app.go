package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	. "./config"
	. "./connection"
	. "./models"
	"github.com/gorilla/mux"
)

var config = Config{}
var dao = MeetingConn{}
var participantsconn = ParticipantsConnection{}

// GET list of meetings
func AllMeetingEndPoint(w http.ResponseWriter, r *http.Request) {
	meetings, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, meetings)
}

// GET list of Participants
func AllParticipantEndPoint(w http.ResponseWriter, r *http.Request) {
	participant, err := participantsconn.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, participant)
}

// GET a meetings by its ID
func FindMeetingByIdEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	meetings, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid meetings ID")
		return
	}
	respondWithJson(w, http.StatusOK, meetings)
}

// GET a meetings by its ID
func FindMeetingByEmailEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	meetings, err := participantsconn.FindById(params["Email"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Email")
		return
	}
	respondWithJson(w, http.StatusOK, meetings)
}

// POST a new meeting
func CreateMeetingEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var meetings Meeting

	// StartTime, _ := mux.Vars(r)["StartTime"]
	// EndTime, _ := mux.Vars(r)["EndTime"]
	// cur, err := dao.Find(context.TODO(), bson.M{{"StartTime": StartTime}, {"EndTime": EndTime}})

	if err := json.NewDecoder(r.Body).Decode(&meetings); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	meetings.ID = bson.NewObjectId()
	meetings.Timestamp = time.Now()

	if err := dao.Insert(meetings); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, meetings)
}

// POST a new meeting
func CreateParticipantEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var participant Participant
	if err := json.NewDecoder(r.Body).Decode(&participant); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	participant.ID = bson.NewObjectId()
	if err := participantsconn.Insert(participant); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, participant)
}

// PUT update an existing meeting
func UpdateMeetingEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var meetings Meeting
	meetings.ID = bson.ObjectIdHex(params["id"])
	if err := json.NewDecoder(r.Body).Decode(&meetings); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(meetings); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing meeting
func DeleteMeetingEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var meetings Meeting
	if err := json.NewDecoder(r.Body).Decode(&meetings); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(meetings); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()

	participantsconn.Server = config.Server
	participantsconn.Database = config.Database
	participantsconn.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/meetings", AllMeetingEndPoint).Methods("GET")
	r.HandleFunc("/participants", AllParticipantEndPoint).Methods("GET")
	r.HandleFunc("/participants", CreateParticipantEndPoint).Methods("POST")
	r.HandleFunc("/meetings", CreateMeetingEndPoint).Methods("POST")
	r.HandleFunc("/meetings/{id}", UpdateMeetingEndPoint).Methods("PUT")
	r.HandleFunc("/meetings", DeleteMeetingEndPoint).Methods("DELETE")
	r.HandleFunc("/meetings/{id}", FindMeetingByIdEndpoint).Methods("GET")
	r.HandleFunc("/meetings/{Email}", FindMeetingByEmailEndpoint).Methods("GET")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}

}
