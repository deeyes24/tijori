package config

type SavedPasswords struct {
	savedPasswords []*SavedPassword
}

type SavedPassword struct {
	ID             int
	UserName       string
	Password       string
	AdditionalInfo string
}
