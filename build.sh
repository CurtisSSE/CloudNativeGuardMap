#!/bin/bash

cd Site
npm install
npm run build
cd ..

go build -o main .
