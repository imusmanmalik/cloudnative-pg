/*
This file is part of Cloud Native PostgreSQL.

Copyright (C) 2019-2021 EnterpriseDB Corporation.
*/

// Package url holds the constants for webserver routing
package url

import (
	"fmt"
)

const (
	// StatusPort is the port for status HTTP requests
	StatusPort int = 8000

	// MetricsPort is the port for HTTP requests
	MetricsPort int = 9187

	// PathHealth is the URL path for Health State
	PathHealth string = "/healthz"

	// PathReady is the URL oath for Ready State
	PathReady string = "/readyz"

	// PathPgStatus is the URL path for PostgreSQL Status
	PathPgStatus string = "/pg/status"

	// PathPgBackup is the URL path for PostgreSQL Backup
	PathPgBackup string = "/pg/backup"

	// PathMetrics is the URL path for Metrics
	PathMetrics string = "/metrics"
)

// Local builds an url for the provided path on localhost, pointing to the status web server
func Local(path string) string {
	return Build("localhost", path)
}

// Build builds an url given the hostname and the path, pointing to the status web server
func Build(hostname, path string) string {
	// If path already starts with '/' we remove it
	if path[0] == '/' {
		path = path[1:]
	}
	return fmt.Sprintf("http://%s:%d/%s", hostname, StatusPort, path)
}