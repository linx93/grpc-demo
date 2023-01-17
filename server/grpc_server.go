package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	service2 "grpc-demo/server/service"
	"log"
	"net"
	"os"
)

func main() {
	//没有证书---开始
	//grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	//没有证书---结束

	//creds, err2 := credentials.NewServerTLSFromFile("cert/server.pem", "cert/server.key")
	//if err2 != nil {
	//	log.Fatal("证书生成错误", err2)
	//}

	//grpcServer := grpc.NewServer(grpc.Creds(creds))
	//有证书[单向认证]---结束

	//有证书[双向认证]------开始
	//从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert, err2 := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	if err2 != nil {
		log.Fatal("证书读取错误", err2)
	}

	//创建一个新的、空的CertPool
	certPool := x509.NewCertPool()
	ca, err2 := os.ReadFile("cert/ca.crt")
	if err2 != nil {
		log.Fatal("ca证书读取错误", err2)
	}

	//尝试解析所传入的PEM编码的证书，如果解析成功会将其加到CertPool中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)

	//构建基于TLS的TransportCredentials选项
	creds := credentials.NewTLS(&tls.Config{
		//设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		//要求必须校验客户端证书。可以根据实际情况选用一下参数
		ClientAuth: tls.RequireAndVerifyClientCert,
		//设置跟证书的集合，校验方式使用ClientAuth中设定的模式
		ClientCAs: certPool,
	})

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	//有证书[双向认证]------结束

	service2.RegisterProductServiceServer(grpcServer, service2.ProductService)
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("启动监听出错", err)
	}

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal("启动服务出错", err)
	}

	fmt.Println("启动grpc服务端成功")

}
