#!/bin/bash

docker rmi -f ebook-reader

docker build -t ebook-reader .

echo "Image built. Run './scripts/docker-run.sh' to run the image."