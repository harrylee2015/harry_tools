package main

import (
	"gitlab.33.cn/chain33/chain33/types"
	//"github.com/gogo/protobuf/proto"
	"fmt"
)

func main() {
	//var replys []*types.ReplyGame
   //game :=&types.Game{
	//   GameId:        "xxxxxxxx",
	//   Value:         20,
	//   HashType:      "sha256",
	//   HashValue:     []byte{1,23,45,67,23},
	//   CreateTime:    11111111111,
	//   CreateAddress: "xyyyyyyyyyyy",
	//   Status:        0,
   //}
   //replyGame :=&types.ReplyGame{game}
   ////replys = append(replys,replyGame)
   ////game.Status=20
   ////replys = append(replys,replyGame)
	//output(reply(replyGame))
//	j :=14
//for i:=0;i<=14;i+=5+1{
//	fmt.Println(i)
//}
}
func reply(reply interface{})*types.Message{
	if res,ok :=reply.(types.Message);ok{
		fmt.Println("11111111111")
      return &res
	}
	return nil
}
func output(reply interface{})interface{}{
	 if res,ok :=reply.(*types.Message);ok {
	 	fmt.Println("xxxxxx")
	 	fmt.Println(*res)
	 	if result,ok1:=(*res).(*types.ReplyGame);ok1{
			fmt.Println("xxxxxx")
			 fmt.Println(result)
		 }
	 }
	return reply
}