# go-cassandra

## Starting a single-node cassandra cluster
```bash
docker-compose up -d
```

## Setting your cassandra cluster
```bash
# create cassandra keyspace
CREATE KEYSPACE streamdemoapi WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};
```

## Makefile specs
- **git** - git add - commit - push commands
- **cassandra** - starts single node cassandra cluster on docker


## References
[cassandra-setup](https://hub.docker.com/_/cassandra)<br>
[gocql](https://github.com/gocql/gocql)<br>
[cassandra-port-specs](https://stackoverflow.com/questions/2359159/cassandra-port-usage-how-are-the-ports-used)<br>

## Author
**Akshit Sadana <akshitsadana@gmail.com>**

- Github: [@Akshit8](https://github.com/Akshit8)
- LinkedIn: [@akshitsadana](https://www.linkedin.com/in/akshit-sadana-b051ab121/)

## License
Licensed under the MIT License