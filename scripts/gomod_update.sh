#!/bin/bash

rm go.mod
rm go.sum

go mod init leetcode_history
go mod tidy

# git add go.mod go.sum
