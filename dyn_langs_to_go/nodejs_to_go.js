var ffi = require('ffi');

var lib = ffi.Library('go_adder/libadder', {
  'Add': [ 'int', [ 'int', 'int' ] ]
});

console.log('Javascript says: about to call Go ..');
var total = lib.Add(39, 3);
console.log('Javascript says: result is ' + total);
