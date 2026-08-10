# Linux Networking & Connectivity

This document covers common Linux networking commands, concepts, connectivity checks, DNS troubleshooting, ports, sockets, HTTP requests, and basic network troubleshooting.

---

## 1. Understanding Linux Networking

Linux networking allows a system to communicate with:

- Other computers on a local network
- Servers on the internet
- Containers
- Virtual machines
- Network services running on the same machine

Some important networking concepts are:

- IP address
- Network interface
- MAC address
- Port
- Protocol
- DNS
- Gateway
- Routing
- TCP
- UDP
- Socket

---

# 2. Network Interfaces

A network interface represents a connection between the system and a network.

Common interfaces include:

- `eth0` - Ethernet interface
- `enp3s0` - Modern Ethernet naming
- `wlan0` - Wireless interface
- `wlp2s0` - Modern wireless naming
- `lo` - Loopback interface

The loopback interface refers to the local machine itself.

The commonly used loopback address is:

```text
127.0.0.1
````

---

## 3. ip

The `ip` command is the primary modern command for inspecting and configuring Linux networking.

### Show network interfaces

```bash
ip link
```

Example:

```text
1: lo: <LOOPBACK,UP,LOWER_UP>
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP>
```

### Show IP addresses

```bash
ip addr
```

or:

```bash
ip a
```

Example:

```text
inet 192.168.1.20/24
```

This means the interface has the IP address:

```text
192.168.1.20
```

with a `/24` network prefix.

### Show a specific interface

```bash
ip addr show eth0
```

### Show routing table

```bash
ip route
```

Example:

```text
default via 192.168.1.1 dev eth0
192.168.1.0/24 dev eth0
```

The `default` route is normally used when traffic is destined for a network that does not have a more specific route.

### Show neighbors

```bash
ip neigh
```

This displays information about devices known through the local network's ARP/neighbor discovery mechanisms.

---

# 4. IP Addresses

An IP address identifies a network interface on an IP network.

IPv4 example:

```text
192.168.1.100
```

IPv6 example:

```text
2001:db8::1
```

Private IPv4 ranges commonly include:

```text
10.0.0.0/8
172.16.0.0/12
192.168.0.0/16
```

These addresses are normally used inside private networks.

---

# 5. localhost

`localhost` refers to the current computer.

Common addresses are:

```text
127.0.0.1
```

for IPv4 and:

```text
::1
```

for IPv6.

For example:

```bash
curl http://localhost:8080
```

attempts to connect to a service running on the local machine on port `8080`.

---

# 6. ping

`ping` tests whether another host is reachable over the network and measures round-trip time.

Example:

```bash
ping google.com
```

Ping a specific IP:

```bash
ping 8.8.8.8
```

Send a limited number of packets:

```bash
ping -c 4 google.com
```

The `-c` option specifies the number of packets to send.

Example output:

```text
64 bytes from ...
time=20.3 ms
```

The `time` value represents approximately how long the packet took to travel to the destination and back.

### Important

A failed `ping` does not always mean that a server is unreachable.

Some systems or firewalls block ICMP packets while still allowing TCP or HTTP connections.

---

# 7. curl

`curl` is a command-line tool for transferring data over network protocols.

It is especially useful for testing HTTP APIs and web servers.

### Make an HTTP request

```bash
curl https://example.com
```

### Show response headers

```bash
curl -I https://example.com
```

### Show detailed connection information

```bash
curl -v https://example.com
```

### Follow redirects

```bash
curl -L https://example.com
```

### Send a POST request

```bash
curl -X POST https://example.com/api
```

### Send JSON

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"name":"Almas"}' \
  http://localhost:8080/api
```

### Download a file

```bash
curl -O https://example.com/file.zip
```

### Specify a timeout

```bash
curl --connect-timeout 5 https://example.com
```

---

# 8. wget

`wget` is primarily used for downloading files from the network.

Example:

```bash
wget https://example.com/file.zip
```

Save the file with a specific name:

```bash
wget -O file.zip https://example.com/download
```

Continue an interrupted download:

```bash
wget -c https://example.com/file.zip
```

Download recursively:

```bash
wget -r https://example.com/
```

Be careful with recursive downloads because they can download a large amount of data.

---

# 9. curl vs wget

Both can download files.

`curl` is particularly useful for:

* HTTP requests
* APIs
* Sending headers
* Sending JSON
* Testing web services
* Debugging HTTP

`wget` is particularly useful for:

* Downloading files
* Resuming downloads
* Recursive downloads

---

# 10. Ports

A port identifies a particular network service on a host.

Examples:

```text
22    SSH
53    DNS
80    HTTP
443   HTTPS
5432  PostgreSQL
3306  MySQL
6379  Redis
8080  Common application port
```

