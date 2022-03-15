/**
 * @fileoverview gRPC-Web generated client stub for rusk
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var transaction_pb = require('./transaction_pb.js')
const proto = {};
proto.rusk = require('./stake_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.rusk.StakeServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.rusk.StakeServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.StakeTransactionRequest,
 *   !proto.rusk.Transaction>}
 */
const methodDescriptor_StakeService_NewStake = new grpc.web.MethodDescriptor(
  '/rusk.StakeService/NewStake',
  grpc.web.MethodType.UNARY,
  proto.rusk.StakeTransactionRequest,
  transaction_pb.Transaction,
  /**
   * @param {!proto.rusk.StakeTransactionRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  transaction_pb.Transaction.deserializeBinary
);


/**
 * @param {!proto.rusk.StakeTransactionRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.Transaction)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.Transaction>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.StakeServiceClient.prototype.newStake =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.StakeService/NewStake',
      request,
      metadata || {},
      methodDescriptor_StakeService_NewStake,
      callback);
};


/**
 * @param {!proto.rusk.StakeTransactionRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.Transaction>}
 *     Promise that resolves to the response
 */
proto.rusk.StakeServicePromiseClient.prototype.newStake =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.StakeService/NewStake',
      request,
      metadata || {},
      methodDescriptor_StakeService_NewStake);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.FindStakeRequest,
 *   !proto.rusk.FindStakeResponse>}
 */
const methodDescriptor_StakeService_FindStake = new grpc.web.MethodDescriptor(
  '/rusk.StakeService/FindStake',
  grpc.web.MethodType.UNARY,
  proto.rusk.FindStakeRequest,
  proto.rusk.FindStakeResponse,
  /**
   * @param {!proto.rusk.FindStakeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.FindStakeResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.FindStakeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.FindStakeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.FindStakeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.StakeServiceClient.prototype.findStake =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.StakeService/FindStake',
      request,
      metadata || {},
      methodDescriptor_StakeService_FindStake,
      callback);
};


/**
 * @param {!proto.rusk.FindStakeRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.FindStakeResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.StakeServicePromiseClient.prototype.findStake =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.StakeService/FindStake',
      request,
      metadata || {},
      methodDescriptor_StakeService_FindStake);
};


module.exports = proto.rusk;

