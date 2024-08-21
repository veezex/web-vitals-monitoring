# Consists of the following parts:

## A golang server that allows recording new metrics into sqlite

All configurations are located in the .env file.

Port from .env should be open for the server to work.

#### Script example

```html
  <script type="module" async src="http://localhost:6510/script"></script>
```

#### Install the dependencies

```bash
apt install sqlite3 certbot go git make -y
```

### Create service 
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

### Create certificate (if use https)

```bash
make certs
```

### Change the permissions of the certificate files and copy the to the project dir
```bash
chown admin:admin /etc/letsencrypt/live/wwsg1.aktee.top/fullchain.pem
chown admin:admin /etc/letsencrypt/live/wwsg1.aktee.top/privkey.pem
```