An IP address identifies the host.

A port identifies the service on that host.

For example:

```text
192.168.1.10:8080
```

means:

```text
Host: 192.168.1.10
Port: 8080
```

---

# 11. ss

`ss` displays socket information.

It is commonly used to inspect listening ports and network connections.

### Show listening TCP sockets

```bash
ss -ltn
```

Meaning:

```text
-l    listening
-t    TCP
-n    don't resolve names
```

### Show listening TCP and UDP sockets

```bash
ss -ltnu
```

### Show all TCP connections

```bash
ss -tn
```

### Show the process using a socket

```bash
sudo ss -ltnp
```

Example:

```text
LISTEN 0 128 0.0.0.0:8080
```

This means something is listening on port `8080`.

---

# 12. Checking Whether a Port Is Open

To check whether a local service is listening:

```bash
ss -ltn
```

To filter for a specific port:

```bash
ss -ltn | grep :8080
```

You can also use `curl` when the service speaks HTTP:

```bash
curl http://localhost:8080
```

A connection failure can indicate:

* Service is not running
* Service is listening on another port
* Service is listening only on another interface
* Firewall is blocking the connection
* Incorrect IP address
* Incorrect protocol

---

# 13. TCP and UDP

TCP and UDP are transport-layer protocols.

## TCP

TCP provides a connection-oriented communication mechanism.

It provides features such as:

* Reliable delivery
* Ordering
* Retransmission
* Connection management

Common TCP applications include:

```text
HTTP
HTTPS
SSH
PostgreSQL
```

## UDP

UDP is connectionless and has lower protocol overhead.

It does not provide TCP's built-in guarantees of reliable ordered delivery.

Common UDP uses include:

```text
DNS
DHCP
Streaming
Some real-time applications
```

---

# 14. DNS

DNS stands for Domain Name System.

DNS translates domain names into IP addresses.

For example:

```text
example.com
```

may resolve to an IP address such as:

```text
93.184.216.34
```

Without DNS, users would generally need to remember IP addresses instead of domain names.

---

# 15. dig

`dig` is a DNS troubleshooting tool.

Look up a domain:

```bash
dig example.com
```

Show only the answer:

```bash
dig +short example.com
```

Query a specific record type:

```bash
dig example.com A
```

IPv6 address:

```bash
dig example.com AAAA
```

Mail records:

```bash
dig example.com MX
```

Name servers:

```bash
dig example.com NS
```

Query a specific DNS server:

```bash
dig @8.8.8.8 example.com
```

---

# 16. nslookup

`nslookup` can also be used to query DNS.

Example:

```bash
nslookup example.com
```

Query a specific DNS server:

```bash
nslookup example.com 8.8.8.8
```

`dig` is generally preferred for detailed DNS troubleshooting, while `nslookup` is still widely available and useful for basic queries.

---

# 17. DNS Troubleshooting

If a domain does not work, first determine whether DNS resolution works.

Try:

```bash
ping example.com
```

or:

```bash
dig +short example.com
```

If the domain does not resolve, investigate DNS.

You can compare different DNS servers:

```bash
dig @1.1.1.1 example.com
```

and:

```bash
dig @8.8.8.8 example.com
```

If one works while another does not, the problem may involve the DNS resolver being used by the system.

---

# 18. hostname

`hostname` displays or modifies the system hostname.

Show hostname:

```bash
hostname
```

Show detailed hostname information:

```bash
hostnamectl
```

Example:

```bash
hostnamectl status
```

The hostname identifies the machine on a network.

---

# 19. traceroute

`traceroute` shows the network path packets take toward a destination.

Example:

```bash
traceroute example.com
```

It can help identify where network connectivity problems occur.

On some systems the command may need to be installed separately.

Example:

```bash
sudo pacman -S traceroute
```

or:

```bash
sudo apt install traceroute
```

---

# 20. SSH

SSH stands for Secure Shell.

It allows a user to remotely access another system securely.

Basic usage:

```bash
ssh username@hostname
```

Example:

```bash
ssh user@192.168.1.20
```

Specify a port:

```bash
ssh -p 2222 user@192.168.1.20
```

Use a private key:

```bash
ssh -i ~/.ssh/id_ed25519 user@server
```

SSH normally uses port:

```text
22
```

---

# 21. scp

`scp` copies files between systems using SSH.

Copy a local file to a remote machine:

```bash
scp file.txt user@server:/home/user/
```

Copy a remote file locally:

```bash
scp user@server:/home/user/file.txt .
```

Copy a directory:

```bash
scp -r directory user@server:/home/user/
```

---

# 22. Network Troubleshooting Workflow

When a service cannot be reached, troubleshoot from the lower layers upward.

A useful sequence is:

