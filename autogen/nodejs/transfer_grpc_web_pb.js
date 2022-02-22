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
proto.rusk = require('./transfer_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.rusk.TransferClient =
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
proto.rusk.TransferPromiseClient =
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
 *   !proto.rusk.TransferTransactionRequest,
 *   !proto.rusk.Transaction>}
 */
const methodDescriptor_Transfer_NewTransfer = new grpc.web.MethodDescriptor(
  '/rusk.Transfer/NewTransfer',
  grpc.web.MethodType.UNARY,
  proto.rusk.TransferTransactionRequest,
  transaction_pb.Transaction,
  /**
   * @param {!proto.rusk.TransferTransactionRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  transaction_pb.Transaction.deserializeBinary
);


/**
 * @param {!proto.rusk.TransferTransactionRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.Transaction)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.Transaction>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.TransferClient.prototype.newTransfer =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.Transfer/NewTransfer',
      request,
      metadata || {},
      methodDescriptor_Transfer_NewTransfer,
      callback);
};


/**
 * @param {!proto.rusk.TransferTransactionRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.Transaction>}
 *     Promise that resolves to the response
 */
proto.rusk.TransferPromiseClient.prototype.newTransfer =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.Transfer/NewTransfer',
      request,
      metadata || {},
      methodDescriptor_Transfer_NewTransfer);
};


module.exports = proto.rusk;

