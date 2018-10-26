package nulls

import (
	"testing"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
)

func TestBoolValid_MarshalXML(t *testing.T) {
	type test struct {
		Val Bool `xml:"val"`
	}

	val := test{
		Val: Bool{
			Bool:  true,
			Valid: true,
		},
	}

	data, err := xml.Marshal(val)
	assert.NoError(t, err)

	assert.Equal(t, "<test><val>true</val></test>", string(data))
}

func TestBoolInvalid_MarshalXML(t *testing.T) {
	type test struct {
		Val Bool `xml:"val"`
	}

	val := test{
		Val: Bool{},
	}

	data, err := xml.Marshal(val)
	assert.NoError(t, err)

	assert.Equal(t, "<test></test>", string(data))
}

func TestBoolValid_MarshalXMLAttr(t *testing.T) {
	type test struct {
		Val Bool `xml:"val,attr"`
	}

	val := test{
		Val: Bool{
			Bool:  true,
			Valid: true,
		},
	}

	body, err := xml.Marshal(&val)
	assert.NoError(t, err)

	assert.Equal(t, "<test val=\"true\"></test>", string(body))
}


func TestBoolInvalid_MarshalXMLAttr(t *testing.T) {
	type test struct {
		Val Bool `xml:"val,attr"`
	}

	val := test{
		Val: Bool{},
	}

	body, err := xml.Marshal(&val)
	assert.NoError(t, err)

	assert.Equal(t, "<test></test>", string(body))
}