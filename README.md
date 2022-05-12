# beetle
golang etl

go mod init beetledb/modules
go mod tidy

brew install golang
brew install homebrew/cask/docker
docker pull scylladb/scylla

Start a scylla server instance

$docker run --name beetledb --hostname beetledb -d scylladb/scylla --smp 1

 This command will start a Scylla single-node cluster in developer mode (see --developer-mode 1) limited by a single CPU core (see --smp). Production grade configuration requires tuning a few kernel parameters such that limiting number of available cores (with --smp 1) is the simplest way to go.

Multiple cores requires setting a proper value to the /proc/sys/fs/aio-max-nr. On many non production systems it will be equal to 65K. The formula to calculate proper value is:

Available AIO on the system - (request AIO per-cpu * ncpus) =
aio_max_nr - aio_nr < (reactor::max_aio + detect_aio_poll + reactor_backend_aio::max_polls) * cpu_cores =
aio_max_nr - aio_nr < (1024 + 2 + 10000) * cpu_cores =
aio_max_nr - aio_nr < 11026 * cpu_cores

where

reactor::max_aio = max_aio_per_queue * max_queues,
max_aio_per_queue = 128,
max_queues = 8.

$ docker exec -it beetledb nodetool status

Run cqlsh utility
$ docker exec -it beetledb cqlsh
Connected to Test Cluster at 172.17.0.2:9042.
[cqlsh 5.0.1 | Cassandra 2.1.8 | CQL spec 3.2.1 | Native protocol v3]
Use HELP for help.
cqlsh>

Make a cluster
$ docker run --name beetledb2  --hostname beetledb2 -d scylladb/scylla --seeds="$(docker inspect --format='{{ .NetworkSettings.IPAddress }}' beetledb)"

$ docker exec -it beetledb supervisorctl restart scylla

To improve I/O performance (not necessery for dev environment unless stress testing):
Create a Scylla data directory /var/lib/scylla on the host, which is used by Scylla container to store all data:

$ sudo mkdir -p /var/lib/scylla/data /var/lib/scylla/commitlog /var/lib/scylla/hints /var/lib/scylla/view_hints
Launch Scylla using Docker's --volume command line option to mount the created host directory as a data volume in the container and disable Scylla's developer mode to run I/O tuning before starting up the Scylla node.

$ docker run --name beetledb --volume /var/lib/scylla:/var/lib/scylla -d scylladb/scylla --developer-mode=0

