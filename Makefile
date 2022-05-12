default: setup start
setup: install

install:
	@docker run --name beetledb --hostname beetledb -d scylladb/scylla --smp 1
	@sleep 15
	@docker run --name beetledb2  --hostname beetledb2 -d scylladb/scylla --seeds="$(docker inspect --format='{{ .NetworkSettings.IPAddress }}' beetledb)"
	@sleep 15
	@docker exec -it beetledb nodetool status

clean:
	@echo 'clean task not implemented'

build:
	@echo 'build task not implemented'

start:
	@docker exec -it beetledb supervisorctl restart scylla

restart:
	@docker exec -it beetledb supervisorctl restart scylla

console:
	@docker exec -it beetledb cqlsh

test:
	@go test