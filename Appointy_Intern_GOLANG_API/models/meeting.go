package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Represents a meeting, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Meeting struct {
	ID           bson.ObjectId `bson:"_id" json:"id"`
	Title        string        `bson:"Title" json:"Title"`
	Participants *Participant  `bson:"Participants" json:"Participants"`
	StartTime    time.Time     `bson:"StartTime" json:"StartTime"`
	EndTime      time.Time     `bson:"EndTime" json:"EndTime"`
	Timestamp    time.Time     `bson:"Timestamp" json:"Timestamp"`
	// TimeDiff     time.Time     `bson:"TimeDiff" json:"TimeDiff"`
}
