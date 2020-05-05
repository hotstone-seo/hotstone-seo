FROM node:10-alpine AS client-builder

RUN apk add --update --no-cache git bash
WORKDIR /usr/src/client
COPY . .

RUN npm install
RUN npm run build
RUN npm pack
RUN ls -hal ./

# === BUILD FINAL ===
FROM node:10-alpine

RUN apk update && apk add --no-cache git bash 

COPY --from=client-builder /usr/src/client/examples/server-side-rendering /app/demo
COPY --from=client-builder /usr/src/client/*.tgz /app/demo/vendor/

WORKDIR /app/demo
RUN ls -hal ./
RUN ls -hal ./vendor

RUN npm install --no-package-lock
RUN npx webpack

CMD ["npm", "run", "start:server"]
