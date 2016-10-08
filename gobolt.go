package main

import (
	// "encoding/binary"
	//"bytes"
	"encoding/json"
	"fmt"
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

type object interface{}

// Open database
func (gbase *Gobolt) Open(gbc GoBoltConf) error {
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

// Get list of all buckets in the database
func (gbase *Gobolt) GetBucketList() ([]string, error) {
	var buckets []string

	err := gbase.gb.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			buckets = append(buckets, string(name))
			return nil
		})
	})
	return buckets, err
}

// Create new bucket
func (gbase *Gobolt) CreateBucket(bucket_name string) (*bolt.Bucket, error) {
	bucket := new(bolt.Bucket)
	err := gbase.gb.Update(func(tx *bolt.Tx) error {
		var err error
		bucket, err = tx.CreateBucketIfNotExists([]byte(bucket_name))
		return err
	})
	return bucket, err
}

// Get all data from bucket
func (gbase *Gobolt) GetBucketData(bucket_name string) (interface{}, error) {
	err := gbase.gb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket_name))
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", bucket_name)
		}

		bucket.ForEach(func(k, v []byte) error {
			s := string(v[:])
			fmt.Printf("key=%s, value=%s\n", k, s)
			return nil
		})

		return nil
	})

	return nil, err
}

// Create new object in bucket or rewrite existing (safe mode enabled)
func (gbase *Gobolt) SetByKey(bucket_name string, key_name string, obj interface{}) error {
	err := gbase.gb.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucket_name))
		if err != nil {
			return err
		}

		//id, _ := bucket.NextSequence()
		//obj.ID = int(id)
		id := []byte(key_name)

		buf, err := json.Marshal(obj)
		if err != nil {
			return err
		}

		err = bucket.Put(id, buf)
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

// Get object by key
func (gbase *Gobolt) GetByKey(bucket_name string, key_name string) (*object, error) {
	err := gbase.gb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket_name))
		if bucket == nil {
			return fmt.Errorf("Bucket %s not found!", bucket_name)
		}

		bucket.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		})

		return nil
	})

	return nil, err
}

type mySubObject struct {
	F1 string `json:"F1"`
	F2 string
}

type myObject struct {
	F1 string
	F2 string
	F3 mySubObject
}

func main() {
	var gbc GoBoltConf
	var bname1 = "bucket_one"
	var bname2 = "bucket_two"

	gbc.path = "testdb.bdb"

	db := new(Gobolt)
	err := db.Open(gbc)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = db.CreateBucket(bname1)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = db.CreateBucket(bname2)
	if err != nil {
		fmt.Println(err.Error())
	}

	mo1 := myObject{F1: "aa", F2: "ab", F3: mySubObject{F1: "aca", F2: "acb"}}
	mo2 := myObject{F1: "ba", F2: "bb", F3: mySubObject{F1: "bca", F2: "bcb"}}
	mo3 := myObject{F1: "ca", F2: "cb", F3: mySubObject{F1: "cca", F2: "ccb"}}

	err = db.SetByKey(bname1, "mo1", mo1)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.SetByKey(bname1, "mo2", mo2)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.SetByKey(bname1, "mo3", mo3)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.SetByKey(bname2, "mo1", mo1)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.SetByKey(bname2, "mo2", mo2)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.SetByKey(bname2, "mo3", mo3)
	if err != nil {
		fmt.Println(err.Error())
	}

	var list []string
	list, err = db.GetBucketList()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("readed from db:")
		fmt.Println(list)
	}

	_, _ = db.GetBucketData(bname1)
	_, _ = db.GetBucketData(bname2)

}
