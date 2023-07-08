#!/bin/bash
go mod download
echo "Installing sqlite3 drivers; this usually takes a while..."
go install github.com/mattn/go-sqlite3

