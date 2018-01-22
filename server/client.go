package server

import (
	"fmt"
	"net"
	"future/common"
	"bytes"
	"encoding/binary"
	"os"
	"crypto/md5"
)

type GameServer struct {
	PlayerID int
	Gold int
	Diamond int
}


func (g *GameServer) AddMoney() bool {

	server := common.GameServerAddress
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return false
	}

	fmt.Println("connect success")

	return send(conn,g)
}


//发送上分请求
func send(conn net.Conn,g *GameServer) bool {

	protocolID := IntToBytes(0x50)
	playerID := IntToBytes(g.PlayerID)
	addGold := Int64ToBytes(g.Gold)
	message := packet(protocolID, playerID, addGold)
	packageLength := IntToBytes(len(message) + 4 + 16)
	h := md5.New()
	h.Write(packageLength)
	h.Write(protocolID)
	h.Write(playerID)
	h.Write(addGold)
	h.Write([]byte("NzrWj")) // 需要加密的字符串为 sign
	serverkey := h.Sum(nil)
	fmt.Printf("%s\n", "abcdefg"[2:4])
	message = packet(message,serverkey[:])
	packageContent := append(packageLength, message...)
	conn.Write(packageContent)
	fmt.Printf("messagelength=%d,md5=%d,packagelen=%d \n", len(packageContent), len(serverkey[:]), BytesToInt(packageLength))
	//fmt.Printf("md5val=%s \n",serverkey)

	return receive(conn)
	//defer conn.Close()
}

func receive(conn net.Conn) bool {
	//缓冲区，存储被截断的数据
	tmpBuffer := make([]byte, 0)
	buffer := make([]byte, 1024)
	for{
		n, err := conn.Read(buffer)
		if err != nil{
			fmt.Println(conn.RemoteAddr().String(), "connection error: ", err)
			return false
		}
		tmpBuffer = buffer[:n]
		if unpack(tmpBuffer) {
			return true
		}
	}
	defer conn.Close()
	return false
}

func unpack(data []byte) bool{
	fmt.Println("::客户端接收到的bufs缓冲器内容::")
	head := make([]byte, 4)
	protocolID := make([]byte, 4)
	msgID := make([]byte, 4)
	bufs := bytes.NewBuffer(data)
	bufs.Read(head)
	bufs.Read(protocolID)
	bufs.Read(msgID)
	packagelength := BytesToInt(head)
	protocol := BytesToInt(protocolID)
	msg := BytesToInt(msgID)
	fmt.Printf("包长度=%d,协议=%d,返回值=%d,函数返回结果=%s",packagelength,protocol,msg,protocol==-0x50)
	return msg==0
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}

//整形转换成字节
func Int64ToBytes(n int) []byte {
	x := int64(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x) //这里有大端序和小端序的问题,不同的编程语言和操作系统对字节的处理略有不同，网络通信中一般用小端序 http://lihaoquan.me/2016/11/5/golang-byteorder.html
	return bytesBuffer.Bytes()
}

func BytesToInt(buf []byte) int32 {
	//return int32(binary.BigEndian.Uint32(buf))
	return int32(binary.LittleEndian.Uint32(buf))
}

func packet(message ...[]byte) []byte {
	return bytes.Join(message, []byte(""))
}
