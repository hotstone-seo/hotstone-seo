# === build backend ===
FROM golang:alpine AS backend-builder

RUN apk update && apk add --no-cache git bash
WORKDIR /usr/src/backend
COPY ${ROOT_REPO}/ .

RUN go get -u -v golang.org/x/tools/cmd/goimports

CMD ["/bin/bash", "/usr/src/backend/typicalw", "json-server"]
