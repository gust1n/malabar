var PROTO_PATH = __dirname + '/../locationService/locationservice/locationservice.proto';

var grpc = require('grpc');
var fs = require('fs');

var locationService_proto = grpc.load(PROTO_PATH);

// TODO: Implement mutual TLS, docs are scarce
// var path = require('path');
// function withConfDir(filename) {
//	return path.join(process.env.HOME, '.locationService', 'client', filename);
//}

function main() {
	var client = new locationService_proto.LocationService('localhost:9000',
		grpc.credentials.createInsecure());
	var call = client.trackUser({userID: 1});
	call.on('data', function(point) {
		console.log(point);
	});
}

main();
