package game

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/33cn/chain33/common"
	"github.com/33cn/chain33/common/address"
	"github.com/33cn/chain33/common/crypto"
	cty "github.com/33cn/chain33/system/dapp/coins/types"
	"github.com/33cn/chain33/types"
	gt "github.com/33cn/plugin/plugin/dapp/game/types"
	"io/ioutil"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

const (
	FuncName_QueryGameListByIds = "QueryGameListByIds"
	//组合查询
	FuncName_QueryGameListByStatusAndAddr = "QueryGameListByStatusAndAddr"
	FuncName_QueryGameById                = "QueryGameById"
	FuncName_QueryGameListCount           = "QueryGameListCount"
	GameX                                 = "game"
	ConfName_ActiveTime                   = GameX + ":" + "activeTime"
	ConfName_DefaultCount                 = GameX + ":" + "defaultCount"
	ConfName_MaxCount                     = GameX + ":" + "maxCount"
	ConfName_MaxGameAmount                = GameX + ":" + "maxGameAmount"
	ConfName_MinGameAmount                = GameX + ":" + "minGameAmount"
)

var (
	//Jrpc_Url = "http://139.219.129.114:9801"
	//Jrpc_Url = "http://139.219.129.114:8801"

	//Jrpc_Url = "http://47.99.129.127:9801"
	Jrpc_Url = "http://192.168.0.194:8801"
	//Jrpc_Url = "http://47.99.129.127:8801"
	// 石头剪刀布游戏，比特元打币地址
	rpsGameAddr = "19MJmA7GcE1NfMwdGqgLJioBjVbzQnVYvR"
	// 石头剪刀布游戏，地址对应私钥
	rpsGamePriv = "5072a3b6ed612845a7c00b88b38e4564093f57ce652212d6e26da9fded83e951"

	// 10个测试地址,对应10个参与游戏的用户
	// cli查询语句：./cli  account balance -a 1Ft6SUwdnhzryBwPDEnchLeNHgvSAXPbAK  -e user.p.fzmtest.coins
	// ./cli  account balance -a 1Ft6SUwdnhzryBwPDEnchLeNHgvSAXPbAK  -e user.p.fzmtest.game
	rpsTestAddress = [10]string{"1Ft6SUwdnhzryBwPDEnchLeNHgvSAXPbAK", "1HHy3xZEtEYHjPJYkvWvR25CDPZNLkwZQ4", "1AuL12ZPYxBV9Pvx1LFcazS5W8pwmtsA4y", "15PXh6D4zLBgGbRFfbssni3Q2C35AvkTxX",
		"18pdQ8UoZkhwGWx3UgE6cmy5G2K2QaS4wS", "1Cwu2kCQTth6YmXki5rgg6DEKj8Rys9fu1", "1GTqLXA8cjUKWDtXjax9Acs2itwbAQierf", "1GaPfaU8dSAdipzUtWevZq8fMwvKWHnVFc",
		"1CtY2o6AAJv8ae7B4ptADRVnkALvPERYoJ", "1NuPawKogkafCncNVmpHmg3qEvCx8GcNWj"}

	rpsTestPrivA = [10]string{"b7c84e44279384e415cac9252500d64ecbbdc0306db77af3029fa26769247a7d", "f56cca65e772931231c92ff4ef0b518c2a92a02c4edd2f1972663c6d3e2d0e24",
		"72a829b5fb54dd7069f62f835fd694cc0887a77f56653423d6c0c1e6cd1eba0c", "8241e88a2ee69eb2a922bd053ce1e9d1df26e317325f6b34405f39021e856f1d",
		"b2c789fa5b4fc129201c960dca877bf530fdea9e0067d70b91f1e829b2ab5079", "71e5c6e084b71fa0273414196a4d7009df51bde29559320bd4a3635ef1fdfd2f",
		"74900ae36e709ff8679fd0cfc0e8a58b9933fea5dd0cf3e1fc295027179f6d63", "fd367b455b088c291b1e37144f7f83a72b5e34c10660c0dbc3d0ec4a039ed18d",
		"61b386c0c15b9931da48d9cda7dbcd23c78e12b78be92fc0a2e76c6de283044f", "868f59b43958db9d3ac93b62930d7ac7744cc23bf965801c87af3b2c8d4898a6"}

	rpsTestPrivB = [10]string{"61b386c0c15b9931da48d9cda7dbcd23c78e12b78be92fc0a2e76c6de283044f", "868f59b43958db9d3ac93b62930d7ac7744cc23bf965801c87af3b2c8d4898a6",
		"b7c84e44279384e415cac9252500d64ecbbdc0306db77af3029fa26769247a7d", "f56cca65e772931231c92ff4ef0b518c2a92a02c4edd2f1972663c6d3e2d0e24",
		"72a829b5fb54dd7069f62f835fd694cc0887a77f56653423d6c0c1e6cd1eba0c", "8241e88a2ee69eb2a922bd053ce1e9d1df26e317325f6b34405f39021e856f1d",
		"fd367b455b088c291b1e37144f7f83a72b5e34c10660c0dbc3d0ec4a039ed18d", "74900ae36e709ff8679fd0cfc0e8a58b9933fea5dd0cf3e1fc295027179f6d63",
		"b2c789fa5b4fc129201c960dca877bf530fdea9e0067d70b91f1e829b2ab5079", "71e5c6e084b71fa0273414196a4d7009df51bde29559320bd4a3635ef1fdfd2f"}

	// 设置10局游戏，对应的10次出拳(A用户首先出拳)
	rpsGuessA = [10]int32{1, 1, 3, 2, 1, 3, 1, 2, 3, 2}
	// 设置10局游戏，对应的10次出拳(B用户的出拳：胜，平，平，负，胜，平，平，负，胜，平)
	rpsGuessB = [10]int32{1, 3, 2, 1, 2, 3, 1, 2, 1, 2}
	secet     = [10]string{"ABCDEF", "BACDEF", "ACBDEF", "ABDCEF", "ABCEDF", "ABCDFE", "CDEFAB", "ABEFCD", "ADEFBC", "ABCFDE"}
	// 游戏币创世地址
	GenesisAddr = "14KEKbYtKKQm4wMthSK9J4La4nAiidGozt"
	// 游戏币地址对应的私钥
	GenesisPriv = "cc38546e9e659d15e6b4893f0ab32a06d103931a8230b0bde71459d2b27d6944"
	ExecName    = "game"
	ParaName    = "user.p.fzmtest."

	//ParaName = ""
	Secet = "ABCDEF"
	//剪刀
	Scissor = 1
	//石头
	Rock = 2
	//布
	Paper = 3
)

