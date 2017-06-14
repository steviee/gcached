package service

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/steviee/gcached/store"
)

/*
  Bucket handlers

  GET /              - returns list of buckets
  GET /<bucket>      - items of bucket, defaultTTL, createdAt
  POST,PUT /<bucket> - creates or updates a bucket, defaultTTL = 0 if omitted
  DELETE /<bucket>   - delete bucket with all items
*/

/*
  Item handlers

  GET /<bucket>/<key>    - get item value and refresh ttl or return status 404 if not found
  PUT /<bucket>/<key>    - update item value and refresh ttl (create implicitly)
  POST /<bucket>/<key>   - create item value and refresh ttl
  DELETE /<bucket>/<key> - delete item
*/

// BucketsIndex returns a list of all buckets as array
func BucketsIndex(w http.ResponseWriter, r *http.Request) {

	// get bucket keys
	list := make([]string, 0)
	for _, v := range store.Buckets {
		list = append(list, v.Key)
	}

	// just sort the list
	sort.Strings(list)

	// send them out!
	if err := json.NewEncoder(w).Encode(list); err != nil {
		panic(err)
	}

}

// BucketCreate creates the bucket
func BucketCreate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bucketKey := vars["key"]
	var ttl int

	ttl, err := strconv.Atoi(r.URL.Query().Get("ttl"))
	if err != nil {
		ttl = 0
	}

	bucket := store.Bucket{Key: bucketKey, DefaultTTL: ttl, Items: make(map[string]store.Item), CreatedAt: time.Now()}
	store.Buckets[bucketKey] = bucket

	w.WriteHeader(http.StatusCreated)
}

// BucketDelete deletes the bucket
func BucketDelete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bucketKey := vars["key"]

	delete(store.Buckets, bucketKey)

	w.WriteHeader(http.StatusOK)
}

// BucketIndex returns information about the requested bucket (or 404 if non-existant)
func BucketIndex(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bucketKey := vars["key"]

	bucket, ok := store.Buckets[bucketKey]

	if ok {
		if err := json.NewEncoder(w).Encode(bucket); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func ItemSet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bucketKey := vars["key"]
	itemKey := vars["itemKey"]
	var ttl int

	ttl, err := strconv.Atoi(r.URL.Query().Get("ttl"))
	if err != nil {
		ttl = 0
	}

	bucket, ok := store.Buckets[bucketKey]
	if ok {
		item := store.Item{Key: itemKey, Value: "x", CreatedAt: time.Now(), ProlongedAt: time.Now(), TTL: ttl}
		bucket.Items[itemKey] = item
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func ItemDelete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bucketKey := vars["key"]
	itemKey := vars["itemKey"]
	bucket, ok := store.Buckets[bucketKey]

	if ok {
		delete(bucket.Items, itemKey)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func ItemShow(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bucketKey := vars["key"]
	itemKey := vars["itemKey"]

	bucket, ok := store.Buckets[bucketKey]
	if ok {
		item, itemOk := bucket.Items[itemKey]
		if itemOk {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				panic(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
