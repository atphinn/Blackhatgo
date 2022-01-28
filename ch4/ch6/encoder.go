package encoder

type BinaryMarshallable interface {
	MarshalBinary(*Metadata) ([]byte, error)
	UnmarshalBinary([]byte, *Metadata) error
}

func marshal(v interface{}, meta *Metadata) ([]byte, error) {
	bm, ok := v.(BinaryMarshallable)
	if ok {
		//Custom marshallable interface found
		buf, err := bm.MarshalBinary(meta)
		if err != nil {
			return nil, err
		}
		return buf, err
	}
}

func unmarshal(buf []byte, v interface{}, meta *Metadata) (interface{}, error) {

	bm, ok := v.(BinaryMarshallable)
	if ok {
		//Custom marshallable interface found
		if err := bm.UnmarshalBinary(buf, meta); err != nil {
			return nil, err
		}
		return bm, nil
	}
}

func parseTags(sf reflect) {

}
