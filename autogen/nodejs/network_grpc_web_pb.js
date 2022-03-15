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
proto.rusk = require('./network_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.rusk.NetworkClient =
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
proto.rusk.NetworkPromiseClient =
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
 *   !proto.rusk.Null,
 *   !proto.rusk.Message>}
 */
const methodDescriptor_Network_Listen = new grpc.web.MethodDescriptor(
  '/rusk.Network/Listen',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.rusk.Null,
  proto.rusk.Message,
  /**
   * @param {!proto.rusk.Null} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.Message.deserializeBinary
);


/**
 * @param {!proto.rusk.Null} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.Message>}
 *     The XHR Node Readable Stream
 */
proto.rusk.NetworkClient.prototype.listen =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/rusk.Network/Listen',
      request,
      metadata || {},
      methodDescriptor_Network_Listen);
};


/**
 * @param {!proto.rusk.Null} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.Message>}
 *     The XHR Node Readable Stream
 */
proto.rusk.NetworkPromiseClient.prototype.listen =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/rusk.Network/Listen',
      request,
      metadata || {},
      methodDescriptor_Network_Listen);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.BroadcastMessage,
 *   !proto.rusk.Null>}
 */
const methodDescriptor_Network_Broadcast = new grpc.web.MethodDescriptor(
  '/rusk.Network/Broadcast',
  grpc.web.MethodType.UNARY,
  proto.rusk.BroadcastMessage,
  proto.rusk.Null,
  /**
   * @param {!proto.rusk.BroadcastMessage} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.Null.deserializeBinary
);


/**
 * @param {!proto.rusk.BroadcastMessage} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.Null)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.Null>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.NetworkClient.prototype.broadcast =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.Network/Broadcast',
      request,
      metadata || {},
      methodDescriptor_Network_Broadcast,
      callback);
};


/**
 * @param {!proto.rusk.BroadcastMessage} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.Null>}
 *     Promise that resolves to the response
 */
proto.rusk.NetworkPromiseClient.prototype.broadcast =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.Network/Broadcast',
      request,
      metadata || {},
      methodDescriptor_Network_Broadcast);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.PropagateMessage,
 *   !proto.rusk.Null>}
 */
const methodDescriptor_Network_Propagate = new grpc.web.MethodDescriptor(
  '/rusk.Network/Propagate',
  grpc.web.MethodType.UNARY,
  proto.rusk.PropagateMessage,
  proto.rusk.Null,
  /**
   * @param {!proto.rusk.PropagateMessage} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.Null.deserializeBinary
);


/**
 * @param {!proto.rusk.PropagateMessage} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.Null)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.Null>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.NetworkClient.prototype.propagate =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.Network/Propagate',
      request,
      metadata || {},
      methodDescriptor_Network_Propagate,
      callback);
};


/**
 * @param {!proto.rusk.PropagateMessage} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.Null>}
 *     Promise that resolves to the response
 */
proto.rusk.NetworkPromiseClient.prototype.propagate =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.Network/Propagate',
      request,
      metadata || {},
      methodDescriptor_Network_Propagate);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.SendMessage,
 *   !proto.rusk.Null>}
 */
const methodDescriptor_Network_Send = new grpc.web.MethodDescriptor(
  '/rusk.Network/Send',
  grpc.web.MethodType.UNARY,
  proto.rusk.SendMessage,
  proto.rusk.Null,
  /**
   * @param {!proto.rusk.SendMessage} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.Null.deserializeBinary
);


/**
 * @param {!proto.rusk.SendMessage} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.Null)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.Null>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.NetworkClient.prototype.send =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.Network/Send',
      request,
      metadata || {},
      methodDescriptor_Network_Send,
      callback);
};


/**
 * @param {!proto.rusk.SendMessage} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.Null>}
 *     Promise that resolves to the response
 */
proto.rusk.NetworkPromiseClient.prototype.send =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.Network/Send',
      request,
      metadata || {},
      methodDescriptor_Network_Send);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.AliveNodesRequest,
 *   !proto.rusk.AliveNodesResponse>}
 */
const methodDescriptor_Network_AliveNodes = new grpc.web.MethodDescriptor(
  '/rusk.Network/AliveNodes',
  grpc.web.MethodType.UNARY,
  proto.rusk.AliveNodesRequest,
  proto.rusk.AliveNodesResponse,
  /**
   * @param {!proto.rusk.AliveNodesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.AliveNodesResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.AliveNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.AliveNodesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.AliveNodesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.NetworkClient.prototype.aliveNodes =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.Network/AliveNodes',
      request,
      metadata || {},
      methodDescriptor_Network_AliveNodes,
      callback);
};


/**
 * @param {!proto.rusk.AliveNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.AliveNodesResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.NetworkPromiseClient.prototype.aliveNodes =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.Network/AliveNodes',
      request,
      metadata || {},
      methodDescriptor_Network_AliveNodes);
};


module.exports = proto.rusk;

