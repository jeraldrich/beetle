package main

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

func main() {
	// Create gocql cluster.
	host := "127.0.0.1:9042"
	cluster := gocql.NewCluster(host)
	// cluster := gocqlxtest.CreateCluster()
	// Create gocql cluster.
	// Wrap session on creation, gocqlx session embeds gocql.Session pointer.
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		fmt.Printf("Failed to connect to cql cluster: [%s]", host)
	}
	defer session.Close()

	var releaseVersion string
	err = session.Query("SELECT release_version FROM system.local", nil).Get(&releaseVersion)
	if err != nil {
		fmt.Printf("Failed to query release_version from scylla: [%s]", host)
	}
}
