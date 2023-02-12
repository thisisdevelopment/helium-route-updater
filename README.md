# Helium Router Updater

Automatically adds / removes euis from your helium routes based on updates from your LNS (currently only chirpstack is directly supported)

## Status

Alpha! First prototype and not yet fully working

Currently, as of feb 12 2023, Helium does not support this in production.   
This depends on all the changes that they are doing for [HIP70](https://github.com/helium/HIP/blob/main/0070-scaling-helium.md)

## Purpose

Make it easy to integrate a new [Oui](https://docs.helium.com/use-the-network/run-a-network-server/buy-an-oui/) into Helium

## Requirements

- your own Helium [Oui](https://docs.helium.com/use-the-network/run-a-network-server/buy-an-oui/) 
- your own Helium DevAddr range
- your own LNS (eg: [Chirpstack](https://www.chirpstack.io/)) configured for helium:
  - with your own DevAddr range
  - allows connections from unknown gateways
  - your gateway bridge should be publicly accessible 
- a preconfigured route 
- enough credit in your wallet to "buy" messages

## Usage

You can use this in 2 ways:
- as a standalone daemon that listens to device updates from the specified lns (lora network server); for now it only support chirpstack
- as a lib; either use the included listeners to listen for device updates or roll your own and integrate it in your own program 

### Standalone usage

```
//TODO
```

### Library usage

```
//TODO
```

# Contributing
You can help to improve this, check out how you can do things [CONTRIBUTING.md](CONTRIBUTING.md)

# License
© [This is Development BV](https://www.thisisdevelopment.nl), 2023~time.Now()
Released under the [MIT License](./LICENSE)
