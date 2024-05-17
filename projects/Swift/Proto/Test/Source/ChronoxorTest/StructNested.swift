//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// FBE version: 1.14.5.0
//------------------------------------------------------------------------------

import Foundation
import ChronoxorFbe
import ChronoxorProto

public protocol StructNestedBase: StructOptionalBase {
    var f1000: EnumSimple { get set }
    var f1001: EnumSimple? { get set }
    var f1002: EnumTyped { get set }
    var f1003: EnumTyped? { get set }
    var f1004: FlagsSimple { get set }
    var f1005: FlagsSimple? { get set }
    var f1006: FlagsTyped { get set }
    var f1007: FlagsTyped? { get set }
    var f1008: StructSimple { get set }
    var f1009: StructSimple? { get set }
    var f1010: StructOptional { get set }
    var f1011: StructOptional? { get set }
}

public protocol StructNestedInheritance: StructOptionalInheritance {
    var parent: StructNested { get set }
}

extension StructNestedInheritance {
    public var parent: StructOptional {
        get { return parent.parent }
        set { parent.parent = newValue }
    }
    public var f1000: EnumSimple {
        get { return parent.f1000 }
        set { parent.f1000 = newValue }
    }
    public var f1001: EnumSimple? {
        get { return parent.f1001 }
        set { parent.f1001 = newValue }
    }
    public var f1002: EnumTyped {
        get { return parent.f1002 }
        set { parent.f1002 = newValue }
    }
    public var f1003: EnumTyped? {
        get { return parent.f1003 }
        set { parent.f1003 = newValue }
    }
    public var f1004: FlagsSimple {
        get { return parent.f1004 }
        set { parent.f1004 = newValue }
    }
    public var f1005: FlagsSimple? {
        get { return parent.f1005 }
        set { parent.f1005 = newValue }
    }
    public var f1006: FlagsTyped {
        get { return parent.f1006 }
        set { parent.f1006 = newValue }
    }
    public var f1007: FlagsTyped? {
        get { return parent.f1007 }
        set { parent.f1007 = newValue }
    }
    public var f1008: StructSimple {
        get { return parent.f1008 }
        set { parent.f1008 = newValue }
    }
    public var f1009: StructSimple? {
        get { return parent.f1009 }
        set { parent.f1009 = newValue }
    }
    public var f1010: StructOptional {
        get { return parent.f1010 }
        set { parent.f1010 = newValue }
    }
    public var f1011: StructOptional? {
        get { return parent.f1011 }
        set { parent.f1011 = newValue }
    }
}

public struct StructNested: StructNestedBase, StructOptionalInheritance, Comparable, Hashable, Codable {
    public var parent: StructOptional
    public var f1000: EnumSimple = ChronoxorTest.EnumSimple()
    public var f1001: EnumSimple? = nil
    public var f1002: EnumTyped = ChronoxorTest.EnumTyped.ENUM_VALUE_2
    public var f1003: EnumTyped? = nil
    public var f1004: FlagsSimple = ChronoxorTest.FlagsSimple()
    public var f1005: FlagsSimple? = nil
    public var f1006: FlagsTyped = FlagsTyped.fromSet(set: [FlagsTyped.FLAG_VALUE_2.value!, FlagsTyped.FLAG_VALUE_4.value!, FlagsTyped.FLAG_VALUE_6.value!])
    public var f1007: FlagsTyped? = nil
    public var f1008: StructSimple = ChronoxorTest.StructSimple()
    public var f1009: StructSimple? = nil
    public var f1010: StructOptional = ChronoxorTest.StructOptional()
    public var f1011: StructOptional? = nil

    public init() { parent = StructOptional() }
    public init(parent: StructOptional, f1000: EnumSimple, f1001: EnumSimple?, f1002: EnumTyped, f1003: EnumTyped?, f1004: FlagsSimple, f1005: FlagsSimple?, f1006: FlagsTyped, f1007: FlagsTyped?, f1008: StructSimple, f1009: StructSimple?, f1010: StructOptional, f1011: StructOptional?) {
        self.parent = parent

        self.f1000 = f1000
        self.f1001 = f1001
        self.f1002 = f1002
        self.f1003 = f1003
        self.f1004 = f1004
        self.f1005 = f1005
        self.f1006 = f1006
        self.f1007 = f1007
        self.f1008 = f1008
        self.f1009 = f1009
        self.f1010 = f1010
        self.f1011 = f1011
    }

    public init(other: StructNested) {
        parent = other.parent
        self.f1000 = other.f1000
        self.f1001 = other.f1001
        self.f1002 = other.f1002
        self.f1003 = other.f1003
        self.f1004 = other.f1004
        self.f1005 = other.f1005
        self.f1006 = other.f1006
        self.f1007 = other.f1007
        self.f1008 = other.f1008
        self.f1009 = other.f1009
        self.f1010 = other.f1010
        self.f1011 = other.f1011
    }

