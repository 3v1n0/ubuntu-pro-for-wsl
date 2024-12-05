(ref::landscape-config)=
# Landscape configuration schema

Both Landscape clients are configured via a single, plain text configuration file (e.g., `landscape.conf` or `landscape.ini`). This file is provided to the Windows host.

The schema for this file is the same as Landscape for Ubuntu desktop or server, with a few additional keys specific to the WSL settings, which can be grouped into keys that affect just the Windows-side client and keys that affect both the Windows-side client and the Ubuntu WSL-side client(s). These additions are documented below.

> See more: [Landscape | Configure Ubuntu Pro for WSL for Landscape](https://ubuntu.com/landscape/docs/register-wsl-hosts-to-landscape/)

Here is an example of what the configuration looks like:

```ini
[host]
url = landscape-server.domain.com:6554

[client]
url = https://landscape-server.domain.com/message-system
ping_url  = http://landscape-server.domain.com/ping
account_name = standalone
log_level = debug
ssl_public_key = C:\Users\user\Downloads\landscape_server.pem
```

## Host

This section contains settings unique to the Windows-side client. Currently these consist of a single key:
- `url`: The URL of your Landscape account followed by a colon (`:`) and the port number. Port 6554 is the default for Landscape Quickstart installations.

## Client

This section contains settings used by both clients. Most keys in this section behave the same way they would on a traditional Landscape setup. Only the following keys behave differently:
- `ssl_public_key`: This key must be a Windows path. The WSL instances will have this path translated automatically.
- `computer_title`: This key will be ignored. Instead, each WSL instance will use its Distro name as computer title.
- `hostagent_uid`: This key will be ignored.

> See more: [GitHub | Landscape client configuration schema](https://github.com/canonical/landscape-client/blob/master/example.conf)