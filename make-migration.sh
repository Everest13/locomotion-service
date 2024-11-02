#!/bin/sh

#bin/goose -dir migrations create $1 sql

cd migrations
goose create locomotion sql

cd ../