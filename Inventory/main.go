package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "net/http/pprof"

	"github.com/IsmaelNicolas/Library_System/author"
	"github.com/IsmaelNicolas/Library_System/book"
	"github.com/IsmaelNicolas/Library_System/config"
	"github.com/IsmaelNicolas/Library_System/models"
	"github.com/IsmaelNicolas/Library_System/stand"
	"github.com/IsmaelNicolas/Library_System/stock"
	"github.com/IsmaelNicolas/Library_System/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type BookServer struct{}

// AuthorServer implements the AuthorServiceServer interface
type AuthorServer struct{}

// StandServer implements the StandServiceServer interface
type StandServer struct{}

// StockServer implements the StockServiceServer interface
type StockServer struct{}

// Implementación de los métodos de StockService
func (s *StockServer) CreateStock(ctx context.Context, req *stock.CreateStockRequest) (*stock.StockResponse, error) {
	// Implementación de la lógica para crear un stock
	return &stock.StockResponse{}, nil
}

func (s *StockServer) ReadStocks(ctx context.Context, req *stock.ReadStocksRequest) (*stock.ReadStocksResponse, error) {
	// Implementación de la lógica para leer stocks
	return &stock.ReadStocksResponse{}, nil
}

// Implementación de los métodos de StandService
func (s *StandServer) CreateStand(ctx context.Context, req *stand.CreateStandRequest) (*stand.StandResponse, error) {
	// Lógica para crear un stand
	return &stand.StandResponse{}, nil
}

func (s *StandServer) ReadStands(ctx context.Context, req *stand.ReadStandsRequest) (*stand.ReadStandsResponse, error) {
	// Lógica para leer stands
	return &stand.ReadStandsResponse{}, nil
}

// Implementación de los métodos de AuthorService
func (s *AuthorServer) CreateAuthor(ctx context.Context, req *author.CreateAuthorRequest) (*author.AuthorResponse, error) {
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Crear el autor en la base de datos utilizando la función CreateAuthor del paquete database
	newAuthor, err := utils.CreateAuthor(db, req.FullName)
	if err != nil {
		return nil, err
	}

	// Preparar la respuesta del servidor con el autor recién creado
	response := &author.AuthorResponse{
		Author: &author.Author{
			Id:       uint64(newAuthor.ID),
			FullName: newAuthor.FullName,
		},
	}

	return response, nil
}

func (s *AuthorServer) ReadAuthors(ctx context.Context, req *author.ReadAuthorsRequest) (*author.ReadAuthorsResponse, error) {
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Leer todos los autores de la base de datos utilizando la función ReadAuthors del paquete database
	authors, err := utils.ReadAuthors(db)
	if err != nil {
		return nil, err
	}

	// Preparar la respuesta del servidor con los autores leídos
	var authorResponses []*author.Author
	for _, a := range authors {
		authorResponses = append(authorResponses, &author.Author{
			Id:       uint64(a.ID),
			FullName: a.FullName,
		})
	}

	response := &author.ReadAuthorsResponse{
		Authors: authorResponses,
	}

	return response, nil
}

func (s *BookServer) CreateBook(ctx context.Context, req *book.CreateBookRequest) (*book.BookResponse, error) {
	// Lógica para crear un libro
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	newBook, err := utils.CreateBook(db, req.Title, req.Year, uint(req.Author.Id))
	if err != nil {
		return nil, err
	}

	author := &author.Author{
		Id:       req.Author.Id,
		FullName: req.Author.FullName, // Suponiendo que Author tiene un campo FullName
	}

	response := &book.BookResponse{
		Book: &book.Book{
			Id:     uint64(newBook.ID),
			Title:  newBook.Title,
			Year:   newBook.Year,
			Author: author, // Usar el campo Author para incluir la información del autor
		},
	}

	return response, nil
}

func (s *BookServer) ReadBooks(ctx context.Context, req *book.ReadBooksRequest) (*book.ReadBooksResponse, error) {
	// Lógica para leer libros
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Leer libros de la base de datos con la información del autor pre-cargada
	var books []*models.Book
	if err := db.Preload("Author").Find(&books).Error; err != nil {
		return nil, err
	}

	// Crear la respuesta con los libros leídos
	var bookResponses []*book.Book
	for _, b := range books {
		// Crear un objeto book.Book para cada libro leído
		bookResponse := &book.Book{
			Id:    uint64(b.ID),
			Title: b.Title,
			Year:  b.Year,
			// Incluir el nombre completo del autor en la respuesta
			Author: &author.Author{
				Id:       uint64(b.Author.ID),
				FullName: b.Author.FullName,
			},
		}
		bookResponses = append(bookResponses, bookResponse)
	}

	response := &book.ReadBooksResponse{
		Books: bookResponses,
	}

	return response, nil
}
func (s *BookServer) ReadBookById(ctx context.Context, req *book.ReadBookByIdRequest) (*book.BookResponse, error) {
	// Lógica para leer un libro por su ID
	return &book.BookResponse{}, nil
}

func (s *BookServer) UpdateBookById(ctx context.Context, req *book.UpdateBookByIdRequest) (*book.BookResponse, error) {
	// Lógica para actualizar un libro por su ID
	return &book.BookResponse{}, nil
}

func main() {
	// Conexión a la base de datos
	db, err := config.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Auto migración de modelos
	if err := db.AutoMigrate(&models.Author{}, &models.Book{}, &models.Stand{}, &models.Stock{}).Error; err != nil {
		log.Fatalf("failed to AutoMigrate: %v", err)
	}
	log.Printf("Database OK")

	// Opciones del servidor gRPC
	serverOptions := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Second * 60,
			Time:              time.Second * 10,
			Timeout:           time.Second * 20,
		}),
		grpc.ConnectionTimeout(time.Second * 10),
	}

	// Crear listener
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Crear servidor gRPC
	s := grpc.NewServer(serverOptions...)

	// Registrar servicios gRPC
	stock.RegisterStockServiceServer(s, &StockServer{})
	stand.RegisterStandServiceServer(s, &StandServer{})
	author.RegisterAuthorServiceServer(s, &AuthorServer{})
	book.RegisterBookServiceServer(s, &BookServer{})

	// Iniciar servidor gRPC
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Mensaje de éxito
	log.Println("SERVER OK")

	// Esperar señal de terminación
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	log.Println("Stopping server....")
	s.GracefulStop()
	log.Println("Bye")
}
