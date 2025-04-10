package main

import (
	"fmt"

	apb "protobufDemo/AppAccessRecordPb"

	"google.golang.org/protobuf/proto"
)

func genAppAccessRecord() []byte {
	pbRec := apb.AppAccessRecord{}

	pbRec.AppId = 1001
	pbRec.AppName = "rupin app"
	pbRec.UserName = "user1@rupin.in"
	pbRec.BytesRx = 50001
	pbRec.BytesTx = 20001

	buf, err := proto.Marshal(&pbRec)
	if err != nil {
		fmt.Println("Error in protobuf encoding!")
		return nil
	}
	return buf
}

func printAppAccessRecord(buf []byte) {
	p := apb.AppAccessRecord{}

	err := proto.Unmarshal(buf, &p)
	if err != nil {
		fmt.Println("Error in  protobuf decoding!")
		return
	}

	fmt.Println("App Access Record = ", p)
	fmt.Println("app id= ", p.GetAppId())
	fmt.Println("app name= ", p.GetAppName())
	fmt.Println("user name= ", p.GetUserName())
	fmt.Println("bytes rx= ", p.GetBytesRx())
	fmt.Println("bytes tx= ", p.GetBytesTx())
}

func main() {
	buf := genAppAccessRecord()
	printAppAccessRecord(buf)
}
