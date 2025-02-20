// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package profile

import (
	fmt "fmt"
	io "io"
	reflect "reflect"
	sync "sync"

	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

var (
	md_Profile               protoreflect.MessageDescriptor
	fd_Profile_username      protoreflect.FieldDescriptor
	fd_Profile_pubKey        protoreflect.FieldDescriptor
	fd_Profile_telegramId    protoreflect.FieldDescriptor
	fd_Profile_cosmosAddress protoreflect.FieldDescriptor
)

func init() {
	file_habbitchain_profile_profile_proto_init()
	md_Profile = File_habbitchain_profile_profile_proto.Messages().ByName("Profile")
	fd_Profile_username = md_Profile.Fields().ByName("username")
	fd_Profile_pubKey = md_Profile.Fields().ByName("pubKey")
	fd_Profile_telegramId = md_Profile.Fields().ByName("telegramId")
	fd_Profile_cosmosAddress = md_Profile.Fields().ByName("cosmosAddress")
}

var _ protoreflect.Message = (*fastReflection_Profile)(nil)

type fastReflection_Profile Profile

func (x *Profile) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Profile)(x)
}

func (x *Profile) slowProtoReflect() protoreflect.Message {
	mi := &file_habbitchain_profile_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Profile_messageType fastReflection_Profile_messageType
var _ protoreflect.MessageType = fastReflection_Profile_messageType{}

type fastReflection_Profile_messageType struct{}

func (x fastReflection_Profile_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Profile)(nil)
}
func (x fastReflection_Profile_messageType) New() protoreflect.Message {
	return new(fastReflection_Profile)
}
func (x fastReflection_Profile_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Profile
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Profile) Descriptor() protoreflect.MessageDescriptor {
	return md_Profile
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Profile) Type() protoreflect.MessageType {
	return _fastReflection_Profile_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Profile) New() protoreflect.Message {
	return new(fastReflection_Profile)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Profile) Interface() protoreflect.ProtoMessage {
	return (*Profile)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Profile) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Username != "" {
		value := protoreflect.ValueOfString(x.Username)
		if !f(fd_Profile_username, value) {
			return
		}
	}
	if x.PubKey != "" {
		value := protoreflect.ValueOfString(x.PubKey)
		if !f(fd_Profile_pubKey, value) {
			return
		}
	}
	if x.TelegramId != int32(0) {
		value := protoreflect.ValueOfInt32(x.TelegramId)
		if !f(fd_Profile_telegramId, value) {
			return
		}
	}
	if x.CosmosAddress != "" {
		value := protoreflect.ValueOfString(x.CosmosAddress)
		if !f(fd_Profile_cosmosAddress, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Profile) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "habbitchain.profile.Profile.username":
		return x.Username != ""
	case "habbitchain.profile.Profile.pubKey":
		return x.PubKey != ""
	case "habbitchain.profile.Profile.telegramId":
		return x.TelegramId != int32(0)
	case "habbitchain.profile.Profile.cosmosAddress":
		return x.CosmosAddress != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: habbitchain.profile.Profile"))
		}
		panic(fmt.Errorf("message habbitchain.profile.Profile does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Profile) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "habbitchain.profile.Profile.username":
		x.Username = ""
	case "habbitchain.profile.Profile.pubKey":
		x.PubKey = ""
	case "habbitchain.profile.Profile.telegramId":
		x.TelegramId = int32(0)
	case "habbitchain.profile.Profile.cosmosAddress":
		x.CosmosAddress = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: habbitchain.profile.Profile"))
		}
		panic(fmt.Errorf("message habbitchain.profile.Profile does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Profile) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "habbitchain.profile.Profile.username":
		value := x.Username
		return protoreflect.ValueOfString(value)
	case "habbitchain.profile.Profile.pubKey":
		value := x.PubKey
		return protoreflect.ValueOfString(value)
	case "habbitchain.profile.Profile.telegramId":
		value := x.TelegramId
		return protoreflect.ValueOfInt32(value)
	case "habbitchain.profile.Profile.cosmosAddress":
		value := x.CosmosAddress
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: habbitchain.profile.Profile"))
		}
		panic(fmt.Errorf("message habbitchain.profile.Profile does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Profile) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "habbitchain.profile.Profile.username":
		x.Username = value.Interface().(string)
	case "habbitchain.profile.Profile.pubKey":
		x.PubKey = value.Interface().(string)
	case "habbitchain.profile.Profile.telegramId":
		x.TelegramId = int32(value.Int())
	case "habbitchain.profile.Profile.cosmosAddress":
		x.CosmosAddress = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: habbitchain.profile.Profile"))
		}
		panic(fmt.Errorf("message habbitchain.profile.Profile does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Profile) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "habbitchain.profile.Profile.username":
		panic(fmt.Errorf("field username of message habbitchain.profile.Profile is not mutable"))
	case "habbitchain.profile.Profile.pubKey":
		panic(fmt.Errorf("field pubKey of message habbitchain.profile.Profile is not mutable"))
	case "habbitchain.profile.Profile.telegramId":
		panic(fmt.Errorf("field telegramId of message habbitchain.profile.Profile is not mutable"))
	case "habbitchain.profile.Profile.cosmosAddress":
		panic(fmt.Errorf("field cosmosAddress of message habbitchain.profile.Profile is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: habbitchain.profile.Profile"))
		}
		panic(fmt.Errorf("message habbitchain.profile.Profile does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Profile) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "habbitchain.profile.Profile.username":
		return protoreflect.ValueOfString("")
	case "habbitchain.profile.Profile.pubKey":
		return protoreflect.ValueOfString("")
	case "habbitchain.profile.Profile.telegramId":
		return protoreflect.ValueOfInt32(int32(0))
	case "habbitchain.profile.Profile.cosmosAddress":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: habbitchain.profile.Profile"))
		}
		panic(fmt.Errorf("message habbitchain.profile.Profile does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Profile) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in habbitchain.profile.Profile", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Profile) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Profile) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Profile) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Profile) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Profile)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Username)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.PubKey)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.TelegramId != 0 {
			n += 1 + runtime.Sov(uint64(x.TelegramId))
		}
		l = len(x.CosmosAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Profile)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.CosmosAddress) > 0 {
			i -= len(x.CosmosAddress)
			copy(dAtA[i:], x.CosmosAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.CosmosAddress)))
			i--
			dAtA[i] = 0x22
		}
		if x.TelegramId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TelegramId))
			i--
			dAtA[i] = 0x18
		}
		if len(x.PubKey) > 0 {
			i -= len(x.PubKey)
			copy(dAtA[i:], x.PubKey)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.PubKey)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Username) > 0 {
			i -= len(x.Username)
			copy(dAtA[i:], x.Username)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Username)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Profile)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Profile: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Profile: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Username = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.PubKey = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TelegramId", wireType)
				}
				x.TelegramId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.TelegramId |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CosmosAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.CosmosAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: habbitchain/profile/profile.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username      string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	PubKey        string `protobuf:"bytes,2,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	TelegramId    int32  `protobuf:"varint,3,opt,name=telegramId,proto3" json:"telegramId,omitempty"`
	CosmosAddress string `protobuf:"bytes,4,opt,name=cosmosAddress,proto3" json:"cosmosAddress,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_habbitchain_profile_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_habbitchain_profile_profile_proto_rawDescGZIP(), []int{0}
}

