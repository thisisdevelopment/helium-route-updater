# Helium Router Updater

Automatically adds / removes euis from your helium routes based on updates from your LNS

## Status

**Beta**  
We use this internally and things may change (either on our side or on the helium side)  

## Purpose

Make it easy to integrate a new [Oui](https://docs.helium.com/use-the-network/run-a-network-server/buy-an-oui/) into Helium

## Requirements

- your own Helium [Oui](https://docs.helium.com/use-the-network/run-a-network-server/buy-an-oui/) 
- your own Helium DevAddr range
- your own LNS (eg: [Chirpstack](https://www.chirpstack.io/)) configured for helium:
  - with your own DevAddr range
  - allows connections from unknown gateways
  - your gateway bridge should be publicly accessible 
  - should store appeui somewhere as this is required for helium; but not natively supported by chirpstack
- a preconfigured route (see adding routes)
- enough credit in your wallet to "buy" messages

## Standalone usage

```
Currently not implemented
```

## Library usage

```go
ch := make(chan listener.DeviceEvent)

keypair := helium_crypto.KeyPairFromString(config.Auth)
go updater.Run(config.Server, config.RouteId, keypair, ch)

devEui, _ := strconv.ParseUint("<your device id>", 16, 64)
appEui, _ := strconv.ParseUint("<your app eui>", 16, 64)
ch <- listener.DeviceEvent{
    DevEui:    devEui,
    AppEui:    appEui,
    EventType: eventType,
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

rid=$(./helium-config-service-cli route new --commit  | cut -f3 -d' ')
./helium-config-service-cli route update server -r $rid --host $SERVER_HOST --port $SERVER_PORT --commit
./helium-config-service-cli route update add-gwmp-region -r $rid eu868 $SERVER_PORT --commit
```

# Contributing
You can help to improve this, check out how you can do things [CONTRIBUTING.md](CONTRIBUTING.md)

# License
© [This is Development BV](https://www.thisisdevelopment.nl), 2023~time.Now()
Released under the [MIT License](./LICENSE)
