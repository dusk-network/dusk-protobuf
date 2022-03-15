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
proto.rusk = require('./reward_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.rusk.RewardClient =
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
proto.rusk.RewardPromiseClient =
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
 *   !proto.rusk.BN256Point,
 *   !proto.rusk.GetBalanceResponse>}
 */
const methodDescriptor_Reward_GetBalance = new grpc.web.MethodDescriptor(
  '/rusk.Reward/GetBalance',
  grpc.web.MethodType.UNARY,
  proto.rusk.BN256Point,
  proto.rusk.GetBalanceResponse,
  /**
   * @param {!proto.rusk.BN256Point} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.GetBalanceResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.BN256Point} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.GetBalanceResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.GetBalanceResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.RewardClient.prototype.getBalance =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.Reward/GetBalance',
      request,
      metadata || {},
      methodDescriptor_Reward_GetBalance,
      callback);
};


/**
 * @param {!proto.rusk.BN256Point} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.GetBalanceResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.RewardPromiseClient.prototype.getBalance =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.Reward/GetBalance',
      request,
      metadata || {},
      methodDescriptor_Reward_GetBalance);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.BN256Point,
 *   !proto.rusk.GetWithdrawalTimeResponse>}
 */
const methodDescriptor_Reward_GetWithdrawalTime = new grpc.web.MethodDescriptor(
  '/rusk.Reward/GetWithdrawalTime',
  grpc.web.MethodType.UNARY,
  proto.rusk.BN256Point,
  proto.rusk.GetWithdrawalTimeResponse,
  /**
   * @param {!proto.rusk.BN256Point} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.GetWithdrawalTimeResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.BN256Point} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.GetWithdrawalTimeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.GetWithdrawalTimeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.RewardClient.prototype.getWithdrawalTime =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.Reward/GetWithdrawalTime',
      request,
      metadata || {},
      methodDescriptor_Reward_GetWithdrawalTime,
      callback);
};


/**
 * @param {!proto.rusk.BN256Point} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.GetWithdrawalTimeResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.RewardPromiseClient.prototype.getWithdrawalTime =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.Reward/GetWithdrawalTime',
      request,
      metadata || {},
      methodDescriptor_Reward_GetWithdrawalTime);
};


module.exports = proto.rusk;

