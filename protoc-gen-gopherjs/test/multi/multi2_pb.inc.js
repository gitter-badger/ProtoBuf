/**
 * @fileoverview
 * @enhanceable
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = $global;
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.multitest.Multi2', null, global);
goog.exportSymbol('proto.multitest.Multi2.Color', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.multitest.Multi2 = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.multitest.Multi2, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.multitest.Multi2.displayName = 'proto.multitest.Multi2';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.multitest.Multi2.prototype.toObject = function(opt_includeInstance) {
  return proto.multitest.Multi2.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.multitest.Multi2} msg The msg instance to transform.
 * @return {!Object}
 */
proto.multitest.Multi2.toObject = function(includeInstance, msg) {
  var f, obj = {
    requiredValue: jspb.Message.getFieldWithDefault(msg, 1, 0),
    color: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.multitest.Multi2}
 */
proto.multitest.Multi2.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.multitest.Multi2;
  return proto.multitest.Multi2.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.multitest.Multi2} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.multitest.Multi2}
 */
proto.multitest.Multi2.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setRequiredValue(value);
      break;
    case 2:
      var value = /** @type {!proto.multitest.Multi2.Color} */ (reader.readEnum());
      msg.setColor(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.multitest.Multi2.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.multitest.Multi2.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.multitest.Multi2} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.multitest.Multi2.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRequiredValue();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getColor();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.multitest.Multi2.Color = {
  BLUE: 0,
  GREEN: 1,
  RED: 2
};

/**
 * optional int32 required_value = 1;
 * @return {number}
 */
proto.multitest.Multi2.prototype.getRequiredValue = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.multitest.Multi2.prototype.setRequiredValue = function(value) {
  jspb.Message.setField(this, 1, value);
};


/**
 * optional Color color = 2;
 * @return {!proto.multitest.Multi2.Color}
 */
proto.multitest.Multi2.prototype.getColor = function() {
  return /** @type {!proto.multitest.Multi2.Color} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {!proto.multitest.Multi2.Color} value */
proto.multitest.Multi2.prototype.setColor = function(value) {
  jspb.Message.setField(this, 2, value);
};


