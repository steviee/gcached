[![CircleCI](https://circleci.com/gh/steviee/gcached/tree/master.svg?style=svg)](https://circleci.com/gh/steviee/gcached/tree/master)

# gcached
**gcached** (like memcached, but with indexable buckets and disk dumping) is a simple project to backup some thoughts about a frequency-capping problem for an ad-delivery-pipeline.

**The problem (abstract)**

I need to do frequent checking against some datastore for recently access items on a per-user base. Items (ads) are stored by ID in the storage. Additionally I need to cap on different intervals like once every minute, 7 times a day or 20 times a week (for example).

For this I would create a bucket per user (user_1) to query for keys like 

```ruby
GET baseurl://user_1/?touched_within=60.seconds (or)
GET baseurl://user_1/?touched_within=24.hours (or)  
GET baseurl://user_1/?touched_within=7.days
```

The result should always include all the items that where created (or updated) within the given timespan. 

Items in the bucket could simply contain IDs and a ttl, added like this:  

```ruby
POST baseurl://user_1/12412?ttl=7.days  
POST baseurl://user_1/93183?ttl=7.days  
POST baseurl://user_1/11312?ttl=7.days  
```

They can either be added by POST url or with a body containing a JSON-document i.e.

Every once-in-a-while a bucket-cleaner should go over about-to-be-expired items and remove them from the respective bucket. 
Also to prevent data-loss an OPLOG should be written to recreate the data in case of a catastrophic data-loss. Normal operation should include some data-dumping for example with [boltdb](https://github.com/boltdb/bolt) to save and re-load the buckets when restarting the service.

This is work-in-progress, so be aware!
