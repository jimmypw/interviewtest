#!/bin/bash

go build killer.go
vagrant destroy -f
vagrant up
clear
vagrant ssh
