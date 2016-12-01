FROM ubuntu:16.04

WORKDIR /app/

COPY ./build/dp-dd-frontend-controller .

ENTRYPOINT ./dp-dd-frontend-controller
