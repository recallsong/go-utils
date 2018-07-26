#!/bin/bash
GOPATH=$(go env GOPATH)
GOPATH_DIR=${GOPATH%%:*}
cd "${GOPATH_DIR}/src/"
echo "GOPATH is ${GOPATH_DIR}"

mkdir -p golang.org/x && cd golang.org/x

if [ ! -n "$1" ] ;then
    dirs=("net" "sys" "tools" "text" "crypto" "time")
    for var in ${dirs[@]};  
    do  
        echo $var  
        git clone --depth=1 https://github.com/golang/${var}.git ${var}
        rm -rf ${var}/.git
    done
else
    var=$1
    echo $var  
    git clone --depth=1 https://github.com/golang/${var}.git ${var}
    rm -rf ${var}/.git
    
fi

