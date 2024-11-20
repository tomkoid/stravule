FROM node:20-alpine AS builder

WORKDIR /app

COPY package*.json .
RUN npm install --verbose

COPY .npmrc .
COPY postcss.config.js .
COPY svelte.config.js .
COPY tailwind.config.ts .
COPY tsconfig.json .
COPY vite.config.ts .
COPY .svelte-kit .
COPY .env .

COPY ./static /app/static
COPY ./src /app/src

RUN npm run build

CMD [ "node", "build" ]