func getRealExecName(execName string) string {
	return ParaName + execName
}
func getJrpc() string {
	return Jrpc_Url
}
func TestTranfer(t *testing.T){
	tx := createTxTranferToAddress(50, "test!", "1CDyeoyRFBnsqpVUVajtJnpSEr5enq4CYs", rpsGamePriv)
	txhash,err:=SendTx(tx)
	if err !=nil {
		t.Error(err)
	}
	t.Log(txhash)
}
func main() {

	//主链上，给rpsGameAddr地址打比特元（买币操作）
	//for i := 0; i < len(rpsTestAddress); i++ {
	//	tx :=createBtyTranfer(1, "tranfer to rpsGameAddr 1 BTY.",rpsGameAddr,rpsTestPrivA[i])
	//	txHash,_:=SendTx(tx)
	//	fmt.Println(txHash)
	//	time.Sleep(10 * time.Second)
	//}

	//time.Sleep(300*time.Second)

	// 平行链上，给创世地址打回游戏币（提币操作）
	//for i := 0; i < len(rpsTestAddress); i++ {
	//	tx := createTxTranferToAddress(1,"tranfer to genesisAddr 1 game coin",GenesisAddr,rpsTestPrivA[i])
	//	txHash,_:=SendTx(tx)
	//	fmt.Println(txHash)
	//	time.Sleep(3 * time.Second)
	//}

	//主链上，从rpsGameAddr地址给用户地址打比特元
	//for i := 0; i < len(rpsTestAddress); i++ {
	//	tx :=createBtyTranfer(1, "tranfer to UserAddress 1 BTY.",rpsTestAddress[i],rpsGamePriv)
	//	SendTx(tx)
	//	time.Sleep(10 * time.Second)
	//}

	//上一步成功后，从平行链的创世地址往10个测试地址中打入游戏币
	//for i := 0; i < 1; i++ {
	//	for i := 0; i < len(rpsTestAddress); i++ {
	tx := createTxTranferToAddress(50, "test!", "1CDyeoyRFBnsqpVUVajtJnpSEr5enq4CYs", rpsGamePriv)
	//		tx := createTxTranferToAddress(10000000,"fzm_test","1Dtp4BkyoKhUahz4mMcEuR2pZVaiTJwKD7",GenesisPriv)
	SendTx(tx)
	//		fmt.Println(txHash)
	//		time.Sleep(3 * time.Second)
	//     }

	//等待打包
	//time.Sleep(20*time.Second)
	//fmt.Println("============================ 游戏创世账户往用户账户打游戏币完成===============================")

	// 用户A往合约地址打入游戏币,锁定在合约地址中
	//for i := 0; i < len(rpsTestAddress); i++ {
	//	tx := createTxTranferToExec(50,"test game exector!","game",rpsTestPrivA[i],false)
	//	txHash,_:=SendTx(tx)
	//	fmt.Println(txHash)
	//	time.Sleep(3 * time.Second)
	//}
	//game.proto
	//	//等待打包
	//	time.Sleep(20*time.Second)
	//	fmt.Println("============================ 用户游戏账户往合约中锁定游戏币完成===============================")
	//bufGuess := []byte("123456" + string(1))
	//fmt.Println(common.Sha256(bufGuess))

	//创建游戏
	//gameIds := []string{}
	//for i := 0; i < 1; i++ {
	//	for i := 0; i < len(rpsGuessA); i++ {
	//		bufGuess := []byte(secet[i] + string(rpsGuessA[i]))
	//		tx := createGame(2, "sha256", common.Sha256(bufGuess), "game", rpsTestPrivA[i])
	//		_, err := SendTx(tx)
	//		if err != nil {
	//			fmt.Errorf("can't get gameId!")
	//			panic(err)
	//		}
	//	}
	//time.Sleep(30 * time.Second)

	// 根据ID查询游戏
	//queryGameById(gameId, ExecName)
	//gameIds = append(gameIds, gameId)

	// 根据游戏状态和游戏地址来查询游戏
	//queryGameListByStatusAndAddr(1, rpsTestAddress[i], ExecName)

	//time.Sleep(30 * time.Second)
	//// 根据游戏列表来查询游戏
	//queryGameListByIds(gameIds, ExecName)
	//
	//fmt.Println("============================ 游戏创建完成===============================")

	//取消游戏
	//for i := 0; i < 1; i++ {
	//	//for i := 0; i < len(gameIds); i++ {
	//	tx := cancelGame(gameIds[i],"game",rpsTestPrivA[i])
	//	SendTx(tx)
	//	time.Sleep(30*time.Second)
	//
	//	// 根据ID查询游戏
	//	queryGameById(gameIds[i], ExecName)
	//
	//	// 根据游戏状态和游戏地址来查询游戏
	//	queryGameListByStatusAndAddr(2, rpsTestAddress[i], ExecName)
	//}
	//time.Sleep(30 * time.Second)
	//// 根据游戏列表来查询游戏
	//queryGameListByIds(gameIds, ExecName)
	//fmt.Println("============================ 取消游戏完成 ===============================")

	// 用户B往合约地址打入游戏币,锁定在合约地址中
	//for i := 0; i < len(rpsTestAddress); i++ {
	//	tx := createTxTranferToExec(100,"test game exector!","game",rpsTestPrivB[i],false)
	//	txHash,_:=SendTx(tx)
	//	fmt.Println(txHash)
	//	time.Sleep(3 * time.Second)
	//}
	//
	////等待打包
	//time.Sleep(20*time.Second)
	//fmt.Println("============================ 用户游戏账户往合约中锁定游戏币完成===============================")

	//匹配游戏
	//for i := 0; i < 1; i++ {
	////for i := 0; i < len(gameIds); i++ {
	//	tx := matchGame(gameIds[i],rpsGuessB[i],"game",rpsTestPrivB[i])
	//	SendTx(tx)
	//	time.Sleep(30*time.Second)
	//
	//	// 根据ID查询游戏
	//	queryGameById(gameIds[i], ExecName)
	//
	//	// 根据游戏状态和游戏地址来查询游戏
	//	queryGameListByStatusAndAddr(1, rpsTestAddress[i], ExecName)
	//}
	//time.Sleep(30 * time.Second)
	//// 根据游戏列表来查询游戏
	//queryGameListByIds(gameIds, ExecName)
	//fmt.Println("============================ 游戏匹配完成 ===============================")

	// 开奖
	//for i := 0; i < 1; i++ {
	////for i := 0; i < len(gameIds); i++ {
	//	tx := closeGame(gameIds[i],Secet,0,"game",rpsTestPrivA[i])
	//	SendTx(tx)
	//	time.Sleep(30*time.Second)
	//
	//	// 根据ID查询游戏
	//	queryGameById(gameIds[i], ExecName)
	//
	//	// 根据游戏状态和游戏地址来查询游戏
	//	queryGameListByStatusAndAddr(3, rpsTestAddress[i], ExecName)
	//}
	//// 根据游戏列表来查询游戏
	//queryGameListByIds(gameIds, ExecName)
	//fmt.Println("============================ 游戏开奖完成 ===============================")

	//从合约地址中提取游戏币
	//tx :=createTxTranferToExec(15,"withdraw game coins.","game","0x868f59b43958db9d3ac93b62930d7ac7744cc23bf965801c87af3b2c8d4898a6",true)
	//txHash,err:=SendTx(tx)
	//if err !=nil {
	//	fmt.Println("send tx have err",err.Error())
	//}
	//fmt.Println(txHash)
	//queryGameById("0x7d9e1aca128e6ea0f05bbbb49038976950d13858dd43a519f72d1472b14a7720",ExecName)
	//fmt.Println(Sha256([]byte("f1ac06aefde25c7d315e2c0f739234c83f4526c1308956ad7126f2f39151a7cab3568086cc33bb55"+string(1))))
	//fmt.Println(Sha256([]byte("f1ac06aefde25c7d315e2c0f739234c83f4526c1308956ad7126f2f39151a7cab3568086cc33bb55"+string(2))))
	//fmt.Println(Sha256([]byte("f1ac06aefde25c7d315e2c0f739234c83f4526c1308956ad7126f2f39151a7cab3568086cc33bb55"+string(3))))
	//fmt.Println(Sha256([]byte("f1ac06aefde25c7d315e2c0f739234c83f4526c1308956ad7126f2f39151a7cab3568086cc33bb55"+string(0))))

	//queryGameById("0x09d875cedd147390796b0c1d0c8c22b3e770b2eb374e514ec865cebdd8c1ca7c", ExecName)
	//fmt.Println(Sha256([]byte("b7900d29ea5bfd27818ddcaac1d3267739302372598351c3339fff3571be10e7db51d990782011a9"+string(1))))
	//fmt.Println(Sha256([]byte("b7900d29ea5bfd27818ddcaac1d3267739302372598351c3339fff3571be10e7db51d990782011a9"+string(2))))
	//fmt.Println(Sha256([]byte("b7900d29ea5bfd27818ddcaac1d3267739302372598351c3339fff3571be10e7db51d990782011a9"+string(3))))
	//queryGameListByStatusAndAddr(1, "", 0,ExecName)
	//
	//bufGuess := []byte(secet[9] + string(rpsGuessA[9]))
	//tx := createGame(2, "sha256", Sha256(bufGuess), "game", rpsTestPrivA[9])
	//SendTx(tx)
	//0xf56fee4be83ae616c6bc0de6ffda82eee83b6323f6420757b180ce9933195fd5
	//	tx := closeGame("0x207fd5b0a5806c4859dd5213cfbf6b661a89a4563eddb290a198d1cf4b76e1da","123456",0,"game",rpsTestPrivA[9])
	//tx := matchGame("0x3b1447f9278207e6e5b8e60046103b509ff219603ce5fbba046d8554cc3e07a6",2,"game",rpsTestPrivA[1])
	//tx := matchGame("0x5660435b3d8a2d09bca17b43b4c3a16adad943354642d3b454f94cd3c04d1bd4",rpsGuessA[9],"game",rpsTestPrivA[9])
	//SendTx(tx)
	//time.Sleep(time.Second)
	//queryGameListByStatusAndAddr(1, "1KgE3vayiqZKhfhMftN7vt2gDv9HoMk941", 0, ExecName)

	//Ids :=[]string{"0x459249cbf17bb550c90b34abb61c342553f4bc6083eba7bbec31ba37672254e5","0x26e975c9c26d30d929e3b8ed954fca1b214d67b84db23a1386b8bbd9cd602249",
	//"0x3bb2cedfdc0e82d93a7d3a4257033eae7738bc3ea2616a80b6ace859b9a915c1","0x7901a69b8508106081cef0649006ae9aaf5bf2b2ddc99e5a5effe8f6c9834dab"
	//,"0xff03c7626fc930c5332c824a7c9047a3ecf099a6aad7d86519f0b2d26178d03e",""}
	//  545  476  59  424
	//queryGameListCount(4,"",ExecName)
	//tx := createManagerTx(ConfName_MaxCount,"50","cc38546e9e659d15e6b4893f0ab32a06d103931a8230b0bde71459d2b27d6944")
	//SendTx(tx)
	//tx := createManagerTx(ConfName_MaxGameAmount,"100","cc38546e9e659d15e6b4893f0ab32a06d103931a8230b0bde71459d2b27d6944")
	//SendTx(tx)
	//tx = createManagerTx(ConfName_MinGameAmount,"2","cc38546e9e659d15e6b4893f0ab32a06d103931a8230b0bde71459d2b27d6944")
	//SendTx(tx)
	//tx = createManagerTx(ConfName_ActiveTime,"24","cc38546e9e659d15e6b4893f0ab32a06d103931a8230b0bde71459d2b27d6944")
	//SendTx(tx)
}
func Sha256(bytes []byte) []byte {
	data := sha256.Sum256(bytes)
	return data[:]
}

