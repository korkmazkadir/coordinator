FROM ubuntu:18.04

WORKDIR /application

# Adds node folder to path to access algorand binaries
ENV PATH="/application:${PATH}"

COPY service-app /application/

ENTRYPOINT ["service-app"]
