package main

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)

func DumpToDisk(buckets map[string]Bucket) {
	db, err := bolt.Open("gcached.db", 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		log.Fatal(err)
	}
	/*
		// store some data
		err = db.Update(func(tx *bolt.Tx) error {
			for _, v := range buckets {
				boltBucket, err := tx.CreateBucketIfNotExists(v.key)
				if err != nil {
					return err
				}

				err = boltBucket.Put(key, value)
				if err != nil {
					return err
				}
			}

			return nil

		})
	*/
	db.Close()
}
