package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/yytt5301/wkim_gosdk/pkg/wksdk"

	wkproto "github.com/WuKongIM/WuKongIMGoProto"
)

func ExObjectString(obj interface{}) string {
	b, _ := json.Marshal(obj)
	return string(b)
}

func userWork(uid, token string) {
	// ================== client 2 ==================
	// 初始化客户端
	cli2 := wksdk.NewClient("tcp://192.168.101.108:5100",
		wksdk.WithUID(uid),
		wksdk.WithToken(token),
		wksdk.WithReconnect(false), wksdk.WithPingInterval(time.Second*5))

	err := cli2.Connect()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	fmt.Println("connect ok.")
	cli2.OnMessage(func(msg *wksdk.Message) {
		fmt.Printf("FromUID=%s, ChannelID=%s, ChannelType=%d\n", msg.FromUID, msg.ChannelID, msg.ChannelType)
		recvPayload := string(msg.Payload)
		fmt.Printf("[%s]==>[%s] receive msg: %s\n", msg.FromUID, uid, string(msg.Payload))

		if recvPayload == "exit" {
			wg.Done()
		} else {
			resMsg := botResponse(recvPayload)
			_, err = cli2.SendMessage([]byte(resMsg), wkproto.Channel{
				ChannelType: msg.ChannelType,
				ChannelID:   msg.FromUID,
			})
		}
	})

	wg.Add(1)
	wg.Wait()
	fmt.Println("userWork() exit")
}

func botResponse(reqMsg string) string {
	return `{"content":"ok","type":1}`
}

func main() {
	fmt.Println(time.Now().String(), "start...")
	// 主循环读取用户输入
	uid := "test2"
	token := "abc"
	go userWork(uid, token)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "q" || input == "Q" {
			fmt.Println("用户输入 q，程序退出")
			break
		}
		// fmt.Println(input)
	}
}
