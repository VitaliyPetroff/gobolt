package main

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	//"github.com/gorilla/schema"
)

// Data base location settings
type GoBoltConf struct {
	Path string
}

// Gobolt struct
type Gobolt struct {
	gb *bolt.DB
}

// Bucket data struct
type BucketData struct {
	Key  string
	Data interface{}
}

// Open database
func (gbase *Gobolt) Open(gbc GoBoltConf) error {
	var err error

	// ToDo: in safe mode don't create database if it's not exist.
	gbase.gb, err = bolt.Open(gbc.Path, 0644, nil)

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
	var bucket_data_list []BucketData

	err := gbase.gb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket_name))
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", bucket_name)
		}

		err := bucket.ForEach(func(k, v []byte) error {
			var bucket_data BucketData
			bucket_data.Key = string(k[:])

			err := json.Unmarshal(v, &bucket_data.Data)
			if err != nil {
				return err
			}

			bucket_data_list = append(bucket_data_list, bucket_data)
			return nil
		})

		return err
	})

	return bucket_data_list, err
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
func (gbase *Gobolt) GetByKey(bucket_name string, key_name string, obj interface{}) error {
	err := gbase.gb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket_name))
		if bucket == nil {
			return fmt.Errorf("Bucket %s not found!", bucket_name)
		}
		err := bucket.ForEach(func(k, v []byte) error {
			err := json.Unmarshal(v, &obj)
			if err != nil {
				return err
			}

			return nil
		})

		return err
	})

	return err
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

	gbc.Path = "testdb.bdb"

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

	// var list interface{}
	// list, err = db.GetBucketList()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println("readed from db:")
	// 	fmt.Println(list)
	// }

	// var bucket_data interface{}
	// bucket_data, err = db.GetBucketData(bname1)
	// if err != nil {
	// 	fmt.Println("error: " + err.Error())
	// } else {
	// 	fmt.Printf("bucket1: %q\n", bucket_data)
	// }

	// bucket_data, err = db.GetBucketData(bname2)
	// if err != nil {
	// 	fmt.Println("error: " + err.Error())
	// } else {
	// 	fmt.Printf("bucket2: %q\n", bucket_data)
	// }

	var mo myObject
	err = db.GetByKey(bname1, "mo4", &mo)
	if err != nil {
		fmt.Println("error: " + err.Error())
	} else {
		fmt.Printf("GetByKey: mo4 ==  %q\n", mo)
		fmt.Printf("mo.F1: %s\n", mo.F1)
		fmt.Printf("mo.F2: %s\n", mo.F2)
	}

}
