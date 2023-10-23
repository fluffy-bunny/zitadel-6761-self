#! /bin/sh

set -eux

cd $GOPATH/src/github.com/zitadel/zitadel/tools
for imp in `cat tools.go | grep "_" | sed -E "s/_ \"(.*.+)\"/\1/g"`; do
    echo "installing $imp"
    go install $imp
done
cd -