//在主链上构建bty的转账交易
func createBtyTranfer(amount int64, note, to, HexPri string) *types.Transaction {
	transfer := &cty.CoinsAction{}

	v := &cty.CoinsAction_Transfer{Transfer: &types.AssetsTransfer{Amount: amount * 1e8, Note: []byte("test"), To: to}}
	transfer.Value = v
	transfer.Ty = cty.CoinsActionTransfer

	tx := &types.Transaction{Execer: []byte("coins"), Payload: types.Encode(transfer), To: to}
	tx.Fee, _ = tx.GetRealFee(1e5)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	tx.Nonce = random.Int63()
	cr, _ := crypto.New("secp256k1")
	hexbytes, _ := common.FromHex(HexPri)
	priv, _ := cr.PrivKeyFromBytes(hexbytes)
	tx.Sign(types.SECP256K1, priv)
	return tx
}

//在平行链上构建交易
func createTxTranferToAddress(amount int64, note, to, HexPri string) *types.Transaction {
	transfer := &cty.CoinsAction{}
	v := &cty.CoinsAction_Transfer{Transfer: &types.AssetsTransfer{Amount: amount * 1e8, Note: []byte("test"), To: to}}
	transfer.Value = v
	transfer.Ty = cty.CoinsActionTransfer

	var tx *types.Transaction
	tx = &types.Transaction{Execer: []byte(getRealExecName("coins")), Payload: types.Encode(transfer), To: address.ExecAddress(getRealExecName("coins"))}
	tx.Fee, _ = tx.GetRealFee(1e5)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	tx.Nonce = random.Int63()
	cr, _ := crypto.New("secp256k1")
	hexbytes, _ := common.FromHex(HexPri)
	priv, _ := cr.PrivKeyFromBytes(hexbytes)
	tx.Sign(types.SECP256K1, priv)
	return tx
}

