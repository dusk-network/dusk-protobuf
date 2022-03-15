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


var echo_pb = require('./echo_pb.js')

var network_pb = require('./network_pb.js')

var prover_pb = require('./prover_pb.js')

var provisioner_pb = require('./provisioner_pb.js')

var transaction_pb = require('./transaction_pb.js')
const proto = {};
proto.rusk = require('./state_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.rusk.StateClient =
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
proto.rusk.StatePromiseClient =
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
 *   !proto.rusk.EchoRequest,
 *   !proto.rusk.EchoResponse>}
 */
const methodDescriptor_State_Echo = new grpc.web.MethodDescriptor(
  '/rusk.State/Echo',
  grpc.web.MethodType.UNARY,
  echo_pb.EchoRequest,
  echo_pb.EchoResponse,
  /**
   * @param {!proto.rusk.EchoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  echo_pb.EchoResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.EchoRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.EchoResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.EchoResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.StateClient.prototype.echo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.State/Echo',
      request,
      metadata || {},
      methodDescriptor_State_Echo,
      callback);
};


/**
 * @param {!proto.rusk.EchoRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.EchoResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.StatePromiseClient.prototype.echo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.State/Echo',
      request,
      metadata || {},
      methodDescriptor_State_Echo);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.PreverifyRequest,
 *   !proto.rusk.PreverifyResponse>}
 */
const methodDescriptor_State_Preverify = new grpc.web.MethodDescriptor(
  '/rusk.State/Preverify',
  grpc.web.MethodType.UNARY,
  proto.rusk.PreverifyRequest,
  proto.rusk.PreverifyResponse,
  /**
   * @param {!proto.rusk.PreverifyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.PreverifyResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.PreverifyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.PreverifyResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.PreverifyResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.StateClient.prototype.preverify =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.State/Preverify',
      request,
      metadata || {},
      methodDescriptor_State_Preverify,
      callback);
};


/**
 * @param {!proto.rusk.PreverifyRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.PreverifyResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.StatePromiseClient.prototype.preverify =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.State/Preverify',
      request,
      metadata || {},
      methodDescriptor_State_Preverify);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.ExecuteStateTransitionRequest,
 *   !proto.rusk.ExecuteStateTransitionResponse>}
 */
const methodDescriptor_State_ExecuteStateTransition = new grpc.web.MethodDescriptor(
  '/rusk.State/ExecuteStateTransition',
  grpc.web.MethodType.UNARY,
  proto.rusk.ExecuteStateTransitionRequest,
  proto.rusk.ExecuteStateTransitionResponse,
  /**
   * @param {!proto.rusk.ExecuteStateTransitionRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.ExecuteStateTransitionResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.ExecuteStateTransitionRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.ExecuteStateTransitionResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.ExecuteStateTransitionResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.StateClient.prototype.executeStateTransition =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.State/ExecuteStateTransition',
      request,
      metadata || {},
      methodDescriptor_State_ExecuteStateTransition,
      callback);
};


/**
 * @param {!proto.rusk.ExecuteStateTransitionRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.ExecuteStateTransitionResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.StatePromiseClient.prototype.executeStateTransition =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.State/ExecuteStateTransition',
      request,
      metadata || {},
      methodDescriptor_State_ExecuteStateTransition);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.VerifyStateTransitionRequest,
 *   !proto.rusk.VerifyStateTransitionResponse>}
 */
const methodDescriptor_State_VerifyStateTransition = new grpc.web.MethodDescriptor(
  '/rusk.State/VerifyStateTransition',
  grpc.web.MethodType.UNARY,
  proto.rusk.VerifyStateTransitionRequest,
  proto.rusk.VerifyStateTransitionResponse,
  /**
   * @param {!proto.rusk.VerifyStateTransitionRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.VerifyStateTransitionResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.VerifyStateTransitionRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.VerifyStateTransitionResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.VerifyStateTransitionResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.StateClient.prototype.verifyStateTransition =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.State/VerifyStateTransition',
      request,
      metadata || {},
      methodDescriptor_State_VerifyStateTransition,
      callback);
};


/**
 * @param {!proto.rusk.VerifyStateTransitionRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.VerifyStateTransitionResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.StatePromiseClient.prototype.verifyStateTransition =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.State/VerifyStateTransition',
      request,
      metadata || {},
      methodDescriptor_State_VerifyStateTransition);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.StateTransitionRequest,
 *   !proto.rusk.StateTransitionResponse>}
 */
const methodDescriptor_State_Accept = new grpc.web.MethodDescriptor(
  '/rusk.State/Accept',
  grpc.web.MethodType.UNARY,
  proto.rusk.StateTransitionRequest,
  proto.rusk.StateTransitionResponse,
  /**
   * @param {!proto.rusk.StateTransitionRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.StateTransitionResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.StateTransitionRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.StateTransitionResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.StateTransitionResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.StateClient.prototype.accept =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.State/Accept',
      request,
      metadata || {},
      methodDescriptor_State_Accept,
      callback);
};


/**
 * @param {!proto.rusk.StateTransitionRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.StateTransitionResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.StatePromiseClient.prototype.accept =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.State/Accept',
      request,
      metadata || {},
      methodDescriptor_State_Accept);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.StateTransitionRequest,
 *   !proto.rusk.StateTransitionResponse>}
 */
const methodDescriptor_State_Finalize = new grpc.web.MethodDescriptor(
  '/rusk.State/Finalize',
  grpc.web.MethodType.UNARY,
  proto.rusk.StateTransitionRequest,
  proto.rusk.StateTransitionResponse,
  /**
   * @param {!proto.rusk.StateTransitionRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.StateTransitionResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.StateTransitionRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.StateTransitionResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.StateTransitionResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.StateClient.prototype.finalize =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.State/Finalize',
      request,
      metadata || {},
      methodDescriptor_State_Finalize,
      callback);
};


/**
 * @param {!proto.rusk.StateTransitionRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.StateTransitionResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.StatePromiseClient.prototype.finalize =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.State/Finalize',
      request,
      metadata || {},
      methodDescriptor_State_Finalize);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.RevertRequest,
 *   !proto.rusk.RevertResponse>}
 */
const methodDescriptor_State_Revert = new grpc.web.MethodDescriptor(
  '/rusk.State/Revert',
  grpc.web.MethodType.UNARY,
  proto.rusk.RevertRequest,
  proto.rusk.RevertResponse,
  /**
   * @param {!proto.rusk.RevertRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.RevertResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.RevertRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.RevertResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.RevertResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.StateClient.prototype.revert =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.State/Revert',
      request,
      metadata || {},
      methodDescriptor_State_Revert,
      callback);
};


/**
 * @param {!proto.rusk.RevertRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.RevertResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.StatePromiseClient.prototype.revert =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.State/Revert',
      request,
      metadata || {},
      methodDescriptor_State_Revert);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.PersistRequest,
 *   !proto.rusk.PersistResponse>}
 */
const methodDescriptor_State_Persist = new grpc.web.MethodDescriptor(
  '/rusk.State/Persist',
  grpc.web.MethodType.UNARY,
  proto.rusk.PersistRequest,
  proto.rusk.PersistResponse,
  /**
   * @param {!proto.rusk.PersistRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.PersistResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.PersistRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.PersistResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.PersistResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.StateClient.prototype.persist =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.State/Persist',
      request,
      metadata || {},
      methodDescriptor_State_Persist,
      callback);
};


/**
 * @param {!proto.rusk.PersistRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.PersistResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.StatePromiseClient.prototype.persist =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.State/Persist',
      request,
      metadata || {},
      methodDescriptor_State_Persist);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.GetProvisionersRequest,
 *   !proto.rusk.GetProvisionersResponse>}
 */
const methodDescriptor_State_GetProvisioners = new grpc.web.MethodDescriptor(
  '/rusk.State/GetProvisioners',
  grpc.web.MethodType.UNARY,
  proto.rusk.GetProvisionersRequest,
  proto.rusk.GetProvisionersResponse,
  /**
   * @param {!proto.rusk.GetProvisionersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.GetProvisionersResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.GetProvisionersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.GetProvisionersResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.GetProvisionersResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.StateClient.prototype.getProvisioners =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.State/GetProvisioners',
      request,
      metadata || {},
      methodDescriptor_State_GetProvisioners,
      callback);
};


/**
 * @param {!proto.rusk.GetProvisionersRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.GetProvisionersResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.StatePromiseClient.prototype.getProvisioners =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.State/GetProvisioners',
      request,
      metadata || {},
      methodDescriptor_State_GetProvisioners);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.GetStateRootRequest,
 *   !proto.rusk.GetStateRootResponse>}
 */
const methodDescriptor_State_GetStateRoot = new grpc.web.MethodDescriptor(
  '/rusk.State/GetStateRoot',
  grpc.web.MethodType.UNARY,
  proto.rusk.GetStateRootRequest,
  proto.rusk.GetStateRootResponse,
  /**
   * @param {!proto.rusk.GetStateRootRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.GetStateRootResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.GetStateRootRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.GetStateRootResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.GetStateRootResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.StateClient.prototype.getStateRoot =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.State/GetStateRoot',
      request,
      metadata || {},
      methodDescriptor_State_GetStateRoot,
      callback);
};


/**
 * @param {!proto.rusk.GetStateRootRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.GetStateRootResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.StatePromiseClient.prototype.getStateRoot =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.State/GetStateRoot',
      request,
      metadata || {},
      methodDescriptor_State_GetStateRoot);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.GetNotesOwnedByRequest,
 *   !proto.rusk.GetNotesOwnedByResponse>}
 */
const methodDescriptor_State_GetNotesOwnedBy = new grpc.web.MethodDescriptor(
  '/rusk.State/GetNotesOwnedBy',
  grpc.web.MethodType.UNARY,
  proto.rusk.GetNotesOwnedByRequest,
  proto.rusk.GetNotesOwnedByResponse,
  /**
   * @param {!proto.rusk.GetNotesOwnedByRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.GetNotesOwnedByResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.GetNotesOwnedByRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.GetNotesOwnedByResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.GetNotesOwnedByResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.StateClient.prototype.getNotesOwnedBy =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.State/GetNotesOwnedBy',
      request,
      metadata || {},
      methodDescriptor_State_GetNotesOwnedBy,
      callback);
};


/**
 * @param {!proto.rusk.GetNotesOwnedByRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.GetNotesOwnedByResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.StatePromiseClient.prototype.getNotesOwnedBy =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.State/GetNotesOwnedBy',
      request,
      metadata || {},
      methodDescriptor_State_GetNotesOwnedBy);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.GetAnchorRequest,
 *   !proto.rusk.GetAnchorResponse>}
 */
const methodDescriptor_State_GetAnchor = new grpc.web.MethodDescriptor(
  '/rusk.State/GetAnchor',
  grpc.web.MethodType.UNARY,
  proto.rusk.GetAnchorRequest,
  proto.rusk.GetAnchorResponse,
  /**
   * @param {!proto.rusk.GetAnchorRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.GetAnchorResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.GetAnchorRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.GetAnchorResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.GetAnchorResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.StateClient.prototype.getAnchor =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.State/GetAnchor',
      request,
      metadata || {},
      methodDescriptor_State_GetAnchor,
      callback);
};


/**
 * @param {!proto.rusk.GetAnchorRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.GetAnchorResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.StatePromiseClient.prototype.getAnchor =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.State/GetAnchor',
      request,
      metadata || {},
      methodDescriptor_State_GetAnchor);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.GetOpeningRequest,
 *   !proto.rusk.GetOpeningResponse>}
 */
const methodDescriptor_State_GetOpening = new grpc.web.MethodDescriptor(
  '/rusk.State/GetOpening',
  grpc.web.MethodType.UNARY,
  proto.rusk.GetOpeningRequest,
  proto.rusk.GetOpeningResponse,
  /**
   * @param {!proto.rusk.GetOpeningRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.GetOpeningResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.GetOpeningRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.GetOpeningResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.GetOpeningResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.StateClient.prototype.getOpening =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.State/GetOpening',
      request,
      metadata || {},
      methodDescriptor_State_GetOpening,
      callback);
};


/**
 * @param {!proto.rusk.GetOpeningRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.GetOpeningResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.StatePromiseClient.prototype.getOpening =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.State/GetOpening',
      request,
      metadata || {},
      methodDescriptor_State_GetOpening);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.GetStakeRequest,
 *   !proto.rusk.GetStakeResponse>}
 */
const methodDescriptor_State_GetStake = new grpc.web.MethodDescriptor(
  '/rusk.State/GetStake',
  grpc.web.MethodType.UNARY,
  proto.rusk.GetStakeRequest,
  proto.rusk.GetStakeResponse,
  /**
   * @param {!proto.rusk.GetStakeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.GetStakeResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.GetStakeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.GetStakeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.GetStakeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.StateClient.prototype.getStake =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.State/GetStake',
      request,
      metadata || {},
      methodDescriptor_State_GetStake,
      callback);
};


/**
 * @param {!proto.rusk.GetStakeRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.GetStakeResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.StatePromiseClient.prototype.getStake =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.State/GetStake',
      request,
      metadata || {},
      methodDescriptor_State_GetStake);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rusk.FindExistingNullifiersRequest,
 *   !proto.rusk.FindExistingNullifiersResponse>}
 */
const methodDescriptor_State_FindExistingNullifiers = new grpc.web.MethodDescriptor(
  '/rusk.State/FindExistingNullifiers',
  grpc.web.MethodType.UNARY,
  proto.rusk.FindExistingNullifiersRequest,
  proto.rusk.FindExistingNullifiersResponse,
  /**
   * @param {!proto.rusk.FindExistingNullifiersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.rusk.FindExistingNullifiersResponse.deserializeBinary
);


/**
 * @param {!proto.rusk.FindExistingNullifiersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.rusk.FindExistingNullifiersResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rusk.FindExistingNullifiersResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rusk.StateClient.prototype.findExistingNullifiers =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rusk.State/FindExistingNullifiers',
      request,
      metadata || {},
      methodDescriptor_State_FindExistingNullifiers,
      callback);
};


/**
 * @param {!proto.rusk.FindExistingNullifiersRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rusk.FindExistingNullifiersResponse>}
 *     Promise that resolves to the response
 */
proto.rusk.StatePromiseClient.prototype.findExistingNullifiers =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rusk.State/FindExistingNullifiers',
      request,
      metadata || {},
      methodDescriptor_State_FindExistingNullifiers);
};


module.exports = proto.rusk;

