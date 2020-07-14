package response

import "encoding/json"

// A CopyrightResponse represents a copyright json response
// Spotify copyright object.
type CopyrightResponse struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

// NewCopyrightResponse returns a new CopyrightResponse given
// a json byte array.
func NewCopyrightResponse(data []byte) (*CopyrightResponse, error) {
	var copyright = new(CopyrightResponse)
	err := json.Unmarshal(data, &copyright)
	if err != nil {
		return nil, err
	}
	return copyright, nil
}

// A Copyright represents a copyright for a Spotify resource.
type Copyright struct {
	text string
	copyrightType string
}

// NewCopyright returns a new Copyright pointer given a
// CopyrightResponse.
func NewCopyright(copyright *CopyrightResponse) *Copyright {
	return &Copyright{
		text: copyright.Text,
		copyrightType: copyright.Type,
	}
}

// Text returns the copyright text.
func (c Copyright) Text() string {
	return c.text
}

// Type returns the copyright type.
func (c Copyright) Type() string {
	return c.copyrightType
}

// IsTypeP returns true if the copyright is a P performance type
// copyright.
func (c Copyright) IsTypeP() bool {
	return c.copyrightType == "P"
}

// IsTypeC returns true if the copyright is a C copyright type
// copyright.
func (c Copyright) IsTypeC() bool {
	return c.copyrightType == "C"
}

// String returns the copyright text. String returns the same
// string as Text.
func (c Copyright) String() string {
	return c.Text()
}
