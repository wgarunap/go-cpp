#!/usr/bin/env bash
cp -pvr swig/cpp.swig cpp/
cd cpp/

if [[ -z $( which swig ) ]]; then
    echo "Currently swig is NOT installed\nPlease install the latest version of swig here -> http://www.swig.org/download.html"
    exit 1
fi

swig -go -cgo -c++ -intgosize 64 cpp.swig
rm -rf cpp.swig
cd ..

echo "Building your project..."
go build -o cpp-swig2
echo "Running your test project..."
./cpp-swig2
rm -rf cpp-swig2