// 往游戏合约中打入游戏币 or 从合约中提取游戏币
func createTxTranferToExec(amount int64, note, execName, HexPri string, isWithdraw bool) *types.Transaction {
	initExecName := execName
	transfer := &cty.CoinsAction{}
	if !isWithdraw {
		if initExecName != "" {
			v := &cty.CoinsAction_TransferToExec{TransferToExec: &types.AssetsTransferToExec{Amount: amount * 1e8, Note: []byte("test"), ExecName: getRealExecName(execName), To: address.ExecAddress(getRealExecName(execName))}}
			transfer.Value = v
			transfer.Ty = cty.CoinsActionTransferToExec
		} else {
			v := &cty.CoinsAction_Transfer{Transfer: &types.AssetsTransfer{Amount: amount * 1e8, Note: []byte("test"), To: address.ExecAddress(getRealExecName(execName))}}
			transfer.Value = v
			transfer.Ty = cty.CoinsActionTransfer
		}
	} else {
		v := &cty.CoinsAction_Withdraw{Withdraw: &types.AssetsWithdraw{Amount: amount * 1e8, Note: []byte("test"), ExecName: getRealExecName(execName), To: address.ExecAddress(getRealExecName(execName))}}
		transfer.Value = v
		transfer.Ty = cty.CoinsActionWithdraw
	}
	var tx *types.Transaction
	tx = &types.Transaction{Execer: []byte(getRealExecName("coins")), Payload: types.Encode(transfer), To: address.ExecAddress(getRealExecName("coins"))}
	tx.Fee, _ = tx.GetRealFee(1e5)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	tx.Nonce = random.Int63()
	cr, _ := crypto.New("secp256k1")
	hexbytes, _ := common.FromHex(HexPri)
	priv, _ := cr.PrivKeyFromBytes(hexbytes)
	tx.Sign(types.SECP256K1, priv)
	return tx
}

