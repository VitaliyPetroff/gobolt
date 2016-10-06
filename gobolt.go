package main

import (
	// "encoding/binary"
	// "encoding/json"
	// "fmt"
	// "strconv"

	"github.com/boltdb/bolt"
)

// Data base location settings
type GoBoltConf struct {
	path string
}

type Gobolt struct {
	gb *bolt.DB
}

// type Object struct {
//   ID   int    `json:"id"`
//   Name string `json:"name"`
//   Data string `json:"data"`
// }

// Open database
func (gbase *Gobolt) OpenGB(gbc GoBoltConf) error {
	var err error
	// ToDo: in safe mode don't create database if it's not exist.
	gbase.gb, err = bolt.Open(gbc.path, 0644, nil)

	return err
}

// Safe close database.
// Don't forget to use it after opening.
// Ex.: defer database.Close()
func (gbase *Gobolt) Close() {
	gbase.gb.Close()
}

// Create new bucket
func (gbase *Gobolt) NewBucket(bucket_name string) (*bolt.Bucket, error) {
	bucket := new(bolt.Bucket)
	err := gbase.gb.Update(func(tx *bolt.Tx) error {
		var err error
		bucket, err = tx.CreateBucketIfNotExists([]byte(bucket_name))
		return err
	})

	return bucket, err
}

// Save object

func main() {

}
