//+build wireinject

package main

import (
	"github.com/FictProger/architecture-lab-3/server/dormitories"
	"github.com/google/wire"
)

// ComposeApiServer will create an instance of CharApiServer according to providers defined in this file.
func ComposeApiServer(port HttpPortNumber) (*DirectionApiServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		NewDbConnection,
		// Add providers from dormitories package.
		dormitories.Providers,
		// Provide DirectionApiServer instantiating the structure and injecting dormitories handler and port number.
		wire.Struct(new(DirectionApiServer), "Port", "DormitoriesHandler"),
	)
	return nil, nil
}
