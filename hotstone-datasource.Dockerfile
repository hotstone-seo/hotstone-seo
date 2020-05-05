# === build backend ===
FROM golang:alpine AS backend-builder

ENV ROOT_REPO .
RUN apk add --update --no-cache git bash npm
WORKDIR /usr/src/backend
COPY ${ROOT_REPO}/ .

RUN go get -u -v golang.org/x/tools/cmd/goimports

CMD ["/bin/bash", "/usr/src/backend/typicalw", "json-server"]
