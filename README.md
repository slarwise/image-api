# Skopeo but an HTTP + JSON API

[skopeo](https://github.com/containers/skopeo) is an awesome CLI for working
with images. `skopeo-api` uses the same library that `skopeo` uses,
[containers/image](https://github.com/containers/image), but provides an HTTP
API instead of a CLI.

Listing tags is the only supported functionality at the moment.

## Usage

```bash
# Start the server on port 8080
go run ./main.go

# List the tags for an image repository
# Escape the slashes
curl http://localhost:8080/tags/ghcr.io%2Fslarwise%2Fskopeo-api
curl http://localhost:8080/tags/docker.elastic.co%2Felasticsearch%2Felasticsearch
curl http://localhost:8080/tags/rancher%2Ffleet
```

Example output:

```json
["latest", "0.0.3", "0.0.2", "0.0.1"]
```
