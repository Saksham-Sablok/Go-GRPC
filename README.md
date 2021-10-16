1. create chat.proto file
2. ecport the variables 
    export GO_PATH=~/go
    export PATH=$PATH:/$GO_PATH/bin
3. compile the .proto file with following ccommand
     protoc --go_out=plugins=grpc:chat chat.proto
4. create aggreement.go file which contains the server defination and defination of the fuinction declared in chat.proto
5. we nwwd to create go mod init in the main directry of the project not for the individual packages
also go.mod file will not have the local packages inside it.
6. to enable logginfg at grpc level. 
$ export GRPC_GO_LOG_VERBOSITY_LEVEL=99
$ export GRPC_GO_LOG_SEVERITY_LEVEL=info
7.
