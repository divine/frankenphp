// Copied from https://github.com/caddyserver/xcaddy/blob/b7fd102f41e12be4735dc77b0391823989812ce8/environment.go#L251
package main

import (
	"github.com/caddyserver/caddy/v2"
	caddycmd "github.com/caddyserver/caddy/v2/cmd"
	"go.uber.org/automaxprocs/maxprocs"

	// plug in Caddy modules here.
	_ "github.com/caddyserver/caddy/v2/modules/standard"
	_ "github.com/dunglas/frankenphp/caddy"
	_ "github.com/dunglas/mercure/caddy"
	_ "github.com/dunglas/vulcain/caddy"
	_ "github.com/RussellLuo/caddy-ext/ratelimit"
)

func init() {
	//nolint:errcheck
	maxprocs.Set(maxprocs.Logger(caddy.Log().Sugar().Debugf))
}

func main() {
	caddycmd.Main()
}
