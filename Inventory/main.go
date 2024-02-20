package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/IsmaelNicolas/Library_System/books" // Replace this with the actual package name
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBookServiceServer
}

var books []*pb.Book

func (s *server) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.BookResponse, error) {
	book := &pb.Book{
		Id:       uint64(len(books) + 1),
		Title:    req.Title,
		Year:     req.Year,
		AuthorId: req.AuthorId,
	}
	books = append(books, book)
	return &pb.BookResponse{Book: book}, nil
}

func (s *server) ReadBooks(ctx context.Context, req *pb.ReadBooksRequest) (*pb.ReadBooksResponse, error) {
	return &pb.ReadBooksResponse{Books: books}, nil
}

func (s *server) ReadBookById(ctx context.Context, req *pb.ReadBookByIdRequest) (*pb.BookResponse, error) {
	for _, book := range books {
		if book.Id == req.Id {
			return &pb.BookResponse{Book: book}, nil
		}
	}
	return nil, fmt.Errorf("book with id %d not found", req.Id)
}

func (s *server) UpdateBookById(ctx context.Context, req *pb.UpdateBookByIdRequest) (*pb.BookResponse, error) {
	for _, book := range books {
		if book.Id == req.Id {
			book.Title = req.Title
			book.Year = req.Year
			book.AuthorId = req.AuthorId
			return &pb.BookResponse{Book: book}, nil
		}
	}
	return nil, fmt.Errorf("book with id %d not found", req.Id)
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBookServiceServer(s, &server{})
	log.Println("Server started on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
