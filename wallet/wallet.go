package wallet
//其中 `"crypto/ecdsa"` 包，就是椭圆曲线加密包。
import (
	"crypto/sha256"
	"github.com/mr-tron/base58"
	"golang.org/x/crypto/ripemd160"
	"log"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)
type Address = string
const KeyBitSize = 256
const version = byte(0x00)
const checkSumLen = 4
type Wallet struct {
	// 私钥为 *bip32.Key 类型
	privateKey *bip32.Key
	// 公钥由私钥计算推导，使用下面的调用
	//publicKey publicKey = privateKey.PublicKey()
	// 助记词
	mnemonic string
	// 地址
	Address Address
}

func NewWallet(pass string)*Wallet  {
	w:=&Wallet{}
	//生成秘钥
	w.GenKey(pass)
	//生成地址
	w.GenAddress()
	return w
}
//生成key
func (w*Wallet)GenKey(pass string)*Wallet  {
	//使用bip39
	//熵（随机）
	//elliptic.p256()生成椭圆
	//rand.Reader,生成随机数
	// 使用 bip39
	// 熵（随机）
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		log.Fatal(err)
	}
	// 助记词
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		log.Fatal(err)
	}
	w.mnemonic = mnemonic

	// key 的种子
	seed := bip39.NewSeed(mnemonic, pass)
	// 生成秘钥
	privateKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		log.Fatal(err)
	}
	w.privateKey = privateKey

	return w
	/*privateKey, err :=ecdsa.GenerateKey(elliptic.P256(),rand.Reader)
	if err!=nil{
		log.Fatal(err)
	}
	w.privateKey=privateKey
	return w*/
}
//生成Address
func (w*Wallet)GenAddress()*Wallet  {
	/*//利用私钥形参公钥
	pubKey:=w.genPubKey()//[]byte
	shaHash:=sha256.Sum256(pubKey)
	rpmd:=ripemd160.New()
	rpmd.Write(shaHash[:])
	pubHash:=rpmd.Sum(nil)
	//计算chenckSum校验值
	h1:=sha256.Sum256(pubKey)
	checkSum:=sha256.Sum256(h1[:])
	//组合继续base64
	data:=append(append([]byte{0},pubHash...),checkSum[:2]...)
	w.Address =base58.Encode(data)
	return w*/
	// 利用私钥获取公钥字符串
	pubKey := w.privateKey.PublicKey().String()

	// hash pubKey
	hashPubKey := HashPubKey([]byte(pubKey))

	// 计算checkSum 校验值
	h1 := sha256.Sum256(append([]byte{0}, hashPubKey...))
	checkSum := sha256.Sum256(h1[:])

	// 组合继续base58
	data := append(append([]byte{version}, hashPubKey...), checkSum[:checkSumLen]...)
	w.Address = base58.Encode(data)

	return w
}
//利用私钥生成公钥
/*func (w*Wallet)genPubKey()[]byte  {
	pubKey:=append(
		w.privateKey.PublicKey.X.Bytes(),
		w.privateKey.PublicKey.Y.Bytes()...
		)
	return pubKey
}*/
// 生成公钥的hash值
func HashPubKey(pubKey []byte) []byte {
	// pubHash: ripemd160(sha256(pubkey))
	shaHash := sha256.Sum256([]byte(pubKey))
	rpmd := ripemd160.New()
	rpmd.Write(shaHash[:])
	pubHash := rpmd.Sum(nil)
	return pubHash
}
func (w *Wallet) GetMnemonic() string {
	return w.mnemonic
}