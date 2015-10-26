var ffi = require('ffi');

var lib = ffi.Library('go_adder/libadder', {
  'Add': [ 'long', [ 'long', 'long' ] ]
});

console.log('JavaScript says: about to call Go ..');
var total = lib.Add(39, 3);
console.log('JavaScript says: result is ' + total);
