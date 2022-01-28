package gss

import (
	"encoding/asn1"

	"github.com/blackhat-go/bhg/ch-6/smb/smb/encoder"
)

func (f *Foo) MarshalBinary(meta *encoder.Metadata) ([]byte, error) {
	buff, err := asn1.Marshal(*f)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (f *Foo) UnmarshalBinary(buf []byte, meta *encoder.Metadata) error {
	data := Foo{}
	if _, err := asn1.Unmarshal(buf, &data); err != nil {
		return err
	}
	*f = data
	return nil
}