//创建游戏
func createGame(amount int64, hashType string, hashValue []byte, execName, HexPri string) *types.Transaction {
	v := &gt.GameCreate{
		Value:     amount * 1e8,
		HashType:  hashType,
		HashValue: hashValue,
	}
	precreate := &gt.GameAction{
		Ty:    gt.GameActionCreate,
		Value: &gt.GameAction_Create{v},
	}
	tx := &types.Transaction{
		Execer:  []byte(getRealExecName(execName)),
		Payload: types.Encode(precreate),
		Fee:     1e5,
		Nonce:   rand.New(rand.NewSource(time.Now().UnixNano())).Int63(),
		To:      address.ExecAddress(getRealExecName(execName)),
	}

	tx.SetRealFee(1e5)
	cr, _ := crypto.New("secp256k1")
	hexbytes, _ := common.FromHex(HexPri)
	priv, _ := cr.PrivKeyFromBytes(hexbytes)
	tx.Sign(types.SECP256K1, priv)
	return tx
}

//匹配游戏
func matchGame(gameId string, guess int32, execName, HexPri string) *types.Transaction {
	v := &gt.GameMatch{
		GameId: gameId,
		Guess:  guess,
	}
	game := &gt.GameAction{
		Ty:    gt.GameActionMatch,
		Value: &gt.GameAction_Match{v},
	}
	tx := &types.Transaction{
		Execer:  []byte(getRealExecName(execName)),
		Payload: types.Encode(game),
		Fee:     1e5,
		Nonce:   rand.New(rand.NewSource(time.Now().UnixNano())).Int63(),
		To:      address.ExecAddress(getRealExecName(execName)),
	}

	tx.SetRealFee(1e5)
	cr, _ := crypto.New("secp256k1")
	hexbytes, _ := common.FromHex(HexPri)
	priv, _ := cr.PrivKeyFromBytes(hexbytes)
	tx.Sign(types.SECP256K1, priv)
	return tx
}

