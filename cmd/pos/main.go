package main

import (
	"context"
	"log"
	"net"

	"github.com/automatedtomato/order-management-system/pkg/api/pos" // go_packageで指定したパス
	"google.golang.org/grpc"
)

// posServiceServerはPOSServiceServerインターフェースを実装するサーバ。
type posServiceServer struct {
	pos.UnimplementedPOSServiceServer // 互換性維持のため埋め込む
}

// CreateOrderはCreateOrder RPCメソッドの実装。
func (c *posServiceServer) CreateOrder(ctx context.Context, req *pos.CreateOrderRequest) (*pos.CreateOrderResponse, error) {
	log.Printf("CreateOrderリクエストを受信しました: %v", req)

	// TODO: 注文処理のビジネスロジックを実装する

	// 仮のリスポンスを返す
	res := &pos.CreateOrderResponse{
		OrderId: "order-id-123",
		Status:  "PENDING",
	}
	return res, nil
}

func main() {
	// 8080ポートでリスンするTCPソケットを作成
	lis, err := net.Listen("tcp", "8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPCサーバを作成
	s := grpc.NewServer()

	// POSServiceをgRPCサーバに登録
	pos.RegisterPOSServiceServer(s, &posServiceServer{})

	log.Println("gRPCサーバをポート8080で起動しています...")

	// gRPCサーバを起動してリクエストの受付を開始
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
