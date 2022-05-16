Beetle (WIP)
===================

High performance ETL implemented in golang concurrent fan-out pattern and scylla node clusters.

I was inspired by reading https://discord.com/blog/how-discord-stores-billions-of-messages

I ran into the same situation with mongo in that after > 1TB+ of data, the read performance was heavily impacted by writes. When I used sharding, I had dropped consistency issues - even when I adjusted the write concerns.

I did some research on cassandra and scylla as alternatives, but I found the scylladb query builder makes it easy to grasp for developers not familiar with column store databases: https://github.com/scylladb/gocqlx

### ETL pipeline
* Messages are retrieved from sources defined in conf.json by a producer that downloads and parses files into data ready to be consumed
* The producer sends the messages data to a pool of consumer workers that will validate, transform and load the messages into a database
* Consumer workers clean the messages marking any that does not pass a filter with dirty_fields: true
* Consumer workers send the failed messages to the failed_imports table and successfull messages to the messages table (todo)

## Setting Up Local Development Environment

#### Install Dependencies
* Install golang: `brew install golang`
* Install Docker: `brew install homebrew/cask/docker`


### Start the scylla cluster

I created a docker compose file which will take the files in scylla_config and create 3 containerized scylla nodes networked with each other. The host machine can start a scylla session at 127.0.0.1:9042

* Start scylla node cluster: `docker compose up`

![ScreenShot](https://github.com/jeraldrich/beetle/blob/main/docker_scylla_cluster.png)


### Configure and run beetle
* Edit the conf.json file to add Urls to retrive messages from: `cp example_conf.json conf.json` 
* Run bettle: `go run .`
* If you want to compile a binary and run: `mkdir -p bin && go build -o bin/beetle && bin/beetle`

This is not necessary to run, but if you want to generate a go file to see all scylla's table metadata (columns, parkeys, sortkeys) do this:

### Generating table metadata (columns, partkeys, sortkeys) from scylla
* Install gocqlx's schemagen: `go get -u "github.com/scylladb/gocqlx/v2/cmd/schemagen"`
* If you have java, make sure you have $GOBIN defined in your path (location to your go bin directory) or you may run java's schemagen
* generate schema go file: `$GOBIN/schemagen -cluster="127.0.0.1:9042" -keyspace="messages" -output="schema" -pkgname="schema"`