//取消游戏
func cancelGame(gameId string, execName, HexPri string) *types.Transaction {
	v := &gt.GameCancel{
		GameId: gameId,
	}
	cancel := &gt.GameAction{
		Ty:    gt.GameActionCancel,
		Value: &gt.GameAction_Cancel{v},
	}
	tx := &types.Transaction{
		Execer:  []byte(getRealExecName(execName)),
		Payload: types.Encode(cancel),
		Fee:     1e5,
		Nonce:   rand.New(rand.NewSource(time.Now().UnixNano())).Int63(),
		To:      address.ExecAddress(getRealExecName(execName)),
	}

	tx.SetRealFee(1e5)
	cr, _ := crypto.New("secp256k1")
	hexbytes, _ := common.FromHex(HexPri)
	priv, _ := cr.PrivKeyFromBytes(hexbytes)
	tx.Sign(types.SECP256K1, priv)
	return tx
}

//结束游戏
func closeGame(gameId, secret string, result int32, execName, HexPri string) *types.Transaction {
	v := &gt.GameClose{
		GameId: gameId,
		Secret: secret,
	}
	close := &gt.GameAction{
		Ty:    gt.GameActionClose,
		Value: &gt.GameAction_Close{v},
	}
	tx := &types.Transaction{
		Execer:  []byte(getRealExecName(execName)),
		Payload: types.Encode(close),
		Fee:     1e5,
		Nonce:   rand.New(rand.NewSource(time.Now().UnixNano())).Int63(),
		To:      address.ExecAddress(getRealExecName(execName)),
	}

	tx.SetRealFee(1e5)
	cr, _ := crypto.New("secp256k1")
	hexbytes, _ := common.FromHex(HexPri)
	priv, _ := cr.PrivKeyFromBytes(hexbytes)
	tx.Sign(types.SECP256K1, priv)
	return tx
}

