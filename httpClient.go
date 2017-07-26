package main

import (
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
	"github.com/golang/protobuf/proto"
	"http-client-go/oidmc_moclient"
	"crypto/md5"
	"strconv"
	"io"
	"encoding/hex"
	"crypto/tls"
)

const (
	URL = "https://172.16.39.136:15423/mo_order"
)

func main() {
	httpPost()
}

func httpPost() {

	orderId := uint32(1)
	userId := uint32(1)
	userName := "hello"

	body, _ := getBody(orderId, userId, userName)

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}


	resp, err := client.Post(URL,
		"text/plain;charset=utf-8",
		strings.NewReader(string(body)))

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("resp error:", err)
	}

	decode(respBody)

	//fmt.Println(string(respBody))
}

func getBody(orderId, userId uint32, userName string) ([]byte, error) {
	req := &oidmc_moclient.MoClientNotifyOrderReq{
		OrderId: orderId,
		UserId: userId,
		UserName: userName,
		Sign: getSign(orderId, userId, userName),
	}

	data, err := proto.Marshal(req)
	if err != nil {
		fmt.Println("marshaling error: ", err)
	}

	return data, err
}

func getSign(orderId, userId uint32, userName string) string {
	salt := []string{"0WV56ndR", "AnuP558M", "3mosdrX8", "1Bm3lqLE"}

	h := md5.New()
	io.WriteString(h, salt[0] + strconv.Itoa(int(orderId)) +
		salt[1] + strconv.Itoa(int(userId)) +
		salt[2] + userName + salt[3])

	strSign := hex.EncodeToString(h.Sum(nil))
	return strSign
}

func decode(data []byte) error {
	res := &oidmc_moclient.MoClientNotifyOrderRes{}
	err := proto.Unmarshal(data, res)
	if err != nil {
		fmt.Println("Unmarshal error", err)
	}

	fmt.Println("res err: ", res.Err)
	fmt.Println("res err msg: ", res.ErrMsg)
	return err
}
