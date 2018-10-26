package nulls

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
)

// String replaces sql.NullString with an implementation
// that supports proper JSON encoding/decoding.
type String sql.NullString

// Interface implements the nullable interface. It returns nil if
// the string is not valid, otherwise it returns the string value.
func (ns String) Interface() interface{} {
	if !ns.Valid {
		return nil
	}
	return ns.String
}

// NewString returns a new, properly instantiated
// String object.
func NewString(s string) String {
	return String{String: s, Valid: true}
}

// Scan implements the Scanner interface.
func (ns *String) Scan(value interface{}) error {
	n := sql.NullString{String: ns.String}
	err := n.Scan(value)
	ns.String, ns.Valid = n.String, n.Valid
	return err
}

// Value implements the driver Valuer interface.
func (ns String) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.String, nil
}

// MarshalJSON marshals the underlying value to a
// proper JSON representation.
func (ns String) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON will unmarshal a JSON value into
// the propert representation of that value.
func (ns *String) UnmarshalJSON(text []byte) error {
	ns.Valid = false
	if string(text) == "null" {
		return nil
	}
	if err := json.Unmarshal(text, &ns.String); err == nil {
		ns.Valid = true
	}
	return nil
}

// UnmarshalText will unmarshal text value into
// the propert representation of that value.
func (ns *String) UnmarshalText(text []byte) error {
	ns.Valid = false
	t := string(text)
	if t == "null" {
		return nil
	}
	ns.String = t
	ns.Valid = true
	return nil
}

func (ns String) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if ns.Valid {
		return e.EncodeElement(ns.String, start)
	}
	return nil
}

func (ns *String) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

	ns.Valid = true
	ns.String = data

	return nil
}

func (ns String) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if ns.Valid {
		return xml.Attr{
			Name:  name,
			Value: ns.String,
		}, nil
	}
	return xml.Attr{}, nil
}

func (ns *String) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "" {
		return nil
	}

	if attr.Value == "null" {
		return nil
	}

	ns.Valid = true
	ns.String = attr.Value

	return nil
}