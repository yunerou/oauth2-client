FROM 886715051234.dkr.ecr.ap-northeast-1.amazonaws.com/payroll_api:base-cache119 as builder
# FROM base as builder
COPY . /app
RUN go get -d
# ARG opts
RUN CGO_ENABLED=0  go build -o /go/bin/myapp
# RUN env ${opts} go build -o /go/bin/myapp

# nonroot User have permission in /opt directory
FROM gcr.io/distroless/static-debian11:debug AS distroless-deployer
LABEL maintainer="yunerou"
USER nonroot:nonroot
COPY --from=builder --chown=nonroot:nonroot /go/bin/myapp /
COPY --from=builder --chown=nonroot:nonroot --chmod=755 /opt /opt
# COPY config.* /
COPY resources /resources
ENTRYPOINT ["/myapp"]
EXPOSE 8080