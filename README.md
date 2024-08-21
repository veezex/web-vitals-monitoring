## A server that allows recording new metrics into SQLite

All configurations are located in the .env file.
Port from .env should be open for the server to work. (for Certbot to work, port 80 should be open)

### Install the dependencies (Debian)
```bash
apt install sqlite3 git make build-essential
```
Additionally, you need to install [Go 1.21 or later](https://go.dev/doc/install).

```bash
git clone https://github.com/veezex/web-vitals-monitoring.git
```
Next, you will need to create a .env (see .env.example) file in the project directory.
After that, you can run the server.

```bash
make build
./serverapp
```

## Create service (optionally)
```bash
sudo vi /etc/systemd/system/ww.service
```
```
[Unit]
Description=Go Server (WW)
After=network.target
[Service]
User=admin
Group=admin
ExecStart=/home/admin/web-vitals-monitoring/serverapp
Restart=always
[Install]
WantedBy=multi-user.target
```

## HTTPS support

### Create certificate
```bash
apt install certbot
make certs
```

### Change the permissions of the certificate files and copy them to the project directory
```bash
chown admin:admin /etc/letsencrypt/live/domain.com/fullchain.pem
chown admin:admin /etc/letsencrypt/live/domain.com/privkey.pem
```