func (x *Profile) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Profile) GetPubKey() string {
	if x != nil {
		return x.PubKey
	}
	return ""
}

func (x *Profile) GetTelegramId() int32 {
	if x != nil {
		return x.TelegramId
	}
	return 0
}

func (x *Profile) GetCosmosAddress() string {
	if x != nil {
		return x.CosmosAddress
	}
	return ""
}

var File_habbitchain_profile_profile_proto protoreflect.FileDescriptor

var file_habbitchain_profile_profile_proto_rawDesc = []byte{
	0x0a, 0x21, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x6c, 0x65,
	0x67, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65,
	0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0xb9,
	0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x62, 0x62, 0x69, 0x74, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x0c, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x23, 0x68, 0x61, 0x62, 0x62,
	0x69, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x61, 0x62, 0x62,
	0x69, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0xa2,
	0x02, 0x03, 0x48, 0x50, 0x58, 0xaa, 0x02, 0x13, 0x48, 0x61, 0x62, 0x62, 0x69, 0x74, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0xca, 0x02, 0x13, 0x48, 0x61,
	0x62, 0x62, 0x69, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0xe2, 0x02, 0x1f, 0x48, 0x61, 0x62, 0x62, 0x69, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5c,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x48, 0x61, 0x62, 0x62, 0x69, 0x74, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_habbitchain_profile_profile_proto_rawDescOnce sync.Once
	file_habbitchain_profile_profile_proto_rawDescData = file_habbitchain_profile_profile_proto_rawDesc
)

func file_habbitchain_profile_profile_proto_rawDescGZIP() []byte {
	file_habbitchain_profile_profile_proto_rawDescOnce.Do(func() {
		file_habbitchain_profile_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_habbitchain_profile_profile_proto_rawDescData)
	})
	return file_habbitchain_profile_profile_proto_rawDescData
}

var file_habbitchain_profile_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_habbitchain_profile_profile_proto_goTypes = []interface{}{
	(*Profile)(nil), // 0: habbitchain.profile.Profile
}
var file_habbitchain_profile_profile_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_habbitchain_profile_profile_proto_init() }
func file_habbitchain_profile_profile_proto_init() {
	if File_habbitchain_profile_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_habbitchain_profile_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_habbitchain_profile_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_habbitchain_profile_profile_proto_goTypes,
		DependencyIndexes: file_habbitchain_profile_profile_proto_depIdxs,
		MessageInfos:      file_habbitchain_profile_profile_proto_msgTypes,
	}.Build()
	File_habbitchain_profile_profile_proto = out.File
	file_habbitchain_profile_profile_proto_rawDesc = nil
	file_habbitchain_profile_profile_proto_goTypes = nil
	file_habbitchain_profile_profile_proto_depIdxs = nil
}
