FROM centos:8

COPY jayce_server /workspace/jayce_server

COPY conf/server.json /workspace/conf/server.json

WORKDIR /workspace

EXPOSE 8090

ENTRYPOINT ["/workspace/jayce_server"]