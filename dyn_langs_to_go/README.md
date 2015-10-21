# Calling Go from interpreted languages

This folder contains code for calling a Go shared library from Python, Ruby,
Javascript (Node.js) and Java.

## Installing dependencies

- For Python, it probably already has [ctypes](https://docs.python.org/2/library/ctypes.html) installed
- For Ruby, `gem install ffi`
- For Javascript, `npm install ffi`
- For Java, install [Maven](https://maven.apache.org/), 
  then `mvn -f java_to_go/pom.xml install` (or `cd java_to_go` then `mvn install` and `cd ..`). 
  That will install [JNA](https://github.com/java-native-access/jna) and compile
  the Java source.

To install all the dependencies use `./install_ffi_deps./sh`

## Running the examples

To run the examples, first run `make` to build the Go shared library, then run
`./run_examples.sh` to run all of the examples.

## Other similar examples

http://stackoverflow.com/questions/15879993/writing-a-ruby-extension-in-go-golang
