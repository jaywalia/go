// 54 base 64

package main

import b64 "encoding/base64"
import "fmt"

func main() {
	p := fmt.Println

	data := "abc123!?$*&()'-=@~"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	p(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	p(string(sDec))
	p()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	p(uEnc)

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	p(string(uDec))
}
