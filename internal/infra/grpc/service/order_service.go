package service

import (
	"context"

	"github.com/felipezschornack/golang-cleanarch/internal/infra/grpc/pb"
	"github.com/felipezschornack/golang-cleanarch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrderUseCase usecase.ListOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrderUseCase:   listOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}

	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	order := &pb.Order{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}

	return &pb.CreateOrderResponse{
		Order: order,
	}, nil
}

func (s *OrderService) ListOrders(context.Context, *pb.Blank) (*pb.OrderListResponse, error) {

	orders, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var orderMessages []*pb.Order
	for _, order := range orders {
		orderMessage := &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}
		orderMessages = append(orderMessages, orderMessage)
	}
	return &pb.OrderListResponse{
		Orders: orderMessages,
	}, nil
}
