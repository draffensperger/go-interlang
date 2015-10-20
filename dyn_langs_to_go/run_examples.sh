echo 'Calling Go library from interpreted languages ...'
echo

echo '---- Ruby ----'
ruby ruby_to_go.rb
echo

echo '---- Python ----'
python python_to_go.py
echo

echo '---- Node.js ----'
# Use the global path, so we can do "npm install ffi -g" and avoid creating a
# local node_modules folder.
NODE_PATH=`npm config get prefix`/lib/node_modules node nodejs_to_go.js
echo

echo '---- Java ----'
java -cp java_to_go/target/classes:java_to_go/target/dependency/* javatogo.JavaToGo
