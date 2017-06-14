package store

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

func StartBackgroundDump() chan bool {
	stop := make(chan bool)
	go func() {
		for {
			DumpToDisk()
			select {
			case <-time.After(60 * time.Second):
			case <-stop:
				return
			}
		}
	}()

	return stop
}

func DumpToDisk() {

	log.Println("Dumping data to disk...")

	if exists("gcached.db") {
		err := os.Remove("gcached.db")
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	db, err := bolt.Open("gcached.db", 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		log.Fatal(err)
	}

	// store some data
	err = db.Update(func(tx *bolt.Tx) error {
		for _, bucket := range buckets {
			boltBucket, err := tx.CreateBucketIfNotExists([]byte(bucket.Key))
			if err != nil {
				return err
			}

			b, err := json.Marshal(bucket)
			if err != nil {
				fmt.Println("error:", err)
			}

			err = boltBucket.Put([]byte(bucket.Key), b)
			if err != nil {
				return err
			}
		}

		return nil

	})
	db.Close()

	//fi, _ := os.Stat("gcache.db")
	//fmt.Println(fmt.Sprintf("%d bytes written to disk. ", fi.Size()))

}

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
