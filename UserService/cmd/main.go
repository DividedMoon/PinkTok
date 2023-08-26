package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"math/big"
	"middleware/auth"
	"middleware/msgno"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user_service/biz/handler"
	biz "user_service/biz/userservice"
	"user_service/internal/model"
)

const (
	certFile = "certificate.pem"
	keyFile  = "private_key.pem"
)

func main() {
	// 创建一个通道来接收中断信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:11011")
	if err != nil {
		hlog.Errorf("resolve tcp addr failed, err:%v", err)
	}
	model.InitDB()

	servername := "user_service"
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(servername),
		provider.WithExportEndpoint("106.54.208.133:4317"),
		provider.WithEnableMetrics(false),
		provider.WithEnableTracing(true),
		provider.WithInsecure())
	defer p.Shutdown(context.Background())
	svr := biz.NewServer(new(handler.UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithMiddleware(msgno.MsgNoMiddleware),
		server.WithMiddleware(auth.AuthenticateServer),
		server.WithMuxTransport(),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: servername,
		}),
	)

	go func() {
		err = svr.Run()
		if err != nil {
			hlog.Errorf(err.Error())
		}
	}()

	<-interrupt
	hlog.Info("接收到中断信号,关闭服务器中")
	err = svr.Stop()
	if err != nil {
		hlog.Errorf(err.Error())
	}
	hlog.Info("服务器已关闭")
}

func getTlsConfig() (tlsConfig *tls.Config, err error) {
	// 检查证书文件是否存在
	hlog.Infof("check certificate files")
	certExist, keyExist := checkCertificateFiles()

	if certExist && keyExist {
		// 加载现有证书和私钥
		hlog.Infof("load certificate and key")
		cert, key, err := loadCertificate()
		if err != nil {
			fmt.Println("Failed to load certificate and key:", err)
			return nil, err
		}

		// 创建 TLS 配置
		tlsCert := tls.Certificate{
			Certificate: [][]byte{cert.Raw},
			PrivateKey:  key,
		}

		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
		}
	} else {
		// 生成自签名证书
		hlog.Infof("generate certificate")
		cert, key, err := generateCertificate()
		if err != nil {
			fmt.Println("Failed to generate certificate:", err)
			return nil, err
		}

		// 将证书和私钥保存到文件
		err = saveCertificate(cert, key)
		if err != nil {
			fmt.Println("Failed to save certificate and key:", err)
			return nil, err
		}

		// 创建 TLS 配置
		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{{Certificate: [][]byte{cert.Raw}, PrivateKey: key}},
		}
	}
	return tlsConfig, nil
}

func checkCertificateFiles() (bool, bool) {
	_, certErr := os.Stat(certFile)
	_, keyErr := os.Stat(keyFile)

	return !os.IsNotExist(certErr), !os.IsNotExist(keyErr)
}

func loadCertificate() (*x509.Certificate, *rsa.PrivateKey, error) {
	// 加载证书文件
	certPEM, err := os.ReadFile(certFile)
	if err != nil {
		return nil, nil, err
	}

	// 解码 PEM 格式证书
	certBlock, _ := pem.Decode(certPEM)
	if certBlock == nil {
		return nil, nil, fmt.Errorf("failed to decode certificate PEM")
	}

	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}

	// 加载私钥文件
	keyPEM, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, nil, err
	}

	// 解码 PEM 格式私钥
	keyBlock, _ := pem.Decode(keyPEM)
	if keyBlock == nil {
		return nil, nil, fmt.Errorf("failed to decode private key PEM")
	}

	key, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		return nil, nil, err
	}

	return cert, key, nil
}

func saveCertificate(cert *x509.Certificate, key *rsa.PrivateKey) error {
	// 将证书保存到文件
	certOut, err := os.Create(certFile)
	if err != nil {
		return err
	}
	defer certOut.Close()

	err = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw})
	if err != nil {
		return err
	}

	// 将私钥保存到文件
	keyOut, err := os.OpenFile(keyFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer keyOut.Close()

	err = pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	if err != nil {
		return err
	}

	return nil
}

func generateCertificate() (*x509.Certificate, *rsa.PrivateKey, error) {
	// 生成 RSA 密钥
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// 创建证书模板
	template := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "localhost"},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0), // 有效期为一年
		BasicConstraintsValid: true,
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")}, // 设置 IP 地址
	}

	// 使用自签名生成证书
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		return nil, nil, err
	}

	cert, err := x509.ParseCertificate(derBytes)
	if err != nil {
		return nil, nil, err
	}

	return cert, key, nil
}
