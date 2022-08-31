package services

import (
	"fmt"

	"github.com/rasul07/books_api_gateway/config"
	"github.com/rasul07/books_api_gateway/genproto/book"
	"github.com/rasul07/books_api_gateway/genproto/category"
	"google.golang.org/grpc"
)

type ServiceManager interface {
	BookService() book.BookServiceClient
	CategoryService() category.CategoryServiceClient
}

type grpcClients struct {
	bookService book.BookServiceClient
	bookCategoryService  category.CategoryServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	connBookService, err := grpc.Dial(
		fmt.Sprintf("%v", conf.RPCPort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	connBookCategoryService, err := grpc.Dial(
		fmt.Sprintf("%v", conf.RPCPort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		bookService: book.NewBookServiceClient(connBookService),
		bookCategoryService:  category.NewCategoryServiceClient(connBookCategoryService),
	}, nil
}

func (g *grpcClients) BookService() book.BookServiceClient {
	return g.bookService
}

func (g *grpcClients) CategoryService() category.CategoryServiceClient {
	return g.bookCategoryService
}
