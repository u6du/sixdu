package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
	"github.com/u6du/ex"
	"github.com/u6du/go-rfc1924/base85"
	"golang.org/x/crypto/ed25519"

	"sixdu/key"
)

/*
func TestDns(dnsLi []string) {

	wait := sync.WaitGroup{}
	wait.Add(len(dnsLi))

	for _, dns := range dnsLi {
		go func(dns string) {
			n := 0
			defer wait.Done()
			for {
				resolve := net.NewResolver(dns)
				li, err := resolve.LookupTXT(context.Background(), "txt.6du.host")
				if err != nil {
					log.Print(dns, " ", err)
					if n > 3 {
						return
					}
				}
				for _, i := range li {
					fmt.Printf("%s %s\n", dns, i)
					return
				}
				n += 1
			}
		}(dns)
	}

	wait.Wait()

}
*/

func main() {
	signStr := "7xW;Y8v!}fq_bBRv7AV`HY>E(YE(iMX_hs|Pm@VQ-l9nRO@IMX`;Ee)_rfDhpG3_KAUtn69a`$%Fw*e`"
	sign, err := base85.DecodeString(signStr)
	ex.Panic(err)
	verify := ed25519.Verify(key.GodPublic, []byte("hello"), sign)
	log.Info().Bool("verify", verify).Msg("")
	//	TestDns(net.DnsDot)

	/*
		for i := 1; i <= 10; i++ {
			bootNode := db.BootNode("txt.6du.host")
			if len(bootNode) == 0 {
				panic(errors.New("boot node not found"))
			} else {
				log.Info().Msg(bootNode)
			}
		}
	*/
	return
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

			home, err := os.UserHomeDir()
			ex.Panic(err)

			home = path.Join(home, ".config", "6du")

			err = os.MkdirAll(home, os.ModePerm)
			ex.Panic(err)

			db, err := sql.Open("sqlite3", path.Join(home, "6du.db"))
			ex.Panic(err)
			defer db.Close()

			println(db)
			//	token := make([]byte, 32)

			//	n, err := cryptoRand.Read(token)
			//	println("cryptoRand.Read", n)
			//	ex.Panic(err)


			target := "6du-boot.6du.world"

			server := "8.8.8.8"

			c := dns.Client{}
			m := dns.Msg{}
			m.SetQuestion(target+".", dns.TypeTXT)
			r, t, err := c.Exchange(&m, server+":53")
			ex.Panic(err)

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
