package nulls

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"encoding/xml"
)

// Int32 adds an implementation for int32
// that supports proper JSON encoding/decoding.
type Int32 struct {
	Int32 int32
	Valid bool // Valid is true if Int32 is not NULL
}

// Interface implements the nullable interface. It returns nil if
// the int32 is not valid, otherwise it returns the int32 value.
func (ns Int32) Interface() interface{} {
	if !ns.Valid {
		return nil
	}
	return ns.Int32
}

// NewInt32 returns a new, properly instantiated
// Int object.
func NewInt32(i int32) Int32 {
	return Int32{Int32: i, Valid: true}
}

// Scan implements the Scanner interface.
func (ns *Int32) Scan(value interface{}) error {
	n := sql.NullInt64{Int64: int64(ns.Int32)}
	err := n.Scan(value)
	ns.Int32, ns.Valid = int32(n.Int64), n.Valid
	return err
}

// Value implements the driver Valuer interface.
func (ns Int32) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return int64(ns.Int32), nil
}

// MarshalJSON marshals the underlying value to a
// proper JSON representation.
func (ns Int32) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Int32)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON will unmarshal a JSON value into
// the propert representation of that value.
func (ns *Int32) UnmarshalJSON(text []byte) error {
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
	j := int32(i)
	ns.Int32 = j
	return nil
}

// UnmarshalText will unmarshal text value into
// the propert representation of that value.
func (ns *Int32) UnmarshalText(text []byte) error {
	return ns.UnmarshalJSON(text)
}

func (ns Int32) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if ns.Valid {
		return e.EncodeElement(ns.Int32, start)
	}
	return nil
}

func (ns *Int32) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
	ns.Int32 = int32(val)

	return nil
}

func (ns Int32) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if ns.Valid {
		var value string
		value = strconv.FormatInt(int64(ns.Int32), strconv.IntSize)

		return xml.Attr{
			Name:  name,
			Value: value,
		}, nil
	}
	return xml.Attr{}, nil
}

func (ns *Int32) UnmarshalXMLAttr(attr xml.Attr) error {
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
	ns.Int32 = int32(val)

	return nil
}
