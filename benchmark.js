var apiBenchmark = require('api-benchmark');
 
var services = {
  "sample": "http://localhost/v1/"
};
 
var routes = {
  "getNormal": "normal",
  "getGoroutine": "goroutine"
};
 
var options = { 
    "minSamples": 20000,
    "runMode": "parallel",
    "maxConcurrentRequests": 100,
    "stopOnError": false
};
 
apiBenchmark.compare(services, routes, options, function(err, results){
    apiBenchmark.getHtml(results, function(error, html) {
    console.log(html);
  });
});