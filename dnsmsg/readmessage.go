package dnsmsg

import (
	"encoding/binary"
	"errors"
)

//char of dots. ascii hex
const dot byte = 0x2e

func GetQuestionName(message []byte) (string, error) {
	const headerByteLen int = 12

	var readCounter int = 0 //現在メッセージの何バイト目まで読み込んだかを示すカウンター

	if len(message) < headerByteLen {
		return "", errors.New("No DNS response data.")
	}
	header := message[:headerByteLen]
	readCounter = headerByteLen

	var qCount uint16 = binary.BigEndian.Uint16(header[4:6])

	if qCount != 1 {
		return "", errors.New("Question count needs 1")
	}

	// ---- read Question section ----
	name, _, err := readName(message, readCounter)
	if err != nil {
		return "", err
	}

	return name, nil
}

// read domain name
// name ex. vaddy.net
// readByte means read byte size
func readName(message []byte, readCounter int) (name string, readByte int, err error) {
	const nullByteLen int = 1

	data := message[readCounter:]
	var labelCount uint8 = 0
	var nameByte []byte = make([]byte, 0, 50)
	for readByte, byteData := range data {
		if byteData == 0x00 {
			nameByte = dropDotInHeadByte(nameByte)
			name = string(nameByte)
			if nameByte[len(nameByte)-1] != dot {
				//最後がドットで終わってなければドットを付与する
				name = name + string(dot)
			}
			return name, readByte + nullByteLen, nil
		}
		if labelCount == 0 { //label count0の場合はラベルの数字のため、圧縮されていないか確認
			if needCheckCompression(byteData) {
				compressedCounter := getFragmentPointer(byteData, data[readByte+1])  //現在の読込バイトと、次の読込予定バイトを使って圧縮先のバイト数を取得
				compressedNameString, _, err := readName(message, compressedCounter) //圧縮先のデータを再帰で取得。圧縮先で読んだバイト数は不要
				if err != nil {
					return "", 0, err
				}
				nameByte = dropDotInHeadByte(nameByte)
				name = string(nameByte) + string(dot) + compressedNameString
				readByte += 2 //label文字数を示す1バイトと、フラグメントのポインタを示す1バイトを読込済みとして足しておく
				return name, readByte, nil
			}

			//labelCount=0はラベル文字数を読み取り、ドットの文字を連結
			nameByte = append(nameByte, dot)
			labelCount = byteData
		} else {
			nameByte = append(nameByte, byteData)
			labelCount--
		}
	}
	return "", 0, errors.New("No terminate byte.")
}

//圧縮先の参照データの先頭からのバイト数を取得
func getFragmentPointer(first byte, second byte) int {
	//  upperの上位2ビットを落とす
	//  underの8bitとupperを8shiftしたものを足したint16の数
	var upper int16 = int16(first & 0b00111111)
	var under int16 = int16(second)
	return int(upper<<8 + under)
}

//先頭ビットが11で始まる0b11000000(63)以上の場合は名前圧縮されているためその判定を行う関数
func needCheckCompression(byteData byte) bool {
	return byteData > 63
}

//名前配列の先頭がドットだった場合は先頭のドットを切り捨てる
func dropDotInHeadByte(nameByte []byte) []byte {
	if nameByte[0] == dot {
		//先頭のドットは不要
		nameByte = nameByte[1:]
	}
	return nameByte
}