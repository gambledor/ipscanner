# Network ip scanner

It scans the network looking for connected devices.
It takes the ip read from the network interface given in input and scans the network.
It has to be run as privileged user because of kernel capabilities.

## How to compile

```bash
go build -o ipScanner
```

## How to run

```bash
sudo ipScanner -i <network_interface>
```

for example

```bash
sudo ipScanner -i eth0
```
