FROM golang:alpine AS build-env
WORKDIR /app
ADD main.go /app
RUN cd /app && \
    go build -o owner-validator

FROM alpine
WORKDIR /app
COPY --from=build-env /app/owner-validator /app
EXPOSE 443
CMD [ "./owner-validator" ]