// 签名交易
func SendTx(tx *types.Transaction) (string, error) {
	hexTx := hex.EncodeToString(types.Encode(tx))
	poststr := fmt.Sprintf(`{"jsonrpc":"2.0","id":2,"method":"Chain33.SendTransaction","params":[{"data":"%v"}]}`,
		hexTx)

	resp, err := http.Post(getJrpc(), "application/json", bytes.NewBufferString(poststr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var txdata = make(map[string]interface{})
	fmt.Println("resp:", string(data))
	err = json.Unmarshal(data, &txdata)
	if err != nil {
		fmt.Println("err:", err.Error())
		return "", err
	}
	if hextx, ok := txdata["result"]; ok {
		return hextx.(string), nil
	}
	return "", fmt.Errorf("not have result!")
}

func queryGameById(gameId string, execName string) {
	funcName := FuncName_QueryGameById
	payload := fmt.Sprintf(`{"gameId":"%v"}`, gameId)
	Query(execName, funcName, payload)
}

func queryGameListByIds(gameIds []string, execName string) {
	funcName := FuncName_QueryGameListByIds
	var str string
	for _, gameId := range gameIds {
		if str == "" {
			str = fmt.Sprintf(`"%v"`, gameId)
		} else {
			str = fmt.Sprintf(`%v,"%v"`, str, gameId)
		}
	}
	payload := fmt.Sprintf(`{"gameIds":[%v]}`, str)
	Query(execName, funcName, payload)
}
func queryGameListByStatusAndAddr(status int32, addr string, index int64, execName string) {
	funcName := FuncName_QueryGameListByStatusAndAddr
	payload := fmt.Sprintf(`{"status":%v,"address":"%v","index":%v,"count":%v,"direction":%v}`, status, addr, index, 10, 0)
	Query(execName, funcName, payload)
}

func queryGameListCount(status int32, addr string, execName string) {
	funcName := FuncName_QueryGameListCount
	payload := fmt.Sprintf(`{"status":%v,"address":"%v"}`, status, addr)
	Query(execName, funcName, payload)
}

func Query(execName, funcName, payload string) {
	execName = getRealExecName(execName)
	poststr := fmt.Sprintf(`{"jsonrpc":"2.0","id":5,"method":"Chain33.Query",
"params":[{"execer":"%v","funcName":"%v","payload":%v}]}`,
		execName, funcName, payload)
	resp, err := http.Post(getJrpc(), "application/json", bytes.NewBufferString(poststr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var txdata = make(map[string]interface{})
	fmt.Println("resp:", string(data))
	err = json.Unmarshal(data, &txdata)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	if hextx, ok := txdata["result"]; ok {
		data2, err := json.Marshal(hextx)
		if err != nil {
			fmt.Println("marshal have err:", err.Error())
		}
		var game gt.Game
		err = json.Unmarshal(data2, &game)
		if err != nil {
			fmt.Println("unmarshal have err:", err.Error())
		}
		fmt.Println(game.GetHashValue())
		return
	}
	//fmt.Println(string(data))

}

func QueryTransaction(txHash string) {
	poststr := fmt.Sprintf(`{"jsonrpc":"2.0","id":11,"method":"Chain33.QueryTransaction","params":[{"hash":"%v"}]}`, txHash)
	resp, err := http.Post(getJrpc(), "application/json", bytes.NewBufferString(poststr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("returned JSON: %s\n", string(b))
}
