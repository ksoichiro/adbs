#!/bin/bash

PGM=adbs
TAG=
if [ $# -ge 1 ]; then
  TAG=$1
else
  TAG=`git rev-parse --short HEAD`
fi

function build() {
  echo "[1;35mBuilding for $1/$2...[m"
  local arch_dir=$PGM-$TAG-bin-$1-$2
  local bin_dir=build/${arch_dir}
  mkdir -p ${bin_dir}
  local ext=
  if [ $1 == "windows" ]; then
    ext=.exe
  fi
  GOOS=$1 GOARCH=$2 go build -o ${bin_dir}/${PGM}${ext}
  local status=$?
  popd > /dev/null 2>&1
  if [ $status -eq 0 ]; then
    for i in $(find ${bin_dir}/ | grep ".DS_Store"); do
      rm -f ${i}
    done
    pushd build > /dev/null 2>&1
    tar czf ${arch_dir}.tar.gz ${arch_dir}
    zip -ry ${arch_dir}.zip ${arch_dir} > /dev/null
    popd > /dev/null 2>&1
  fi
}

if [ -d build ]; then
  rm -rf build/*
fi

build darwin 386
build darwin amd64
build linux 386
build linux amd64
build linux arm
build windows 386
build windows amd64
