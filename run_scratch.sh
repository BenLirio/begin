#!/bin/bash
name=begin
docker image build -t $name /Users/ben/src/github.com/BenLirio/begin
docker container run -it $name
