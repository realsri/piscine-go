#!/bin/bash
curl -s https://assets.01-edu.org/superhero/all.json | jq '.[] | select(.id == 70) | .name'
