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
proto.rusk = require('./prover_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.rusk.ProverClient =
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
proto.rusk.ProverPromiseClient =
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
 *   !proto.rusk.ExecuteProverRequest,
 *   !proto.rusk.ExecuteProverResponse>}
 */
const methodDescriptor_Prover_ProveExecute = new grpc.web.MethodDescriptor(
  '/rusk.Prover/ProveExecute',
  grpc.web.MethodType.UNARY,
  proto.rusk.ExecuteProverRequest,
  proto.rusk.ExecuteProverResponse,
  /**
   * @param {!proto.rusk.ExecuteProverRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.ExecuteProverResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.ExecuteProverRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.ExecuteProverResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.ExecuteProverResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.ProverClient.prototype.proveExecute =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.Prover/ProveExecute',
      request,
      metadata || {},
      methodDescriptor_Prover_ProveExecute,
      callback);
};


/**
 * @param {!proto.rusk.ExecuteProverRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.ExecuteProverResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.ProverPromiseClient.prototype.proveExecute =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.Prover/ProveExecute',
      request,
      metadata || {},
      methodDescriptor_Prover_ProveExecute);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.StctProverRequest,
 *   !proto.rusk.StctProverResponse>}
 */
const methodDescriptor_Prover_ProveStct = new grpc.web.MethodDescriptor(
  '/rusk.Prover/ProveStct',
  grpc.web.MethodType.UNARY,
  proto.rusk.StctProverRequest,
  proto.rusk.StctProverResponse,
  /**
   * @param {!proto.rusk.StctProverRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.StctProverResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.StctProverRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.StctProverResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.StctProverResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.ProverClient.prototype.proveStct =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.Prover/ProveStct',
      request,
      metadata || {},
      methodDescriptor_Prover_ProveStct,
      callback);
};


/**
 * @param {!proto.rusk.StctProverRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.StctProverResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.ProverPromiseClient.prototype.proveStct =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.Prover/ProveStct',
      request,
      metadata || {},
      methodDescriptor_Prover_ProveStct);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.StcoProverRequest,
 *   !proto.rusk.StcoProverResponse>}
 */
const methodDescriptor_Prover_ProveStco = new grpc.web.MethodDescriptor(
  '/rusk.Prover/ProveStco',
  grpc.web.MethodType.UNARY,
  proto.rusk.StcoProverRequest,
  proto.rusk.StcoProverResponse,
  /**
   * @param {!proto.rusk.StcoProverRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.StcoProverResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.StcoProverRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.StcoProverResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.StcoProverResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.ProverClient.prototype.proveStco =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.Prover/ProveStco',
      request,
      metadata || {},
      methodDescriptor_Prover_ProveStco,
      callback);
};


/**
 * @param {!proto.rusk.StcoProverRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.StcoProverResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.ProverPromiseClient.prototype.proveStco =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.Prover/ProveStco',
      request,
      metadata || {},
      methodDescriptor_Prover_ProveStco);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.WfctProverRequest,
 *   !proto.rusk.WfctProverResponse>}
 */
const methodDescriptor_Prover_ProveWfct = new grpc.web.MethodDescriptor(
  '/rusk.Prover/ProveWfct',
  grpc.web.MethodType.UNARY,
  proto.rusk.WfctProverRequest,
  proto.rusk.WfctProverResponse,
  /**
   * @param {!proto.rusk.WfctProverRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.WfctProverResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.WfctProverRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.WfctProverResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.WfctProverResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.ProverClient.prototype.proveWfct =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.Prover/ProveWfct',
      request,
      metadata || {},
      methodDescriptor_Prover_ProveWfct,
      callback);
};


/**
 * @param {!proto.rusk.WfctProverRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.WfctProverResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.ProverPromiseClient.prototype.proveWfct =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.Prover/ProveWfct',
      request,
      metadata || {},
      methodDescriptor_Prover_ProveWfct);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.WfcoProverRequest,
 *   !proto.rusk.WfcoProverResponse>}
 */
const methodDescriptor_Prover_ProveWfco = new grpc.web.MethodDescriptor(
  '/rusk.Prover/ProveWfco',
  grpc.web.MethodType.UNARY,
  proto.rusk.WfcoProverRequest,
  proto.rusk.WfcoProverResponse,
  /**
   * @param {!proto.rusk.WfcoProverRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.WfcoProverResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.WfcoProverRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.WfcoProverResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.WfcoProverResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.ProverClient.prototype.proveWfco =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.Prover/ProveWfco',
      request,
      metadata || {},
      methodDescriptor_Prover_ProveWfco,
      callback);
};


/**
 * @param {!proto.rusk.WfcoProverRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.WfcoProverResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.ProverPromiseClient.prototype.proveWfco =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.Prover/ProveWfco',
      request,
      metadata || {},
      methodDescriptor_Prover_ProveWfco);
};


module.exports = proto.rusk;

