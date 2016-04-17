#!/usr/bin/env node

var lineReader = require('readline').createInterface({
  input: require('fs').createReadStream('out.csv')
});

var arr = []

lineReader.on('line', function (line) {
  var res = +line.substr(line.indexOf(':') + 1);
  if(arr.length > 0 ) {
    var diff = arr[arr.length - 1] - res;
    console.log(diff);
  }
  arr.push(res);
})

