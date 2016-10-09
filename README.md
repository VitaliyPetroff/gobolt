# gobolt

[![Build Status](https://drone.io/github.com/VitaliyPetroff/gobolt/status.png)](https://drone.io/github.com/VitaliyPetroff/gobolt/latest)
**Do not use! It's in development stage!**

Simplified interface for work with BoldDB database

##Roadmap

**DataBase**
 - [x] Open database "Open"
 - [x] Function to safe close database "Close"
 - [x] Get list of all bucket "GetBucketList"
 - [ ] Get database struct

**Bucket**
 - [x] Create new bucket "CreateBucket"
 - [ ] Create indexed bucket "CreateIdBucket"
 - [ ] Get bucket info (indexed or not, use timestamp or not) "GetBucketInfo" (indexed: bool, usetime: bool)
 - [x] Get bucket data

**Object**
- [x] Create new object in bucket or rewrite existing "SetByKey" (safe mode enabled)
- [x] Get object by key "GetByKey"
- [ ] Update object in bucket with full rewrite or only filled vals "UpdateByKey" (safe mode enabled)
- [ ] Delete object in bucket "DeleteByKey"
- [ ] Get list of objects by value "GetListByVal"
- [ ] Get list of objects by using 'between search'
- [ ] Get list of objects by using 'less or more search'

**Indexed Object**
- [ ] Create new object in bucket and get new id "SetWithNewId"

**Other**
 - [ ] Safe mode for work with database
 - [ ] Limit for objects list size
 - [ ] Sort list of objects by key
 - [ ] Sort list of objects by time
 - [ ] Make tests
 - [ ] Create documentation with examples
 - [ ] Something else

 Coverage: 30%
