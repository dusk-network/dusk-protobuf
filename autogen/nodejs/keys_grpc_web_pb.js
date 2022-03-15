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

const proto = {};
proto.rusk = require('./keys_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.rusk.KeysClient =
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
proto.rusk.KeysPromiseClient =
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
 *   !proto.rusk.GenerateKeysRequest,
 *   !proto.rusk.GenerateKeysResponse>}
 */
const methodDescriptor_Keys_GenerateKeys = new grpc.web.MethodDescriptor(
  '/rusk.Keys/GenerateKeys',
  grpc.web.MethodType.UNARY,
  proto.rusk.GenerateKeysRequest,
  proto.rusk.GenerateKeysResponse,
  /**
   * @param {!proto.rusk.GenerateKeysRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.GenerateKeysResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.GenerateKeysRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.GenerateKeysResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.GenerateKeysResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.KeysClient.prototype.generateKeys =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.Keys/GenerateKeys',
      request,
      metadata || {},
      methodDescriptor_Keys_GenerateKeys,
      callback);
};


/**
 * @param {!proto.rusk.GenerateKeysRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.GenerateKeysResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.KeysPromiseClient.prototype.generateKeys =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.Keys/GenerateKeys',
      request,
      metadata || {},
      methodDescriptor_Keys_GenerateKeys);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.PublicKey,
 *   !proto.rusk.StealthAddress>}
 */
const methodDescriptor_Keys_GenerateStealthAddress = new grpc.web.MethodDescriptor(
  '/rusk.Keys/GenerateStealthAddress',
  grpc.web.MethodType.UNARY,
  proto.rusk.PublicKey,
  proto.rusk.StealthAddress,
  /**
   * @param {!proto.rusk.PublicKey} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.StealthAddress.deserializeBinary
);


/**
 * @param {!proto.rusk.PublicKey} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.StealthAddress)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.StealthAddress>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.KeysClient.prototype.generateStealthAddress =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.Keys/GenerateStealthAddress',
      request,
      metadata || {},
      methodDescriptor_Keys_GenerateStealthAddress,
      callback);
};


/**
 * @param {!proto.rusk.PublicKey} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.StealthAddress>}
 *     Promise that resolves to the response
 */
proto.rusk.KeysPromiseClient.prototype.generateStealthAddress =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.Keys/GenerateStealthAddress',
      request,
      metadata || {},
      methodDescriptor_Keys_GenerateStealthAddress);
};


module.exports = proto.rusk;

