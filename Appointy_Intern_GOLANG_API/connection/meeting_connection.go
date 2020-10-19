package connection

import (
	"log"

	. "../models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MeetingConn struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "users"
)

// Establish a connection to database
func (m *MeetingConn) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of meetings
func (m *MeetingConn) FindAll() ([]Meeting, error) {
	var meetings []Meeting
	err := db.C(COLLECTION).Find(bson.M{}).All(&meetings)
	return meetings, err
}

// Find a meeting by its id
func (m *MeetingConn) FindById(id string) (Meeting, error) {
	var meeting Meeting
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&meeting)
	return meeting, err
}

// func (m *MeetingConn) FindByEmail(id string) (Meeting, error) {
// 	var meeting Meeting
// 	err := db.C(COLLECTION).FindBy(bson.ObjectIdHex(id)).One(&meeting)
// 	return meeting, err
// }

// Insert a meeting into database
func (m *MeetingConn) Insert(meeting Meeting) error {
	err := db.C(COLLECTION).Insert(&meeting)
	return err
}

// Delete an existing meeting
func (m *MeetingConn) Delete(meeting Meeting) error {
	err := db.C(COLLECTION).Remove(&meeting)
	return err
}

// Update an existing meeting
func (m *MeetingConn) Update(meeting Meeting) error {
	err := db.C(COLLECTION).UpdateId(meeting.ID, &meeting)
	return err
}
