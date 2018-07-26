#!/bin/bash
mkdir -p lib
echo 'build domainfinder'
go build -o domainfinder
echo 'build synonyms'
cd ../synonyms
go build -o ../domainfinder/lib/synonyms
echo 'build avaliable'
cd ../avaliable
go build -o ../domainfinder/lib/avaliable
echo 'build sprinkle'
cd ../sprinkle
go build -o ../domainfinder/lib/sprinkle
echo 'build coolify'
cd ../coolify
go build -o ../domainfinder/lib/coolify
echo 'build domainify'
cd ../domainify
go build -o ../domainfinder/lib/domainify
echo 'finished'