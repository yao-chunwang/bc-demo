package main

//import "golang.org/x/crypto/ripemd160"
import (
	"fmt"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)
//"github.com/tyler-smith/go-bip32"


func main() {
	//// 数据库连接
	//dbpath := "data"
	//db, err := leveldb.OpenFile(dbpath, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// 释放数据库连接
	//defer db.Close()
	//
	//// 区块链测试
	//bc := blockchain.NewBlockchain(db)
	//// 添加创世区块
	//bc.AddGensisBlock()
	//// 添加新区块
	//bc.
	//	AddBlock("First Block").
	//	AddBlock("Second Block")
	//
	//bc.Iterate()
	//bits := 2
	//nonce := 0
	//for {
	//	data := "block data" + strconv.Itoa(nonce)
	//	hash := sha256.Sum256([]byte(data))
	//	h := fmt.Sprintf("%x", hash)
	//	fmt.Println(h, nonce)
	//	if strings.HasPrefix(h, "000000") {
	//		fmt.Printf("本机挖矿成功")
	//		return
	//	}
	//	nonce ++
	//}


	// hashcash
	//bits := 16 // 256 前 8 位为 0
	//target := big.NewInt(1) // 00000 ... 0001
	//// 00000000 10000000 0000000 0000000 ...... 256
	//// 采用左移位的方案，构建目标比较数
	//// 00000001 LSH 1 =  00000010
	//// 00000001 LSH 2 =  00000100
	//target.Lsh(target, uint(256-bits+1))
	//fmt.Println(target.String())
	//fmt.Println("----------Minting------------")
	//nonce := 0
	//// 服务字符串
	//serviceStr := "block data"
	//var hashInt big.Int
	//for {
	//	// 服务字符串 连接 随机数
	//	data :=  serviceStr + strconv.Itoa(nonce)
	//	hash := sha256.Sum256([]byte(data))
	//	hashInt.SetBytes(hash[:])
	//
	//	fmt.Println(hashInt.String(), nonce)
	//	if hashInt.Cmp(target) == -1 { // compare, hashInt < target
	//		fmt.Printf("本机挖矿成功")
	//		return
	//	}
	//	nonce ++
	//}


	//t := time.Now()
	//bufer := bytes.Buffer{}
	//enc := gob.NewEncoder(&bufer)
	//enc.Encode(t)
	//
	//dec := gob.NewDecoder(&bufer)
	//var t1 time.Time
	//dec.Decode(&t1)
	////fmt.Println(t.String())
	////fmt.Println(t1.String())
	//fmt.Println(t.Format("2006-01-02 15:04:05.999999999 -0700 MST"))
	//fmt.Println(t1.Format("2006-01-02 15:04:05.999999999 -0700 MST"))

	//rh := ripemd160.New() // 新建hash器）
	//rh.Write([]byte("blockchain"))
	//fmt.Println(rh.Sum(nil))

	// base58
	//encoded := "1QCaxc8hutpdZ62iKZsn1TCG3nh7uPZojq"
	//num, err := base58.Decode(encoded)
	//if err != nil {
	//	fmt.Printf("Demo %v, got error %s\n", encoded, err)
	//}
	//
	//chk := base58.Encode(num)
	//if encoded == string(chk) {
	//	fmt.Printf ( "Successfully decoded then re-encoded %s\n", encoded )
	//}


	// bip39
	// Generate a mnemonic for memorization or user-friendly seeds
	//// 熵，理解成，混乱程度
	//entropy, _ := bip39.NewEntropy(256)
	//// 助记词
	//mnemonic, _ := bip39.NewMnemonic(entropy)
	//fmt.Println(mnemonic)
	//// 生成密钥对
	//// 基于助记词和短语密码（用户需要指定的）生成密钥种子
	//seed := bip39.NewSeed(mnemonic, "Secret Passphrase")
	//// 构建私钥，基于种子
	//masterKey, _ := bip32.NewMasterKey(seed)
	//// 构建公钥
	//publicKey := masterKey.PublicKey()
	//
	//fmt.Println("PrivateKey: ", masterKey.String())
	//fmt.Println("PublicKey: ", publicKey.String())

	// 反推
	userMnemonic := "inform undo plate foil orient where virus balcony team cargo loyal measure easily category coach develop minute copper loyal math spy hawk figure tornado"
	// 生成熵
	//userEntropy, _ := bip39.EntropyFromMnemonic(userMnemonic)
	// 生成种子
	userSeed := bip39.NewSeed(userMnemonic, "Secret Passphrase")
	userMasterKey, _ := bip32.NewMasterKey(userSeed)
	userPublicKey := userMasterKey.PublicKey()
	fmt.Println("PrivateKey: ", userMasterKey.String())
	fmt.Println("PublicKey: ", userPublicKey.String())


}
