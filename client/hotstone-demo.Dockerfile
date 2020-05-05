FROM node:10-alpine

RUN apk add --update --no-cache git bash
WORKDIR /usr/src/client
COPY . .

RUN npm run build
RUN npm pack

WORKDIR /usr/src/client/examples/server-side-rendering
RUN npm install

CMD ["npx", "nodemon", "server"]
