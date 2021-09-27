FROM node:16-alpine3.14

WORKDIR /app/mc
COPY package*.json ./

# update and install dependency
RUN apk update && apk upgrade

# Install dependecies
RUN npm install

COPY . .

RUN npm run build


CMD [ "npm", "run", "start" ]
