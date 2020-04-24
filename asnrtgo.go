package asnrtgo

// Encoding Rules
const (
	BASIC_ENCODING_RULES            = 0
	CANONICAL_ENCODING_RULES        = 1
	DISTINGUISHED_ENCODING_RULES    = 2
	UNALIGNED_PACKED_ENCODING_RULES = 3
	ALIGNED_PACKED_ENCODING_RULES   = 4
)

type Dynamic interface{}

type Null struct {
}

type Enum interface {
	Name() string
	Ordinal() int
}

type BitString struct {
	Bytes      []byte // bits packed into bytes.
	UnusedBits uint8  // unused bits in tailing byte.
}

type AsnType interface {
	Tag() uint
	MatchTag(tag uint) bool
}

type Buffer interface {
	EncodingRules() byte
}

type OctetString []byte

type ObjectIdentifier []int


type AsnrtAPI struct {
	Length         func(self BitString) uint
	Bit            func(self BitString, index uint) bool
	SetBit         func(self BitString, index uint, b bool)
	AllocateBuffer func(length int, encodingRules byte) (Buffer, error)
	Encode         func(b Buffer, v Dynamic, t AsnType) error
	Decode         func(b Buffer, v Dynamic, t AsnType) error
	DecodeMetadata func(metadata []byte) map[int]AsnType
}
