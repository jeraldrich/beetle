# beetle
Beetle (WIP)
===================

High performance ETL implemented in golang concurrent fan-out pattern and scylla node clusters.

I was inspired by reading https://discord.com/blog/how-discord-stores-billions-of-messages

I ran into the same situation with mongo in that once you have around 2TB of data, the read performance was heavily impacted by writes. When I used sharding, I had dropped consistency issues - even when I adjusted the write concerns.

I did some research on cassandra and scylla as alternatives, but I found the scylladb query builder makes it easy to grasp for developers not familiar with column store databases: https://github.com/scylladb/gocqlx


## Setting Up Local Development Environment

#### Install Dependencies
* Install brew (osx package manager): http://brew.sh
* Install gpg (hash check all homebrew downloads): `brew install gpg`
* Install golang: `brew install golang`
* Install Docker: `brew install homebrew/cask/docker`
* Install scylla docker image: `docker pull scylladb/scylla`
* Start scylla node cluster: `docker compose up`
