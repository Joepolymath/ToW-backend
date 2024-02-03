package models

type User struct {
	FirstName	string		`json:"first_name"`
	LastName		string		`json:"last_name"`
	Tagline		string		`json:"tagline"`
	Country		string		`json:"country"`
	ZipCode		string		`json:"zip_code"`
	Email			string		`json:"email"`
	Profession	string		`json:"profession"`
	Income		uint32		`json:"income"`
	Company		string		`json:"company"`
	Links			string		`json:"links"`
	UserID		uint32		`json:"user_id"`
	Followers	uint32		`json:"followers"`
}
