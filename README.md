# Consists of the following parts:

## A golang server that allows recording new metrics into sqlite

All configurations are located in the .env file.

Port from .env should be open for the server to work. (for certbot to work port 80 should be open)

#### Install the dependencies

```bash
apt install sqlite3 certbot git make -y
```
Additionally, you need to install the go compiler.

### Create service (optionally)
```bash
sudo vi /etc/systemd/system/ww.service
```

```bash
[Unit]
Description=Go Server (WW)
After=network.target

[Service]
User=admin
Group=admin
ExecStart=/home/admin/executable
Restart=always

[Install]
WantedBy=multi-user.target
```

## HTTPS support

### Create certificate

```bash
make certs
```

### Change the permissions of the certificate files and copy the to the project dir
```bash
chown admin:admin /etc/letsencrypt/live/domain.com/fullchain.pem
chown admin:admin /etc/letsencrypt/live/domain.com/privkey.pem
```