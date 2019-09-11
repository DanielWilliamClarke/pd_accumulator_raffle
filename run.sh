#!/bin/bash

function sectionTitle() {
    echo "----------------------------------------------------"
    echo $1
    echo "----------------------------------------------------"
}

sectionTitle "Install npm dependancies"
pushd src/ui
npm install
npm run build:react
popd

sectionTitle "Set permissions"
chmod -R 777 src

sectionTitle "Build and run raffle"
pushd src
go run raffle.go
popd
