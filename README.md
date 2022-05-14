Beetle (WIP)
===================

High performance ETL implemented in golang concurrent fan-out pattern and scylla node clusters.

I was inspired by reading https://discord.com/blog/how-discord-stores-billions-of-messages

I ran into the same situation with mongo in that once you have around 2TB of data, the read performance was heavily impacted by writes. When I used sharding, I had dropped consistency issues - even when I adjusted the write concerns.

I did some research on cassandra and scylla as alternatives, but I found the scylladb query builder makes it easy to grasp for developers not familiar with column store databases: https://github.com/scylladb/gocqlx

## Setting Up Local Development Environment

#### Install Dependencies
* Install golang: `brew install golang`
* Install Docker: `brew install homebrew/cask/docker`

### Running Beetle
* Start scylla node cluster: `docker compose up`
* Configure which sources your producers will fetch from in producers.cfg
* Run bettle: `go run .`
* If you want to compile a binary and run: `mkdir -p bin && go build -o bin/beetle && bin/beetle`

This is not necessary to run, but if you want to generate a schema (outputs a go file) to see all metadata (columns, parkeys, sortkeys) do this:
### Generating table metadata (columns, partkeys, sortkeys) from scylla
* go get -u "github.com/scylladb/gocqlx/v2/cmd/schemagen"
* $GOBIN/schemagen -cluster="127.0.0.1:9042" -keyspace="messages" -output="schema" -pkgname="schema"

Make sure you have $GOBIN defined in your path or you may run java's schemagen if you have java on your system.