# just a test dockerfile

FROM alpine:latest

WORKDIR /app
COPY README.md .

RUN echo "doing some runs" > test.md

RUN apk add bash

ENTRYPOINT ["bash"]
