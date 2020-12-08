package dormitories

import "github.com/google/wire"

// Set of providers for channels components.
var Providers = wire.NewSet(NewDormitories, HttpHandler)
