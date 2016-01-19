package codec

import (
	"encoding/binary"
	"encoding/json"
	log "github.com/cihub/seelog"
	"github.com/lrsec/go-util/aes"
	"linbox_go/messages"
)

const (
	protocol_version uint16 = 1

	default_password = "medtree-im-passw"
)

type MsgCodec struct {
	aesCodec *aes.AESCodec
}

func NewMsgCodec() (*MsgCodec, error) {
	aesCodec, err := aes.NewAESCodec(default_password)
	if err != nil {
		return nil, err
	}

	codec := new(MsgCodec)
	codec.aesCodec = aesCodec

	return codec, nil
}

func (codec *MsgCodec) Encode(rrType messages.RequestResponseType, content interface{}) ([]byte, error) {
	contentRaw, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	encrypted := codec.aesCodec.Encrypt(contentRaw, codec.aesCodec.Password[:16])
	var length uint32 = uint32(len(encrypted))

	result := make([]byte, length+2+2+4)

	binary.BigEndian.PutUint16(result[0:2], protocol_version)
	binary.BigEndian.PutUint16(result[2:4], uint16(rrType))
	binary.BigEndian.PutUint32(result[4:8], length)
	copy(result[8:], encrypted)

	return result, nil
}

func (codec *MsgCodec) Decode(content []byte, result interface{}) error {
	decrypted := codec.aesCodec.Decrypt(content, codec.aesCodec.Password[:16])

	log.Debug("Decrypted: ", string(decrypted))

	err := json.Unmarshal(decrypted, result)

	return err
}
