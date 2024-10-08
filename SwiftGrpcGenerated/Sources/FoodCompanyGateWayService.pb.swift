// DO NOT EDIT.
// swift-format-ignore-file
// swiftlint:disable all
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: FoodCompanyGateWayService.proto
//
// For information on using the generated types, please see the documentation:
//   https://github.com/apple/swift-protobuf/

import SwiftProtobuf

// If the compiler emits an error on this type, it is because this file
// was generated by a version of the `protoc` Swift plug-in that is
// incompatible with the version of SwiftProtobuf to which you are linking.
// Please ensure that you are building against the same version of the API
// that was used to generate this file.
fileprivate struct _GeneratedWithProtocGenSwiftVersion: SwiftProtobuf.ProtobufAPIVersionCheck {
  struct _2: SwiftProtobuf.ProtobufAPIVersion_2 {}
  typealias Version = _2
}

public struct GetOrderStatusRequest: Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var id: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct GetOrderStatusResponse: Sendable {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var status: GetOrderStatusResponse.Status = .done

  public var owner: String {
    get {return _owner ?? String()}
    set {_owner = newValue}
  }
  /// Returns true if `owner` has been explicitly set.
  public var hasOwner: Bool {return self._owner != nil}
  /// Clears the value of `owner`. Subsequent reads from it will return its default value.
  public mutating func clearOwner() {self._owner = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public enum Status: SwiftProtobuf.Enum, Swift.CaseIterable {
    public typealias RawValue = Int
    case done // = 0
    case pending // = 1
    case cancelled // = 2
    case UNRECOGNIZED(Int)

    public init() {
      self = .done
    }

    public init?(rawValue: Int) {
      switch rawValue {
      case 0: self = .done
      case 1: self = .pending
      case 2: self = .cancelled
      default: self = .UNRECOGNIZED(rawValue)
      }
    }

    public var rawValue: Int {
      switch self {
      case .done: return 0
      case .pending: return 1
      case .cancelled: return 2
      case .UNRECOGNIZED(let i): return i
      }
    }

    // The compiler won't synthesize support with the UNRECOGNIZED case.
    public static let allCases: [GetOrderStatusResponse.Status] = [
      .done,
      .pending,
      .cancelled,
    ]

  }

  public init() {}

  fileprivate var _owner: String? = nil
}

// MARK: - Code below here is support for the SwiftProtobuf runtime.

extension GetOrderStatusRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = "GetOrderStatusRequest"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "id"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.id) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.id.isEmpty {
      try visitor.visitSingularStringField(value: self.id, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: GetOrderStatusRequest, rhs: GetOrderStatusRequest) -> Bool {
    if lhs.id != rhs.id {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension GetOrderStatusResponse: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = "GetOrderStatusResponse"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "status"),
    2: .same(proto: "owner"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularEnumField(value: &self.status) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self._owner) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    if self.status != .done {
      try visitor.visitSingularEnumField(value: self.status, fieldNumber: 1)
    }
    try { if let v = self._owner {
      try visitor.visitSingularStringField(value: v, fieldNumber: 2)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: GetOrderStatusResponse, rhs: GetOrderStatusResponse) -> Bool {
    if lhs.status != rhs.status {return false}
    if lhs._owner != rhs._owner {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension GetOrderStatusResponse.Status: SwiftProtobuf._ProtoNameProviding {
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "Done"),
    1: .same(proto: "Pending"),
    2: .same(proto: "Cancelled"),
  ]
}