    public init(from decoder: Decoder) throws {
        parent = try StructOptional(from: decoder)
        let container = try decoder.container(keyedBy: CodingKeys.self)
        f1000 = try container.decode(ChronoxorTest.EnumSimple.self, forKey: .f1000)
        f1001 = try container.decode(ChronoxorTest.EnumSimple?.self, forKey: .f1001)
        f1002 = try container.decode(ChronoxorTest.EnumTyped.self, forKey: .f1002)
        f1003 = try container.decode(ChronoxorTest.EnumTyped?.self, forKey: .f1003)
        f1004 = try container.decode(ChronoxorTest.FlagsSimple.self, forKey: .f1004)
        f1005 = try container.decode(ChronoxorTest.FlagsSimple?.self, forKey: .f1005)
        f1006 = try container.decode(ChronoxorTest.FlagsTyped.self, forKey: .f1006)
        f1007 = try container.decode(ChronoxorTest.FlagsTyped?.self, forKey: .f1007)
        f1008 = try container.decode(ChronoxorTest.StructSimple.self, forKey: .f1008)
        f1009 = try container.decode(ChronoxorTest.StructSimple?.self, forKey: .f1009)
        f1010 = try container.decode(ChronoxorTest.StructOptional.self, forKey: .f1010)
        f1011 = try container.decode(ChronoxorTest.StructOptional?.self, forKey: .f1011)
    }

    public func clone() throws -> StructNested {
        // Serialize the struct to the FBE stream
        let writer = StructNestedModel()
        try _ = writer.serialize(value: self)

        // Deserialize the struct from the FBE stream
        let reader = StructNestedModel()
        reader.attach(buffer: writer.buffer)
        return reader.deserialize()
    }

    public static func < (lhs: StructNested, rhs: StructNested) -> Bool {
        return true
    }

    public static func == (lhs: StructNested, rhs: StructNested) -> Bool {
        return true
    }

    public func hash(into hasher: inout Hasher) {
        parent.hash(into: &hasher)
    }

    public var description: String {
        var sb = String()
        sb.append("StructNested(")
        sb.append(parent.description)
        sb.append(",f1000="); sb.append(f1000.description)
        sb.append(",f1001=");  if let f1001 = f1001 { sb.append(f1001.description) } else { sb.append("null") }
        sb.append(",f1002="); sb.append(f1002.description)
        sb.append(",f1003=");  if let f1003 = f1003 { sb.append(f1003.description) } else { sb.append("null") }
        sb.append(",f1004="); sb.append(f1004.description)
        sb.append(",f1005=");  if let f1005 = f1005 { sb.append(f1005.description) } else { sb.append("null") }
        sb.append(",f1006="); sb.append(f1006.description)
        sb.append(",f1007=");  if let f1007 = f1007 { sb.append(f1007.description) } else { sb.append("null") }
        sb.append(",f1008="); sb.append(f1008.description)
        sb.append(",f1009=");  if let f1009 = f1009 { sb.append(f1009.description) } else { sb.append("null") }
        sb.append(",f1010="); sb.append(f1010.description)
        sb.append(",f1011=");  if let f1011 = f1011 { sb.append(f1011.description) } else { sb.append("null") }
        sb.append(")")
        return sb
    }
    private enum CodingKeys: String, CodingKey {
        case f1000
        case f1001
        case f1002
        case f1003
        case f1004
        case f1005
        case f1006
        case f1007
        case f1008
        case f1009
        case f1010
        case f1011
    }

    public func encode(to encoder: Encoder) throws {
        try parent.encode(to: encoder)
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encode(f1000, forKey: .f1000)
        try container.encode(f1001, forKey: .f1001)
        try container.encode(f1002, forKey: .f1002)
        try container.encode(f1003, forKey: .f1003)
        try container.encode(f1004, forKey: .f1004)
        try container.encode(f1005, forKey: .f1005)
        try container.encode(f1006, forKey: .f1006)
        try container.encode(f1007, forKey: .f1007)
        try container.encode(f1008, forKey: .f1008)
        try container.encode(f1009, forKey: .f1009)
        try container.encode(f1010, forKey: .f1010)
        try container.encode(f1011, forKey: .f1011)
    }

    public func toJson() throws -> String {
        return String(data: try JSONEncoder().encode(self), encoding: .utf8)!
    }

    public static func fromJson(_ json: String) throws -> StructNested {
        return try JSONDecoder().decode(StructNested.self, from: json.data(using: .utf8)!)
    }
}
