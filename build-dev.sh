#!/bin/bash
go build
shopt -s extglob
strip tproxy
echo "Image: urtho/tproxy"
docker build . -t urtho/tproxy:latest 
docker push -a urtho/tproxy
