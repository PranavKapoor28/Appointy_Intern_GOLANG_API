# Appointy_Intern_GOLANG_API
The task is to develop a basic version of meeting scheduling API. It was required to develop the API for the system.

# Basic external modules for REST API

mux - Request router and dispatcher for matching incoming requests to their respective handler
mgo - MongoDB driver
toml - Parse the configuration file (MongoDB server & credentials)

# Functions:

1.This is an API developed in GO along with MongoDB. 
2.This API schedules a meeting given id,title,participants and other details. 
3.One can get the meeting details by providing the meeting ID. 
4.It can also return the list of meetings arranged within a given time range. 
5.Finally it can return all the meetings in which a particular participant is present.
6.The meting can be updated and deleted.
7. On creation of a meeting it automatically generates a Timestamp.
