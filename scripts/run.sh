#!/bin/bash

read -p "Port: " PORT

cd ./app && env PORT="$PORT" ./main