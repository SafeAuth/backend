package model

type ValidateUser struct {
	ValidUser bool
	Admin     bool
	Banned    bool
	Username  string
	Uid       int
	JWT       string
	ApiKey    string
}
