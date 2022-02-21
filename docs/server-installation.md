# Server installation

## DNS configuration

One domain and one subdomain must point to the IP address of a [dedicated server](server-requirements.md).

| FQDN           | Record type | Address |
| -------------- | ----------- | ------- |
| ((DOMAIN))     | A           | X.X.X.X |
| git.((DOMAIN)) | A           | X.X.X.X |

## Network configuration

### Inbound

| Port | Transport | Application      |
| ---- | --------- | ---------------- |
| 80   | TCP, UDP  | HTTP/(1.1, 2, 3) |
| 443  | TCP, UDP  | HTTP/(1.1, 2, 3) |
| 521  | TCP       | SSH              |

### Outbound

| Port | Transport | Application |
| ---- | --------- | ----------- |
| 587  | TCP       | SMTP        |

### OS installation

Download and boot the ISO image `amd64` of [Debian](https://www.debian.org/distrib/netinst)

Select :

- "Advanced options ..."
- "... Automated install"

The network is automatically configured with DHCP, you can also configure it manually.

At the prompt "Location of initial preconfiguration file:", enter the URL :

```
raw.githubusercontent.com/01-edu/public/master/sh/debian/preseed.cfg
```

and select "Continue".

--> Please let us know when the server's remote access is ready and we will configure it.
