FROM node:22.0.0-alpine AS builder

WORKDIR /web
RUN apk upgrade --update && apk add bash && apk --no-cache add git

COPY .. .

RUN yarn install
RUN yarn build

FROM node:22.0.0-alpine AS dev

WORKDIR /web
RUN apk upgrade --update && apk add bash && apk --no-cache add git

COPY .. .

RUN yarn install
RUN yarn build

EXPOSE 3000

CMD ["yarn", "dev"]

FROM node:22.0.0-alpine AS prod

WORKDIR /web
COPY --from=builder /web/package.json ./
COPY --from=builder /web/yarn.lock ./
COPY --from=builder /web/.output ./.output
RUN yarn install

EXPOSE 3000

CMD ["yarn", "start"]
