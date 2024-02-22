# Skopeo but an HTTP + JSON API

[skopeo](https://github.com/containers/skopeo) is an awesome CLI for working
with images. `skopeo-api` uses the same library that `skopeo` uses,
[containers/image](https://github.com/containers/image), but provides an HTTP
API instead of a CLI.

## Usage

```bash
go run ./main.go

# Second terminal
curl http://localhost:8080/tags/:image-repo
```

Example output:

```json
["latest", "0.0.3", "0.0.2", "0.0.1"]
```
