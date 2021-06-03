#build project
FROM golang:1.13-alpine as builder
ARG BUILDFLAGS="-mod=vendor"
ARG LDFLAGS=""
ARG PRODUCT=calc
ARG WORKINGDIR=calc

RUN apk add --update make

WORKDIR /${WORKINGDIR}

COPY . .

# build project
RUN make build -W tidy \
    SHELL=/bin/sh \
    GOOS=linux \
    GOENVS="CGO_ENABLED=0" \
    BUILDFLAGS="${BUILDFLAGS}" \
    LDFLAGS="-extldflags '-static' ${LDFLAGS}"

# run executable
FROM scratch
ARG PRODUCT=calc
ARG WORKINGDIR=calc

WORKDIR /

COPY --from=builder /${WORKINGDIR}/${PRODUCT} .

ENTRYPOINT [ "/calc" ]