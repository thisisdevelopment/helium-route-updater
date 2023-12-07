package main

import (
	"github.com/thisisdevelopment/helium-route-updater/cmd/helium-utils/cli"
	"os"
)

func main() {
	err := cli.RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

/**
 * Commands to implement:
 * - check --all <= by default only passive checks
 * ✓ check config
 * - check lns config <= how to get regions / devaddrs / netid / allow_unknown_gateways
 * - check lns connection
 * - check helium auth (validate oui matches delegate keypair, make sure not to use owner keypair)
 * - check helium route
 * - check sync <= checks if route is in sync with lns
 * - check ingress (public ip check + connection check + create device + per region: simulate join + simulate uplink => validate devaddr/netid, if join fails; check if it works when adding the gateway)
 * ✓ keypair info
 * ✓ keypair private
 * ✓ keypair public
 * - helium route create <= how to get regions / devaddrs?
 * - helium route info
 * - helium org info
 * - helium org create?
 * - helium org balance
 * - helium gateway lookup
 * - helium info
 * - lns info
 * - lns device create
 * - lns device join
 * - lns device uplink
 * - lns device delete
 * - lns gateway create
 * - lns gateway delete
 * - sync <= runs a single sync
 * - info
 */
