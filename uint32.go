package nulls

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"strconv"
)

// UInt32 adds an implementation for int
// that supports proper JSON encoding/decoding.
type UInt32 struct {
	UInt32 uint32
	Valid  bool // Valid is true if Int is not NULL
}

// Interface implements the nullable interface. It returns nil if
// the uint32 is not valid, otherwise it returns the uint32 value.
func (ns UInt32) Interface() interface{} {
	if !ns.Valid {
		return nil
	}
	return ns.UInt32
}

// NewUInt32 returns a new, properly instantiated
// Int object.
func NewUInt32(i uint32) UInt32 {
	return UInt32{UInt32: i, Valid: true}
}

// Scan implements the Scanner interface.
func (ns *UInt32) Scan(value interface{}) error {
	n := sql.NullInt64{Int64: int64(ns.UInt32)}
	err := n.Scan(value)
	ns.UInt32, ns.Valid = uint32(n.Int64), n.Valid
	return err
}

// Value implements the driver Valuer interface.
func (ns UInt32) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return int64(ns.UInt32), nil
}

// MarshalJSON marshals the underlying value to a
// proper JSON representation.
func (ns UInt32) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.UInt32)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON will unmarshal a JSON value into
// the propert representation of that value.
func (ns *UInt32) UnmarshalJSON(text []byte) error {
	txt := string(text)
	ns.Valid = true
	if txt == "null" {
		ns.Valid = false
		return nil
	}
	i, err := strconv.ParseInt(txt, 10, 32)
	if err != nil {
		ns.Valid = false
		return err
	}
	j := uint32(i)
	ns.UInt32 = j
	return nil
}

// UnmarshalText will unmarshal text value into
// the propert representation of that value.
func (ns *UInt32) UnmarshalText(text []byte) error {
	return ns.UnmarshalJSON(text)
}

func (ns UInt32) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if ns.Valid {
		return e.EncodeElement(ns.UInt32, start)
	}
	return nil
}

func (ns *UInt32) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var data string
	err := d.DecodeElement(&data, &start)

	if err != nil {
		return err
	}
	if data == "" {
		return nil
	}

	if data == "null" {
		return nil
	}

	val, err := strconv.ParseInt(data, 10, strconv.IntSize)

	if err != nil {
		return err
	}

	ns.Valid = true
	ns.UInt32 = uint32(val)

	return nil
}

func (ns UInt32) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if ns.Valid {
		var value string
		value = strconv.FormatInt(int64(ns.UInt32), 10)

		return xml.Attr{
			Name:  name,
			Value: value,
		}, nil
	}
	return xml.Attr{}, nil
}

func (ns *UInt32) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "" {
		return nil
	}

	if attr.Value == "null" {
		return nil
	}

	val, err := strconv.ParseInt(attr.Value, 10, strconv.IntSize)

	if err != nil {
		return err
	}

	ns.Valid = true
	ns.UInt32 = uint32(val)

	return nil
}
