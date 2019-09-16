#!/usr/bin/env bash

[[ $# = 0 || $# > 1 || $1 =~ ^[+-]?[0-9]*\.[0-9]+$ || $1 =~ ^[a-zA-Z]*$ ]] && echo "Usage: leap.sh <year>" && exit 1

if [[ $(($1 % 4)) -eq 0 ]]; then
  if [[ $(($1 % 100)) -eq 0 ]]; then
    if [[ $(($1 % 400)) -eq 0 ]]; then
      echo true
    else
      echo false
    fi
  else
    echo true
  fi
else
  echo false
fi