### Step 1: Check network interfaces

```bash
ip addr
```

Make sure the expected interface exists and has an IP address.

### Step 2: Check the routing table

```bash
ip route
```

Look for a default route.

### Step 3: Test IP connectivity

```bash
ping 8.8.8.8
```

If this fails, investigate basic network connectivity.

### Step 4: Test DNS

```bash
dig +short example.com
```

If IP connectivity works but DNS does not, investigate DNS.

### Step 5: Test the destination

```bash
ping example.com
```

Remember that ping can be blocked.

### Step 6: Test the service port

```bash
ss -ltn
```

for local services, or:

```bash
curl http://example.com
```

for HTTP services.

### Step 7: Inspect the service

If the port should belong to a local service, check whether the service is running.

For systemd services:

```bash
systemctl status service-name
```

---

# 23. Common Network Problems

## "Could not resolve host"

This usually indicates a DNS resolution problem.

Check:

```bash
dig example.com
```

and:

```bash
cat /etc/resolv.conf
```

---

## "Connection refused"

This generally means the destination was reachable but no service accepted the connection on that port.

Possible causes:

* Service is not running
* Wrong port
* Service is listening on another address
* Service configuration is incorrect

Check local listeners:

```bash
ss -ltnp
```

---

## "Connection timed out"

A timeout can indicate:

* Firewall blocking traffic
* Network routing problem
* Server unavailable
* Incorrect address
* Service not responding

Check:

```bash
ip route
```

and:

```bash
ping <ip>
```

where appropriate.

---

## "Network is unreachable"

This usually indicates that the system does not have a usable route to the destination.

Check:

```bash
ip route
```

---

## "No route to host"

This indicates that the system or network cannot find a usable route to the destination, or that the destination/network is rejecting the traffic in a way that results in this error.

Check:

```bash
ip route
```

---

# 24. HTTP Status Codes

When using `curl`, HTTP responses may include status codes.

Common codes:

```text
200 OK
201 Created
204 No Content
301 Moved Permanently
302 Found
400 Bad Request
401 Unauthorized
403 Forbidden
404 Not Found
500 Internal Server Error
502 Bad Gateway
503 Service Unavailable
```

Show response headers with:

```bash
curl -I https://example.com
```

---

# 25. Testing an HTTP API

A simple GET request:

```bash
curl http://localhost:8080/api/users
```

A POST request with JSON:

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"username":"alice"}' \
  http://localhost:8080/api/users
```

For debugging:

```bash
curl -v http://localhost:8080
```

The `-v` option displays detailed information about the connection and request.

---

# 26. Useful Network Commands Summary

| Command       | Purpose                             |
| ------------- | ----------------------------------- |
| `ip addr`     | Show IP addresses and interfaces    |
| `ip link`     | Show network interfaces             |
| `ip route`    | Show routing table                  |
| `ip neigh`    | Show network neighbors              |
| `ping`        | Test basic network reachability     |
| `curl`        | Make network/HTTP requests          |
| `wget`        | Download files                      |
| `ss`          | Inspect sockets and listening ports |
| `dig`         | Query DNS                           |
| `nslookup`    | Query DNS                           |
| `hostname`    | Show system hostname                |
| `hostnamectl` | Show/manage hostname information    |
| `traceroute`  | Show network path                   |
| `ssh`         | Secure remote shell                 |
| `scp`         | Copy files over SSH                 |

---

# 27. Quick Troubleshooting Reference

### Check your IP

```bash
ip addr
```

### Check your route

```bash
ip route
```

### Test internet connectivity by IP

```bash
ping 8.8.8.8
```

### Test DNS

```bash
dig +short example.com
```

### Test HTTP

```bash
curl -v https://example.com
```

### Check local listening ports

```bash
ss -ltnp
```

### Check a specific HTTP service

```bash
curl http://localhost:8080
```

### Check hostname

```bash
hostnamectl
```

### Connect to a remote server

```bash
ssh user@server
```

---

# 28. Important Distinctions

### IP address vs Port

IP address identifies the host/interface.

```text
192.168.1.10
```

Port identifies a service on that host.

```text
8080
```

Together:

```text
192.168.1.10:8080
```

### DNS vs Routing

DNS answers:

```text
"What IP address belongs to this domain?"
```

Routing answers:

```text
"Where should this packet go?"
```

### ping vs curl

`ping` tests network-level reachability using ICMP.

```bash
ping example.com
```

`curl` tests application-level communication, commonly HTTP.

```bash
curl https://example.com
```

A server can reject `ping` while still responding successfully to `curl`.

### TCP vs UDP

TCP provides reliable, ordered communication.

UDP provides a simpler connectionless transport mechanism without TCP's built-in reliability and ordering guarantees.