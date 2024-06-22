FROM node:22.0.0-alpine AS builder

WORKDIR /view
RUN apk upgrade --update && apk add bash && apk --no-cache add git

COPY .. .

RUN yarn install
RUN yarn build

FROM node:22.0.0-alpine AS dev

WORKDIR /view
RUN apk upgrade --update && apk add bash && apk --no-cache add git

COPY .. .

RUN yarn install

EXPOSE 3000

CMD ["yarn", "dev"]

FROM node:22.0.0-alpine AS prod

WORKDIR /view
COPY --from=builder /view/package.json ./
COPY --from=builder /view/yarn.lock ./
COPY --from=builder /view/.output ./.output
RUN yarn install

EXPOSE 3000

CMD ["yarn", "start"]
