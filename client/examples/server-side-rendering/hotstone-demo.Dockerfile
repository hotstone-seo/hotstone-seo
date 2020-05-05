FROM node:10-alpine

RUN apk add --update --no-cache git bash
WORKDIR /usr/src/app
COPY . .

RUN npm install

CMD ["npx", "nodemon", "server"]
