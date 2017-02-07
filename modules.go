package main

import (
	_ "github.com/amaysim-au/logspout/adapters/raw"
	_ "github.com/amaysim-au/logspout/adapters/syslog"
	_ "github.com/amaysim-au/logspout/httpstream"
	_ "github.com/amaysim-au/logspout/routesapi"
	_ "github.com/amaysim-au/logspout/transports/tcp"
	_ "github.com/amaysim-au/logspout/transports/udp"
	_ "github.com/amaysim-au/logspout/transports/tls"
	_ "github.com/amaysim-au/logspout-http/http"
)
