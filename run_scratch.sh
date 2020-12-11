#!/bin/bash
name=begin
docker image build -t $name /Users/ben/src/github.com/BenLirio/begin
docker container run -p 8080:8080 -it $name
