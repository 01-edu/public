# Server installation

## 🌐 DNS Configuration

The following DNS records should be configured in your domain's [zone file](https://en.wikipedia.org/wiki/Zone_file) or through the web interface of your dns provider/domain registrar.  
- A _domain/subdomain_ pointing to the public IP address (using `A Record`) of your [dedicated server](server-requirements.md).
- A _subdomain_ called `git` pointing to the above mentioned *domain/subdomain* (using `CNAME Record`) or it's IP address (using `A Record`).

Your newly configured DNS records should look like this:

| FQDN           | Record type | Target        |
| -------------- | ----------- | ------------- |
| ((DOMAIN))     | A           | X.X.X.X       |
| git.((DOMAIN)) | CNAME       | ((DOMAIN))    |

Here is an _example_ of the DNS records for the domain `example.org` with the public IP address of `93.184.216.34`:

| FQDN            | Record type | Target        |
| --------------  | ----------- | ------------- |
| example.org     | A           | 93.184.216.34 |
| git.example.org | CNAME       | example.org   |


## 🛠️ Network Configuration

### ➡️ Inbound

| Port        | Protocol(s) | Service/Application |
| ----------- | ----------- | ------------------- |
| 80          | TCP         | HTTP/(1.1, 2, 3)    |
| 443         | TCP         | HTTP(S)/(1.1, 2, 3) |
| 521         | TCP         | SSH                 |
| 8080 - 8090 | TCP         | HTTP/(1.1, 2, 3)    |       

### ⬅️ Outbound

| Port        | Protocol(s) | Service/Application  |
| ----------- | ----------- | -------------------- |
| 587         | TCP         | SMTP                 |
| 8080 - 8090 | TCP         | HTTP/(1.1, 2, 3)     | 

## 💿 OS installation

1. Download and boot the `amd64` variant of the [Debian](https://www.debian.org/distrib/netinst) ISO image.

2. Select :
  - "**Advanced options ...**"
  - "**... Automated install**"

3. The network is automatically configured using your DHCP server. Additionally, you can also configure it manually to suit your preference.

4. At the prompt "Location of initial preconfiguration file:", please enter the following URL :

  ```console
  raw.githubusercontent.com/01-edu/public/master/sh/debian/preseed.cfg
  ```

5. Then select "**Continue**" and follow the on-screen instructions.


## 🏁 Finishing up
Once the server is ready to be accessed remotely, please let us know via approriate communication channels and we will proceed with configuring the server.
