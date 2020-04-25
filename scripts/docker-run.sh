#!/bin/bash

read -p "Port (will be bound to 9000 inside the container): " PORT

docker run -p "$PORT":9000 ebook-reader