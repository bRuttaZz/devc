#!/bin/sh

# MEANT TO BE RUN FROM ROOT DIR
# eg:
#   sh scripts/build-all-arch.sh
# The script is for compiling devc using the makefile for all the major 
# architecture of linux kernal (By now, due to the 
# less complicated nature it can be cross compiled atleast from amd64)
set -e

archs=('amd64' '386' 'arm' 'arm64')

do_compile() {
    for arch in ${archs[@]}; do
        echo -e "\n[BUILD] Compiling for $arch"
        make ARCH=$arch;
    done;
}

do_compile