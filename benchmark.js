var apiBenchmark = require('api-benchmark');
 
var services = {
  "sample": "http://localhost/v1/"
};
 
var routes = {
  "getNormal": "normal",
  "getGoroutine": "goroutine",
  "getChannel": "gochannels"
};
 
var options = { 
    "minSamples": 200000,
    "runMode": "parallel",
    "maxConcurrentRequests": 1000,
    "stopOnError": false
};
 
apiBenchmark.compare(services, routes, options, function(err, results){
    apiBenchmark.getHtml(results, function(error, html) {
    console.log(html);
  });
});