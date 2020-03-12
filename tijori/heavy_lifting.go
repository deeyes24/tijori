package tijori

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/tijori/config"
)

var passwordStore = []byte("passwordstore")
var db *bolt.DB

//Load
func LoadSavedPasswords() []*config.SavedPassword {
	var storedPasswords []*config.SavedPassword

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(passwordStore)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var p *config.SavedPassword

			err := json.Unmarshal(v, &p)
			if err != nil {
				return err
			} else {
				storedPasswords = append(storedPasswords, p)
			}

		}
		return nil

	})

	if err != nil {
		fmt.Println("caught error ", err.Error())
	}

	return storedPasswords

}

//Add
func AddtoSavedPasswords(config config.SavedPassword) {

	byteArray, err := json.Marshal(config)
	if err != nil {
		fmt.Println("caught error ", err.Error())
	}
	err = db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket(passwordStore)
		return b.Put(getKey(b), byteArray)
	})
	if err != nil {
		fmt.Println("caught error while trying to add ", err.Error())
		os.Exit(1)
	}

}

//Fetch
func FetchSavedPasswordFor(id int) config.SavedPassword {

	saved := LoadSavedPasswords()
	if id == 0 || id > len(saved) {
		fmt.Println("id cannot be 0 or more than ", len(saved))
		os.Exit(1)
	}

	return *saved[id-1]

}

func getKey(b *bolt.Bucket) []byte {
	id, _ := b.NextSequence()
	key := itob(int(id))
	return key
}

func Init(dbPath string) error {

	var err error

	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(passwordStore)
		return err
	})
}

func must(err error) {
	if err != nil {
		panic("Caught an error with message " + err.Error())
	}
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b

}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
