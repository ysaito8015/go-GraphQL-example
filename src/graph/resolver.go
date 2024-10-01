package graph

import (
	"mygraphql/graph/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	// This resolver is used to resolve the query and mutation fields.
	Srv services.Services
}
