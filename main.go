package main

import (
	"net"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/u6du/ex"
)

func serve() {

	conn, err := net.ListenPacket("udp", ":49101")
	ex.Panic(err)
	defer conn.Close()

	for {
		buf := make([]byte, 65535)
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			continue
		}
		println("recv", buf[:n])
		println("addr", addr.String())
	}
}

func client() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 0})
	ex.Panic(err)
	conn.WriteTo([]byte("test"), &net.UDPAddr{IP: net.ParseIP("54.193.97.48"), Port: 49101})
}

func main() {
	go serve()

	n := 0
	for {
		println(">", n)
		time.Sleep(10 * time.Second)
		n += 1
	}
	/*
		secret, err := bls.RandKey(rand.Reader)

		secretByte := secret.Marshal()
		b64token := make([]byte, ascii85.MaxEncodedLen(len(secretByte)))

		ascii85.Encode(b64token, secretByte)
		token := string(b64token)
		println("secret", token, len(token))

		//	secret, err := bls.SecretKeyFromBytes(token)
		//	ex.Panic(err)
		txt := []byte(`流的前两个字节包含主要版本和次要版本，每个版本都作为编码数字的单个字节。在 Ruby 中实现的版本是4.8（存储为“x04x08”），并受到 Ruby 1.8.0和更高版本的支持。

			统一格式的不同主要版本不兼容，不能被其他主要版本理解。更新的次要版本可以理解较小版本的格式。格式4.7可以通过4.8实现加载，但格式4.8不能通过4.7实现加载。

			版本字节后面是描述序列化对象的流。该流包含嵌套对象（与 Ruby 对象相同），但流中的对象不一定有直接映射到 Ruby 对象模型。

			流中的每个对象都由一个字节来描述它的类型，后面跟着一个或多个描述该对象的字节。当下面提到“对象”时，它意味着下面定义 Ruby 对象的任何类型。

			true, false, nil`)

		sign := secret.Sign(txt, SignTypeDns)

		println("sign.Marshal len", len(sign.Marshal()))
		signByte := sign.Marshal()

		b64sign := base64.RawURLEncoding.EncodeToString(sign.Marshal())

		b86token := make([]byte, ascii85.MaxEncodedLen(len(signByte)))
		ascii85.Encode(b86token, signByte)
		b85 := string(b86token)
		println("signByte len", len(signByte), len(b64sign), len(b85))
		println(b85)

		public := secret.PublicKey()
		println("public", len(public.Marshal()), base64.RawURLEncoding.EncodeToString(public.Marshal()))

		signBytes, err := base64.RawURLEncoding.DecodeString(b64sign)
		sign, err = bls.SignatureFromBytes(signBytes)
		println("sign.Verify", sign.Verify(txt, public, SignTypeDns))

	*/
}
