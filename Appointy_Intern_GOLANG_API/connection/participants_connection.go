package connection

import (
	"log"

	. "../models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ParticipantsConnection struct {
	Server   string
	Database string
}

var database *mgo.Database

const (
	COLLECTION1 = "participants"
)

// Establish a connection to database
func (m *ParticipantsConnection) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	database = session.DB(m.Database)
}

// Find list of participants
func (m *ParticipantsConnection) FindAll() ([]Participant, error) {
	var participants []Participant
	err := database.C(COLLECTION1).Find(bson.M{}).All(&participants)
	return participants, err
}

// Find a participants by its id
func (m *ParticipantsConnection) FindById(id string) (Participant, error) {
	var participants Participant
	err := database.C(COLLECTION1).FindId(bson.ObjectIdHex(id)).One(&participants)
	return participants, err
}

// Insert a participants into database
func (m *ParticipantsConnection) Insert(participants Participant) error {
	err := database.C(COLLECTION1).Insert(&participants)
	return err
}

// Delete an existing participants
func (m *ParticipantsConnection) Delete(participants Participant) error {
	err := database.C(COLLECTION1).Remove(&participants)
	return err
}

// Update an existing participants
func (m *ParticipantsConnection) Update(participants Participant) error {
	err := database.C(COLLECTION1).UpdateId(participants.ID, &participants)
	return err
}
