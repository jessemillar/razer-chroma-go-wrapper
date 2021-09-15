# Razer Chroma HTTP Wrapper

[![Go Report Card](https://goreportcard.com/badge/github.com/jessemillar/razer-chroma-http-wrapper)](https://goreportcard.com/report/github.com/jessemillar/razer-chroma-http-wrapper) [![Man Hours](https://img.shields.io/endpoint?url=https%3A%2F%2Fmh.jessemillar.com%2Fhours%3Frepo%3Dhttps%3A%2F%2Fgithub.com%2Fjessemillar%2Frazer-chroma-http-wrapper.git)](https://jessemillar.com/r/man-hours)

The Razer Chroma HTTP Wrapper is a system tray application that exists to wrap the (rather confusing) [Razer Chroma REST SDK](https://assets.razerzone.com/dev_portal/REST/html/index.html) to enable easier scripting. The HTTP Wrapper exposes localhost HTTP endpoints that can be queried from scripts or applications to cause your Razer peripherals to flash and notify you of an event.

## Quickstart

> TODO Provide links to the binary

1. Download the binary
1. Run the binary
1. Hit one of the endpoints listed below

## Endpoints

> All endpoints are HTTP GET requests for scripting simplicity. Params embedded in the URI are required, anything after `?` are optional.

```
https://localhost:<port>/color/<hex-color>
https://localhost:<port>/flash/color/<hex-color>?count=3&duration=1000&interval=500
```

## Configuration

Configuration is handled by the [`config.toml`](./config.toml) file. This file needs to be in the same directory as the `razer-chroma-http-wrapper` executable. Supported configuration values are defined below:

- `server_port`: An integer value specifying which port to listen on for HTTP GET requests
- `default_color`: A hex color value that the Razer peripheral defaults to when not currently displaying an alert

## Compatibility

### Operating System

My use case is Windows 10 with the [Razer Goliathus Chroma Extended](https://www.razer.com/gaming-mouse-mats/Razer-Goliathus-Chroma/RZ02-02500300-R3M1) mousepad. The Razer Chroma HTTP Wrapper may work with other OSs and other peripherals, but it is not guaranteed. I'm happy to accept PRs to add desired functionality.

### Razer Chroma SDK

Please note, the Razer Chroma HTTP Wrapper was written for this version of the Chroma SDK. It may work with others, but has not been tested.

```
{
    "core": "3.20.02",
    "device": "3.20.02",
    "version": "3.20.03"
}
```
