package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-demo/client/service"
	"log"
	"os"
)

func main() {
	//1. 新建连接，端口是服务端开放的8001端口
	//无证书
	//conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatal("连接服务端失败！", err)
	//}

	//有证书[单向认证]
	//creds, err2 := credentials.NewClientTLSFromFile("cert/server.pem", "*.xionglin.site")
	//if err2 != nil {
	//	log.Fatal("证书生成错误", err2)
	//}

	//有证书[双向认证]
	cert, err2 := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	if err2 != nil {
		log.Fatal("证书读取错误", err2)
	}
	//创建一个新的、空的CertPool
	certPool := x509.NewCertPool()
	ca, _ := os.ReadFile("cert/ca.crt")
	//尝试解析所传入的PEM编码的证书。如果解析成功会将其加到CertPool中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	//构建基于TLS的TransportCredentials选项
	creds := credentials.NewTLS(&tls.Config{
		//设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		ServerName:   "*.xionglin.site",
		RootCAs:      certPool,
	})

	//连接
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal("连接服务端失败！", err)
	}

	//退出时关闭链接
	defer conn.Close()

	client := service.NewProductServiceClient(conn)

	req := &service.Request{
		Id: 665,
	}
	res, err := client.GetProductStock(context.Background(), req)
	if err != nil {
		log.Fatal("查询库存出错", err)
	}

	log.Printf("查询库存成功【stock=%v】\n", res.Stock)

}
