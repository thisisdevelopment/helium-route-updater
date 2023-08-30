# Helium Router Updater

Automatically adds / removes euis / skfs from your helium routes based on updates from your LNS

## Status

**Beta**  
We use this internally and things may change (either on our side or on the helium side)

## Purpose

Make it easy to integrate your own LNS into Helium

## Requirements

- your own Helium [Oui](https://docs.helium.com/use-the-network/run-a-network-server/buy-an-oui/)
- your own Helium DevAddr range
- your own LNS (eg: [Chirpstack](https://www.chirpstack.io/)) configured for helium:
    - with your own DevAddr range
    - allows connections from unknown gateways
    - your gateway bridge should be publicly accessible
    - should store appEui somewhere as this is required for helium ~~but not natively supported by chirpstack~~ this is
      possible as of chirpstack 4.4; it is called joinEui
- a preconfigured route (see adding routes)
- enough DC credit in your wallet to "buy" messages (you can use https://dc-portal.helium.com/ to buy DC credits and
  view the current DC credits of your oui)

## Building

```bash
docker-compose run --entrypoint "" helium-route-updater ./build.sh
```

## Standalone usage

Make sure the following environment variables are set (via in a `.env` file or via actual environment variables)

```
LNS_TYPE=chirpstack
LNS_API_AUTH=<tenantid>:<apikey>
LNS_API_ENDPOINT=<proto>://<chirpstack-server-hostname>
LNS_LISTEN=redis://<redis-server-hostname>:<port>/<dbnum>
LNS_AUTO_ROAMING=<boolean>
HELIUM_ROUTE_ID=<route uuid>
HELIUM_KEYPAIR=<hex encoded keypair or file containing the keypair (should start with ./ or /)>
HELIUM_SERVER=https://config.iot.mainnet.helium.io:6080
```

Run the `./bin/helium-route-updater` command or run
the [thisisdeveloment/helium-route-updater](https://hub.docker.com/repository/docker/thisisdevelopment/helium-route-updater/general)
docker container

The only odd one here is `LNS_AUTO_ROAMING`; it will update the region of the linked device profile of a device if it
detects a different region in the `JOIN_REQUEST/REJOIN_REQUEST`. This one is not helium specific; but as helium
is one of the few global LoRaWAN networks and it fits easily in the rest of the structure it is an optional feature here.

**Caution**: The `LNS_AUTO_ROAMING` feature currently assumes each device has it own device profile. If you have a setup 
were you share device profiles between multiple devices this wil break things!! 

## Chirpstack Docker Compose usage

1. Copy your delegate keypair to the same directory as your docker-compose.yml
2. Add the below services definition to your docker-compose.yml.
3. Fill in the missing <> environment variable values.
4. Run `docker-compose down && docker-compose up -d`

```
helium-route-updater:
    image: thisisdevelopment/helium-route-updater
    environment:
      - LNS_TYPE=chirpstack
      - LNS_API_AUTH=<chirpstack apikey>
      - LNS_API_ENDPOINT=http://chirpstack:8080
      - LNS_LISTEN=redis://redis:6379/0
      - LNS_AUTO_ROAMING=false
      - HELIUM_ROUTE_ID=<route uuid>
      - HELIUM_KEYPAIR=./delegate_keypair.bin
      - HELIUM_SERVER=https://config.iot.mainnet.helium.io:6080
    depends_on:
      - chirpstack
    volumes:
      - ./delegate_keypair.bin:/app/delegate_keypair.bin
```

## Library usage

Please use the ssl/tls endpoint for production: https://config.iot.mainnet.helium.io:6080
See https://github.com/helium/oracles/issues/543

```go
ch := make(chan listener.DeviceEvent)
keypair := helium_crypto.KeyPairFromString(config.Auth)
client := helium_api.NewClient(config.Server, keypair)
go updater.Run(client, config.RouteId, ch)

devEui, _ := strconv.ParseUint("<your device id>", 16, 64)
joinEui, _ := strconv.ParseUint("<your app eui>", 16, 64)
device = &types.Device{devEui: devEui, joinEui: joinEui}
ch <- types.DeviceEvent{
Update: []*types.Device{device},
Delete: []*types.Device{},
}

```

## Adding routes

### Step 1: build helium-config-service-cli

This is used to create a new route for your oui

```bash
git clone https://github.com/helium/helium-config-service-cli ./source
docker run -v $(pwd)/source:/app -w /app rust:1.70 bash -c "apt-get update && apt-get install -y protobuf-compiler && cargo install --path ."
sudo mv source/target/release/helium-config-service-cli ./
sudo chown $(id -u):$(id -g) ./helium-config-service-cli
```

#### Step 2: create route

For this you'll need the info from when you created your oui.
Change your regions/ports where needed. For chirpstack you'll need a separate port per region.

```base
export HELIUM_KEYPAIR_BIN=<path to your private (delegate) key file>
export HELIUM_NET_ID=<your netid>
export HELIUM_OUI=<your oui>
export HELIUM_MAX_COPIES=5
export SERVER_HOST=<your public LNS server hostname>
export SERVER_PORT=<your public LNS server port>
export DEVADDR_START=<start of your devaddr range>
export DEVADDR_END=<end of your devaddr range> 

rid=$(./helium-config-service-cli route new --commit  | head -n 1 | cut -f3 -d' ')
./helium-config-service-cli route update server -r $rid --host $SERVER_HOST --port $SERVER_PORT --commit
./helium-config-service-cli route update add-gwmp-region -r $rid eu868 $SERVER_PORT --commit
./helium-config-service-cli route devaddrs add --route-id $rid -s $DEVADDR_START -e $DEVADDR_END --commit
```

# Contributing

You can help to improve this, check out how you can do things [CONTRIBUTING.md](CONTRIBUTING.md)

# License

© [This is Development BV](https://www.thisisdevelopment.nl), 2023~time.Now()
Released under the [MIT License](./LICENSE)
