/*
* This go file is for go build command to identify this as a go package,
* then go compiler will fetch all the files and will execute the swig file accordingly
*/

//cp -pvr ../swig/cpp.swig . && swig -go -cgo -c++ -intgosize 64 cpp.swig

package cpp
