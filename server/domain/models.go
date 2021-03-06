package domain

import "golang.org/x/text/date"

// User is the struct of the administritor info.
type User struct {
	ID int
	Username string
	Password string
	Salt string
	Email string
}

// Slave is the struct of the servers to be sync with.
type Slave struct {
	ID int
	IP string
	EncryptKey string
	Version int
	LastSyncTime string
}

// Certification is the struct of SSL certification.
type Certification struct {
	ID int
	Domain string
	Path string
	Period string
	RelatedSlaves []int
}