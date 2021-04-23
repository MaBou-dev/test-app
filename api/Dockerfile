# context is folder api (folder of the api part / golang source code)
FROM golang:1.14.3-alpine AS base
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download
COPY . ./

FROM base AS build
RUN go build -o /out/api .

# copy binary build to a scratch image to reduce layer and image size
FROM scratch AS bin
COPY --from=build /out/api /
ENTRYPOINT ["./api"]
