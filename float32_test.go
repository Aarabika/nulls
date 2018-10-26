package nulls

import (
	"encoding/xml"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFloat32Valid_MarshalXML(t *testing.T) {
	type test struct {
		Val Float32 `xml:"val"`
	}

	val := test{
		Val: Float32{
			Float32: 3.22,
			Valid:   true,
		},
	}

	data, err := xml.Marshal(val)
	assert.NoError(t, err)

	assert.Equal(t, "<test><val>3.22</val></test>", string(data))
}

func TestFloat32Invalid_MarshalXML(t *testing.T) {
	type test struct {
		Val Float32 `xml:"val"`
	}

	val := test{
		Val: Float32{},
	}

	data, err := xml.Marshal(val)
	assert.NoError(t, err)

	assert.Equal(t, "<test></test>", string(data))
}

func TestFloat32Invalid_UnmarshalXML(t *testing.T) {
	type test struct {
		Val Float32 `xml:"val"`
	}

	var val test

	err := xml.Unmarshal([]byte(xml.Header+`<test></test>`), &val)
	assert.NoError(t, err)

	assert.Equal(t, false, val.Val.Valid)
}

func TestFloat32Valid_UnmarshalXML(t *testing.T) {
	type test struct {
		Val Float32 `xml:"val"`
	}

	var val test

	err := xml.Unmarshal([]byte(xml.Header+`<test><val>3.22</wval></test>`), &val)
	assert.NoError(t, err)

	assert.Equal(t, true, val.Val.Valid)
	assert.Equal(t, float32(3.22), val.Val.Float32)
}

func TestFloat32Valid_MarshalXMLAttr(t *testing.T) {
	type test struct {
		Val Float32 `xml:"val,attr"`
	}

	val := test{
		Val: Float32{
			Float32: 3.22,
			Valid:   true,
		},
	}

	body, err := xml.Marshal(&val)
	assert.NoError(t, err)

	assert.Equal(t, "<test val=\"3.22\"></test>", string(body))
}

func TestFloat32Invalid_MarshalXMLAttr(t *testing.T) {
	type test struct {
		Val Float32 `xml:"val,attr"`
	}

	val := test{
		Val: Float32{},
	}

	body, err := xml.Marshal(&val)
	assert.NoError(t, err)

	assert.Equal(t, "<test></test>", string(body))
}

func TestFloat32Invalid_UnmarshalXMLAttr(t *testing.T) {
	type test struct {
		Val Float32 `xml:"val,attr"`
	}

	var val test

	err := xml.Unmarshal([]byte(xml.Header+`<test></test>`), &val)
	assert.NoError(t, err)

	assert.Equal(t, false, val.Val.Valid)
}

func TestFloat32Valid_UnmarshalXMLAttr(t *testing.T) {
	type test struct {
		Val Float32 `xml:"val,attr"`
	}

	var val test

	err := xml.Unmarshal([]byte(xml.Header+`<test val="3.22"></test>`), &val)
	assert.NoError(t, err)

	assert.Equal(t, true, val.Val.Valid)
	assert.Equal(t, float32(3.22), val.Val.Float32)
}
