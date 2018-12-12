## To setup - Just edit PROJECT_PATH VARIABLE PROPERLY AND RUN SCRIPT AS sh start_restserver.sh





## This file does things listed below :
# 1) - It puts the system variables in place
# 2) - It builds the code, and creates a server binary in current directory
# 3) - It executes the server binary

PROJECT_PATH="/Users/nsachdeva/notelibrary/notelibrary";

cd "$PROJECT_PATH/src/app"

echo "Setting GOPATH";
export GOPATH=$PROJECT_PATH;
echo "GOPATH Set to - ", $GOPATH;

echo "Setting vendor experiment variable";
export GO15VENDOREXPERIMENT=1;
echo "GO15VENDOREXPERIMENT Set to - ", $GO15VENDOREXPERIMENT;

echo "Building Server";
go build restserver.go;
echo "Restserver Built"

echo "Starting Server - find out  sh vs ./";
./restserver

#### ENDS HERE ####
