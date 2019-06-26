package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ameshkov/dnscrypt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/miekg/dns"
	. "github.com/urwork/throw"
)

const SignTypeDns = 2

func main() {

	target := "6du-boot.6du.space"

	//stampStr := "sdns://AQIAAAAAAAAAFDE3Ni4xMDMuMTMwLjEzMDo1NDQzINErR_JS3PLCu_iZEIbq95zkSV2LFsigxDIuUso_OQhzIjIuZG5zY3J5cHQuZGVmYXVsdC5uczEuYWRndWFyZC5jb20"
	stampStr := "sdns://AQYAAAAAAAAADTkuOS45LjEwOjg0NDMgZ8hHuMh1jNEgJFVDvnVnRt803x2EwAuMRwNo34Idhj4ZMi5kbnNjcnlwdC1jZXJ0LnF1YWQ5Lm5ldA"
	// Initializing the DNSCrypt client
	client := dnscrypt.Client{Proto: "udp", Timeout: 10 * time.Second}

	// Fetching and validating the server certificate
	serverInfo, _, err := client.Dial(stampStr)
	Throw(err)
	println(serverInfo.ProviderName)
	println(serverInfo.ServerAddress)
	fmt.Printf("%x\n", serverInfo.PublicKey)

	// Create a DNS request
	req := dns.Msg{}
	req.Id = dns.Id()
	req.RecursionDesired = true
	req.Question = []dns.Question{
		{Name: target + ".", Qtype: dns.TypeTXT, Qclass: dns.ClassINET},
	}

	// Get the DNS response
	reply, _, err := client.Exchange(&req, serverInfo)
	Throw(err)
	for _, ans := range reply.Answer {
		record := ans.(*dns.TXT)
		println(record.String())
		log.Printf("%s", record.Txt)
	}
	println("end")
	/*
			home, err := os.UserHomeDir()
			Throw(err)

			home = path.Join(home, ".config", "6du")

			err = os.MkdirAll(home, os.ModePerm)
			Throw(err)

			db, err := sql.Open("sqlite3", path.Join(home, "6du.db"))
			Throw(err)
			defer db.Close()

			println(db)
			//	token := make([]byte, 32)

			//	n, err := cryptoRand.Read(token)
			//	println("cryptoRand.Read", n)
			//	Throw(err)
			secret, err := bls.RandKey(cryptoRand.Reader)

			b64token := base64.RawURLEncoding.EncodeToString(secret.Marshal())
			println("secret", b64token)

			//	secret, err := bls.SecretKeyFromBytes(token)
			//	Throw(err)
			txt := []byte(`流的前两个字节包含主要版本和次要版本，每个版本都作为编码数字的单个字节。在 Ruby 中实现的版本是4.8（存储为“x04x08”），并受到 Ruby 1.8.0和更高版本的支持。

		统一格式的不同主要版本不兼容，不能被其他主要版本理解。更新的次要版本可以理解较小版本的格式。格式4.7可以通过4.8实现加载，但格式4.8不能通过4.7实现加载。

		版本字节后面是描述序列化对象的流。该流包含嵌套对象（与 Ruby 对象相同），但流中的对象不一定有直接映射到 Ruby 对象模型。

		流中的每个对象都由一个字节来描述它的类型，后面跟着一个或多个描述该对象的字节。当下面提到“对象”时，它意味着下面定义 Ruby 对象的任何类型。

		true, false, nil`)

			sign := secret.Sign(txt, Sign_type_dns)

			println("sign.Marshal len", len(sign.Marshal()))

			b64sign := base64.RawURLEncoding.EncodeToString(sign.Marshal())
			println(b64sign, len(b64sign))

			public := secret.PublicKey()
			println("public", base64.RawURLEncoding.EncodeToString(public.Marshal()))

			signBytes, err := base64.RawURLEncoding.DecodeString(b64sign)
			sign, err = bls.SignatureFromBytes(signBytes)
			println("sign.Verify", sign.Verify(txt, public, Sign_type_dns))

			target := "6du-boot.6du.world"

			server := "8.8.8.8"

			c := dns.Client{}
			m := dns.Msg{}
			m.SetQuestion(target+".", dns.TypeTXT)
			r, t, err := c.Exchange(&m, server+":53")
			Throw(err)

			log.Printf("Took %v", t)
			if len(r.Answer) == 0 {
				log.Fatal("No results")
			}
			for _, ans := range r.Answer {
				record := ans.(*dns.TXT)
				log.Printf("%s", record.Txt)
			}
	*/
}
