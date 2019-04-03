package nulls

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"strconv"
)

// Int64 replaces sql.Int64 with an implementation
// that supports proper JSON encoding/decoding.
type Int64 sql.NullInt64

// Interface implements the nullable interface. It returns nil if
// the int64 is not valid, otherwise it returns the int64 value.
func (ns Int64) Interface() interface{} {
	if !ns.Valid {
		return nil
	}
	return ns.Int64
}

// NewInt64 returns a new, properly instantiated
// Int64 object.
func NewInt64(i int64) Int64 {
	return Int64{Int64: i, Valid: true}
}

// Scan implements the Scanner interface.
func (ns *Int64) Scan(value interface{}) error {
	n := sql.NullInt64{Int64: ns.Int64}
	err := n.Scan(value)
	ns.Int64, ns.Valid = n.Int64, n.Valid
	return err
}

// Value implements the driver Valuer interface.
func (ns Int64) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.Int64, nil
}

// MarshalJSON marshals the underlying value to a
// proper JSON representation.
func (ns Int64) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Int64)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON will unmarshal a JSON value into
// the propert representation of that value.
func (ns *Int64) UnmarshalJSON(text []byte) error {
	t := string(text)
	ns.Valid = true
	if t == "null" {
		ns.Valid = false
		return nil
	}
	i, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		ns.Valid = false
		return err
	}
	ns.Int64 = i
	return nil
}

// UnmarshalText will unmarshal text value into
// the propert representation of that value.
func (ns *Int64) UnmarshalText(text []byte) error {
	return ns.UnmarshalJSON(text)
}

func (ns Int64) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if ns.Valid {
		return e.EncodeElement(ns.Int64, start)
	}
	return nil
}

func (ns *Int64) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

	val, err := strconv.ParseInt(data, 10, 64)

	if err != nil {
		return err
	}

	ns.Valid = true
	ns.Int64 = val

	return nil
}

func (ns Int64) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if ns.Valid {
		var value string
		value = strconv.FormatInt(int64(ns.Int64), 10)

		return xml.Attr{
			Name:  name,
			Value: value,
		}, nil
	}
	return xml.Attr{}, nil
}

func (ns *Int64) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "" {
		return nil
	}

	if attr.Value == "null" {
		return nil
	}

	val, err := strconv.ParseInt(attr.Value, 10, 64)

	if err != nil {
		return err
	}

	ns.Valid = true
	ns.Int64 = val

	return nil
}
