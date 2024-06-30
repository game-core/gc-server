FROM node:22.0.0-alpine AS builder

WORKDIR /app
RUN apk upgrade --update && apk add bash && apk --no-cache add git
COPY view/admin /app
RUN yarn install
RUN yarn build

# Prod
FROM node:22.0.0-alpine AS prod

WORKDIR /app
COPY --from=builder /app/package.json .
COPY --from=builder /app/yarn.lock .
COPY --from=builder /app/.output ./.output
RUN yarn install

EXPOSE 3000

CMD ["yarn", "start"]

# Dev
FROM node:22.0.0-alpine AS dev

WORKDIR /app
RUN apk upgrade --update && apk add bash && apk --no-cache add git
COPY view/admin /app
RUN yarn install

EXPOSE 3000

CMD ["yarn", "dev"]
