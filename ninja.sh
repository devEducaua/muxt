#!/bin/sh

out="build/"
mkdir -p "$out"

rm -f "build.ninja"

append() {
    echo "$1" >> "build.ninja"
}

append "
target = muxt
builddir = build

rule mk
    command = mkdir -p \$out
    description = MK \$out

rule go
    command = go build -o \$out cmd/main.go
    description = GO \$out

build \$builddir: mk
"

files=$(fd -e go | tr '\n' ' ')

append "build \$builddir/\$target: go $files"

ninja $*
