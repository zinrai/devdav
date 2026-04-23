# devdav

WebDAV server for development environments. Serves a local directory over HTTP
so that it can be mounted from another machine.

Uses [`golang.org/x/net/webdav`](https://pkg.go.dev/golang.org/x/net/webdav) for
protocol handling. No authentication, no TLS, no access control. Intended for
trusted local networks.

## Usage

```
devdav [-addr host:port] [-dir path] [-verbose]
```

Flags:

- `-addr`: listen address. Defaults to `127.0.0.1:8080`.
- `-dir`: directory to serve. Defaults to the current directory.
- `-verbose`: log file access (HTTP method and path) for each request.

Example:

```
devdav -addr 0.0.0.0:8080 -dir /home/zinrai/share -verbose
```

## Mounting from a client

Using [`rclone`](https://rclone.org/):

```
rclone config create devdav webdav url=http://<host>:8080 vendor=other
mkdir -p ~/devdav
rclone mount devdav: ~/devdav --daemon
```

To unmount:

```
fusermount -u ~/devdav
```

## Non-goals

- Production deployment
- Multi-user access control
- Public internet exposure
- Feature parity with Apache `mod_dav` or nginx `dav_module`

## License

This project is licensed under the [MIT License](./LICENSE).
