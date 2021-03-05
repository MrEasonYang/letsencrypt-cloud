package domain

import "golang.org/x/text/date"

type User struct {
	ID int
	Username string
	Password string
	Salt string
	Email string
}

type Slave struct {
	ID int
	IP string
	EncryptKey string
	Version int
	LastSyncTime string
}

