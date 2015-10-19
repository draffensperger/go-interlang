
cd adder
http://stackoverflow.com/questions/15879993/writing-a-ruby-extension-in-go-golang


cd adder

go build -buildmode=c-shared -o libadder.so

Uses Unix domain sockets for communication
https://github.com/DavidHuie/quartz

http://stackoverflow.com/questions/15879993/writing-a-ruby-extension-in-go-golang

Ruby

gem install ffi

Javascript

npm install ffi -g
export NODE_PATH=/opt/lib/node_modules
NODE_PATH=/usr/local/lib/node_modules node nodejs_to_go.js

or npm install -g ffi
cp go_adder/libadder.so go_adder/libadder.dylib

export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:./go_adder


cd java_to_go
javac JavaToGo.java
javah JavaToGo.java

javac *.java
LD_LIBRARY_PATH=/Users/dave/Dropbox/Programming/golang/src/github.com/draffensperger/go-inter-lang-calls/go_from_dyn_langs/go_adder java JavaToGo

java -Djava.library.path="/Users/dave/Dropbox/Programming/golang/src/github.com/draffensperger/go-inter-lang-calls/go_from_dyn_langs/go_adder" JavaToGo 

gcc java_to_go.c

gcc -shared -o java_to_go.so java_to_go.c

http://stackoverflow.com/questions/5963266/call-c-function-from-java

http://stackoverflow.com/questions/20771803/where-can-i-find-the-jni-headers-for-mac-os


gcc -shared -o java_to_go.dylib -I/System/Library/Frameworks/JavaVM.framework/Headers java_to_go.c -L../go_adder -ladder


LD_LIBRARY_PATH=../go_adder/ java JavaToGo

LD_LIBRARY_PATH=/Users/dave/Dropbox/Programming/golang/src/github.com/draffensperger/go-inter-lang-calls/go_from_dyn_langs/go_adder java JavaToGo

http://stackoverflow.com/questions/5963266/call-c-function-from-java
http://mrjoelkemp.com/2012/01/getting-started-with-jni-and-c-on-osx-lion/
http://stackoverflow.com/questions/14529720/how-to-make-jni-h-be-found

http://stackoverflow.com/questions/6041934/can-shared-library-call-another-shared-library

http://stackoverflow.com/questions/3950635/how-to-compile-dynamic-library-for-a-jni-application-on-linux

http://stackoverflow.com/questions/13443364/how-can-i-use-linux-shared-libraries-in-java

http://stackoverflow.com/questions/1556421/use-jni-instead-of-jna-to-call-native-code

https://github.com/java-native-access/jna

http://stackoverflow.com/questions/8826881/maven-install-on-mac-os-x
brew install maven

java -cp target/classes:target/dependency/* javatogo.JavaToGo

mvn install && java -cp target/classes:target/dependency/* javatogo.JavaToGo
