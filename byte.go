package interim

import (
	"bytes"
	"encoding/gob"
)

func encodeToByte(data interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	err := gob.NewEncoder(&buffer).Encode(data)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func decodeFromByte(encoded []byte, target interface{}) error {
	buffer := bytes.NewBuffer(encoded)
	err := gob.NewDecoder(buffer).Decode(target)
	if err != nil {
		return err
	}
	return nil
}
