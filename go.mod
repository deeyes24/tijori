module github.com/deeyes21/tijori

go 1.15

replace github.com/tijori/cmd => ./cmd

replace github.com/tijori/tijori => ./tijori

require (
	github.com/atotto/clipboard v0.1.4 // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.2.1 // indirect
	github.com/tijori/cmd v0.0.0-00010101000000-000000000000
	github.com/tijori/config v0.0.0-00010101000000-000000000000 // indirect
	github.com/tijori/tijori v0.0.0-00010101000000-000000000000
)

replace github.com/tijori/config => ./config
