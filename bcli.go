package main

import (
	"bc-demo/blockchain"
	"bc-demo/wallet"
	"flag"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"os"
	"strings"
)

func main()  {
	//初始化数据库
	//数据库连接
	dbpath:="data"
	db,err:=leveldb.OpenFile(dbpath,nil)
	if err!=nil{
		log.Fatal(err)
	}
	//释放数据库连接
	defer  db.Close()
	//初始化区块链
	bc:=blockchain.NewBlockchain(db)

	//初始化第一个命令参数
	arg1:=""
	//若用户指定了参数，则第一个用户参数为命令参数
	if len(os.Args)>=2{
		arg1=os.Args[1]
	}
	//基于命令参数，执行对应的功能
	switch strings.ToLower(arg1) {
	case "create:block":
		// 为 createblock 命令增加一个 flag 集合。标志集合
		// 错误处理为，一旦解析失败，则 exit
		fs := flag.NewFlagSet("create:block", flag.ExitOnError)
		// 在集合中，添加需要解析的 flag 标志
		address := fs.String("address", "", "")
		// 解析命令行参数,
		fs.Parse(os.Args[2:])
		// 完成区块的创建
		bc.AddBlock(*address)
		//生成钱包地址
	case "create:wallet":
		// 命令行标志集（参数集 -flag）
		fs := flag.NewFlagSet("create:wallet", flag.ExitOnError)
		// pass 标志, *string
		pass := fs.String("pass", "", "")
		w := wallet.NewWallet(*pass)
		fmt.Printf("your mnemonic: %s\n", w.GetMnemonic())
		fmt.Printf("your address: %s \n", w.Address)
	case "show":
		bc.Iterate()
	case "balance":
		fs := flag.NewFlagSet("balance", flag.ExitOnError)
		address := fs.String("address", "", "")
		fs.Parse(os.Args[2:])
		// 完成区块的创建
		fmt.Printf("Address:%s\nBalance:%d\n",
			*address, bc.GetBalance(*address),
		)
	case "init":
		fs:=flag.NewFlagSet("init",flag.ExitOnError)
		address:=fs.String("address","","")
		fs.Parse(os.Args[2:])
		//清空
		bc.Clear()
		//添加创世区块
		bc.AddGensisBlock(*address)
	case "help":
		fallthrough
	default:
		Usage()
		
	}
}
//输出bcli的帮助信息
func Usage()  {
	fmt.Println("bcli is a tool for Blockchain")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Printf("\t%s\t\t%s\n", "bcli create:block -txs=<txs>", "create block on blockchain")
	fmt.Printf("\t%s\t\t%s\n", "bcli create:wallet -pass=<pass>", "create wallet base on pass")
	fmt.Printf("\t%s\t%s\n", "bcli init -address=<address>", "initial blockchain.")
	fmt.Printf("\t%s\t%s\n", "bcli balance -address=<address>", "get address 's balance.")
	fmt.Printf("\t%s\t\t\t%s\n", "bcli help", "help info for bcli.")
	fmt.Printf("\t%s\t\t\t%s\n", "bcli show", "show blocks in chain.")
}