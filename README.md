# API for container image info

Like [skopeo](https://github.com/containers/skopeo), but an an HTTP API instead
of a CLI. Listing tags is the only supported functionality at the moment.

## Run it locally

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

## Run on kubernetes

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: skopeo-api
  labels:
    app: skopeo-api
spec:
  selector:
    matchLabels:
      app: skopeo-api
  template:
    metadata:
      labels:
        app: skopeo-api
    spec:
      containers:
        - name: skopeo-api
          image: ghcr.io/slarwise/skopeo-api:0.0.1
          ports:
            - containerPort: 8080
              name: http
```
