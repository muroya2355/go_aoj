package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

// crypto/x509 ライブラリを利用した証明書の作成
func main() {

	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		Organization:       []string{"Mannig Publication Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}

	template := x509.Certificate{
		//シリアルナンバー（ランダム）
		SerialNumber: serialNumber,
		//識別名
		Subject: subject,
		//有効期間（開始日）
		NotBefore: time.Now(),
		//有効期間（終了日）
		NotAfter: time.Now().Add(365 * 24 * time.Hour),
		//この X509 証明書がサーバ認証に使用されることを示す
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		//証明書は 127.0.0.1 でのみ効力を持つ
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},
	}

	// RSA秘密鍵を生成
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	// 構造体Certificate と公開鍵、秘密鍵等を引数にして DER 形式のバイトデータスライスを作成
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	// 証明書データを符号化、ファイル cert.pem を作成
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	// pk を符号化して ファイル key.pem に保存
	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()

}
