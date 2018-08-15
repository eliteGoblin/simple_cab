# simple_cab
simple cab trip counter


## Sample test script

```shell
curl -XPOST localhost:8080/v1/trip_info/2013-12-01/count -d@query.json -vvv
curl -XPOST localhost:8080/v1/trip_info/2013-12-01/count -d@query_nocache.json -vvv
curl -XPOST localhost:8080/v1/trip_info/2013-12-06/count -d@query2.json -vvv
curl -XPOST localhost:8080/v1/trip_info/2013-12-31/count -d@query2.json -vvv
curl -XPOST localhost:8080/v1/trip_info/2099-12-31/count -d@query2.json -vvv
curl -XPUT localhost:8080/v1/trip_info/update_cache -vvv
```


#### Further Thoughts

*  If cache miss, should then go to DB directly; if found, then fill the cache
*  Add cache miss rate counter (using atomic package), report it using statsd protocol to monitor it
*  UpdateCache should add wait queue, optimize when lots of updateCache cmd coming
*  Should add some unit tests or TDD (normally my code coverage rate is around 80% - 90%)
*  More goroutines to accelerate index built
*  Put it in docker