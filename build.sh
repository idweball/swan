#!/bin/bash

usage(){
  echo "Usage of: `basename $0` {build|clean|run}"
  exit 1
}

workspace=`cd $(dirname $0);pwd`
cd $workspace

case $1 in
  "build")
    bash scripts/build.sh
    ;;
  "clean")
    bash scripts/clean.sh
    ;;
  "run")
    bash scripts/run.sh
    ;;
  *)
  usage
  ;;
esac