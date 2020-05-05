# === build frontend ===
FROM node:12 AS frontend-builder
ENV ROOT_REPO .
RUN mkdir /usr/src/frontend
WORKDIR /usr/src/frontend
ENV PATH /usr/src/frontend/node_modules/.bin:$PATH
COPY ${ROOT_REPO}/ui /usr/src/frontend
RUN npm install
RUN npm run build

# === build backend ===
FROM golang:alpine AS backend-builder

RUN apk update && apk add --no-cache git bash
WORKDIR /usr/src/backend
COPY ${ROOT_REPO}/ .

RUN go get -u -v golang.org/x/tools/cmd/goimports

RUN go build -o bin/hotstone-seo  cmd/hotstone-seo/main.go

# === BUILD FINAL ===
FROM golang:alpine

RUN apk update && apk add --no-cache git bash 

COPY --from=frontend-builder /usr/src/frontend/build /app/build
COPY --from=backend-builder /usr/src/backend/bin/hotstone-seo /app/hotstone-seo

COPY --from=backend-builder /usr/src/backend /src/backend

WORKDIR /app
CMD ["./hotstone-seo"]
