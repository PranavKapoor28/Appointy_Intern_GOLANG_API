package models

import "gopkg.in/mgo.v2/bson"

// Represents Participnats, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Participant struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	Name  string        `bson:"Name" json:"Name"`
	Email string        `bson:"Email" json:"Email"`
	RSVP  string        `bson:"RSVP" json:"RSVP"`
}
