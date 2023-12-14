FROM golang:alpine
WORKDIR /src
COPY . .
RUN go build -o /bin/controller

FROM scratch
COPY --from=0 /bin/controller /bin/controller
CMD [ "/bin/controller" ]