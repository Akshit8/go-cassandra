# go-cassandra

## Starting a single-node cassandra cluster
```bash
docker-compose up -d

# getting cqlsh
docker exec -it cassandraDB cqlsh
```

## Creating your cassandra cluster schema
```bash
# create cassandra keyspace
CREATE KEYSPACE akshit WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};

# create table messages
use akshit;
create table messages (
id UUID,
user_id UUID,
Message text,
PRIMARY KEY(id)
);

# create table users
use akshit;
CREATE TABLE users (
id UUID,
firstname text,
lastname text,
age int,
email text,
city text,
PRIMARY KEY (id)
);
```

## Config Mgmt
```bash
cd config
cat sample.config.yml > config.yml

# populate the file with your variables
```

## Stream
Helps to Build scalable activity feeds.
<br>
https://getstream.io/activity-feeds/

## Makefile specs
- **git** - git add - commit - push commands
- **cassandra** - starts single node cassandra cluster on docker
- **server** - run the server

## References
[cassandra-setup](https://hub.docker.com/_/cassandra)<br>
[gocql](https://github.com/gocql/gocql)<br>
[cassandra-port-specs](https://stackoverflow.com/questions/2359159/cassandra-port-usage-how-are-the-ports-used)<br>
[cnfg-mgmt-go](https://medium.com/@bnprashanth256/reading-configuration-files-and-environment-variables-in-go-golang-c2607f912b63)<br>
[stream-demo](https://getstream.io/get_started/?signup=1#flat_feed)<br>
[stream-sdk-v2](https://github.com/getstream/stream-go2)<br>
[stream-feed](https://getstream.io/blog/building-a-performant-api-using-go-and-cassandra/)

## Author
**Akshit Sadana <akshitsadana@gmail.com>**

- Github: [@Akshit8](https://github.com/Akshit8)
- LinkedIn: [@akshitsadana](https://www.linkedin.com/in/akshit-sadana-b051ab121/)

## License
Licensed under the MIT License