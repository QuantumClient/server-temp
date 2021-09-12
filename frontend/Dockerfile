FROM node:16-alpine3.14

WORKDIR /app/frontend
COPY package*.json .

# update and install dependency
RUN apk update && apk upgrade

# Install dependecies
RUN npm install

ENV HOST 0.0.0.0
# Build nuxt server

COPY . .

RUN npm run build


CMD [ "npm", "run", "start" ]
