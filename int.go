package nulls

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"encoding/xml"
)

// Int adds an implementation for int
// that supports proper JSON encoding/decoding.
type Int struct {
	Int   int
	Valid bool // Valid is true if Int is not NULL
}

// Interface implements the nullable interface. It returns nil if
// the int is not valid, otherwise it returns the int value.
func (ns Int) Interface() interface{} {
	if !ns.Valid {
		return nil
	}
	return ns.Int
}

// NewInt returns a new, properly instantiated
// Int object.
func NewInt(i int) Int {
	return Int{Int: i, Valid: true}
}

// Scan implements the Scanner interface.
func (ns *Int) Scan(value interface{}) error {
	n := sql.NullInt64{Int64: int64(ns.Int)}
	err := n.Scan(value)
	ns.Int, ns.Valid = int(n.Int64), n.Valid
	return err
}

// Value implements the driver Valuer interface.
func (ns Int) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return int64(ns.Int), nil
}

// MarshalJSON marshals the underlying value to a
// proper JSON representation.
func (ns Int) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Int)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON will unmarshal a JSON value into
// the propert representation of that value.
func (ns *Int) UnmarshalJSON(text []byte) error {
	if i, err := strconv.ParseInt(string(text), 10, strconv.IntSize); err == nil {
		ns.Valid = true
		ns.Int = int(i)
	}
	return nil
}

// UnmarshalText will unmarshal text value into
// the propert representation of that value.
func (ns *Int) UnmarshalText(text []byte) error {
	return ns.UnmarshalJSON(text)
}

func (ns Int) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if ns.Valid {
		return e.EncodeElement(ns.Int, start)
	}
	return nil
}

func (ns *Int) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
	ns.Int = int(val)

	return nil
}

func (ns Int) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if ns.Valid {
		var value string
		value = strconv.FormatInt(int64(ns.Int), strconv.IntSize)

		return xml.Attr{
			Name:  name,
			Value: value,
		}, nil
	}
	return xml.Attr{}, nil
}

func (ns *Int) UnmarshalXMLAttr(attr xml.Attr) error {
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
	ns.Int = int(val)

	return nil
}