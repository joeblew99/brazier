# Mobile


##  notes

This is a good candidate because it has BoltDB
But also Io via JSONRPC, GRPC, and CLI. 


Use the -work flag to gomobile to output working directory and to preserve the generated files
Use the -v flag to see all the output 

Hello is used a base test
in the hello folder:
- gomobile bind --target=android -v
- gomobile bind --target=ios -v

in the cli folder:
- gomobile bind --target=android -v
- gomobile bind --target=ios -v

to  inspect the java output
- jar tf classes.jar
