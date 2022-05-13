# beetle
golang etl

go mod init beetledb/modules
go mod tidy

brew install golang
brew install homebrew/cask/docker
docker pull scylladb/scylla

Start a scylla server instance

$ docker exec -it beetledb nodetool status

Run cqlsh utility
$ docker exec -it beetledb cqlsh
Connected to Test Cluster at 172.17.0.2:9042.
[cqlsh 5.0.1 | Cassandra 2.1.8 | CQL spec 3.2.1 | Native protocol v3]
Use HELP for help.
cqlsh>

Make a cluster
$ docker run --name beetledb2  --hostname beetledb2 -d scylladb/scylla --seeds="$(docker inspect --format='{{ .NetworkSettings.IPAddress }}' beetledb)"

$ docker run --name beetledb --volume /var/lib/scylla:/var/lib/scylla -d scylladb/scylla --developer-mode=0

