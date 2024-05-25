package entity

type User struct {
	Id             string
	Email          string
	Verified_email bool
	Name           string
	Given_name     string
	Family_picture string
	Picture        string
	Locale         string
}
