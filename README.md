# Simple React Server

Simple HTTP server as a replacement for nginx docker image for serving static react frontend applications.

## Usage

Add your builded artefacts at `/var/www`. If for some reason that path is not ok for you you can override it.

## Overrides

| Attribute | Default Value | Environment Variable | Run Parameter |
|---|---|---|---|
| Bind Address | `:3000` | `BIND_ADDRESS` | `--bind-address` |
| Serve Directory | `/var/www` | `SERVE_DIRECTORY` | `--serve-directory` |

### Example Dockerfile for your React application

```Dockerfile
FROM registry.dmz.croz.net/croz/simple-react-server

COPY build /var/www
```