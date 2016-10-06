# gobolt
Simplified interface for work with BoldDB database

##Roadmap

**DataBase**
 - [ ] Open database "Open"
 - [ ] Function to safe close database "Close"

**Bucket**
 - [ ] Create new bucket "CreateBucket"
 - [ ] Create indexed bucket "CreateIdBucket"
 - [ ] Get bucket info (indexed or not, use timestamp or not) "GetBucketInfo" (indexed: bool, usetime: bool).

**Object**
- [ ] Create new object in bucket or rewrite existing "SetByKey" (safe mode enabled)
- [ ] Get object by key "GetByKey"
- [ ] Update object in bucket with full rewrite or only filled vals "UpdateByKey" (safe mode enabled)
- [ ] Delete object in bucket "DeleteByKey"
- [ ] Get list of objects by value "GetListByVal"

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

 Coverage: 0%
