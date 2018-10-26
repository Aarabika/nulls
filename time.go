package nulls

import (
	"database/sql/driver"
	"encoding/json"
	"time"
	"encoding/xml"
)

// Time replaces sql.NullTime with an implementation
// that supports proper JSON encoding/decoding.
type Time struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Interface implements the nullable interface. It returns nil if
// the Time is not valid, otherwise it returns the Time value.
func (ns Time) Interface() interface{} {
	if !ns.Valid {
		return nil
	}
	return ns.Time
}

// NewTime returns a new, properly instantiated
// Time object.
func NewTime(t time.Time) Time {
	return Time{Time: t, Valid: true}
}

// Scan implements the Scanner interface.
func (ns *Time) Scan(value interface{}) error {
	ns.Time, ns.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (ns Time) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.Time, nil
}

// MarshalJSON marshals the underlying value to a
// proper JSON representation.
func (ns Time) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Time)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON will unmarshal a JSON value into
// the propert representation of that value.
func (ns *Time) UnmarshalJSON(text []byte) error {
	ns.Valid = false
	txt := string(text)
	if txt == "null" || txt == "" {
		return nil
	}

	t := time.Time{}
	err := t.UnmarshalJSON(text)
	if err == nil {
		ns.Time = t
		ns.Valid = true
	}

	return err
}

// UnmarshalText will unmarshal text value into
// the propert representation of that value.
func (ns *Time) UnmarshalText(text []byte) error {
	return ns.UnmarshalJSON(text)
}

func (ns Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if ns.Valid {
		return e.EncodeElement(ns.Time, start)
	}
	return nil
}

func (ns *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

	t := time.Time{}
	err = t.UnmarshalJSON([]byte(data))

	if err == nil {
		ns.Time = t
		ns.Valid = true
	}

	return nil
}

func (ns Time) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if ns.Valid {
		return xml.Attr{
			Name:  name,
			Value: ns.Time.String(),
		}, nil
	}
	return xml.Attr{}, nil
}

func (ns *Time) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "" {
		return nil
	}

	if attr.Value == "null" {
		return nil
	}

	t := time.Time{}
	err := t.UnmarshalJSON([]byte(attr.Value))

	if err == nil {
		ns.Time = t
		ns.Valid = true
	}

	return nil
}