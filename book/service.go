package book

import "fmt"

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookReqeuest BookRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookReqeuest BookRequest) (Book, error) {
	price, err := bookReqeuest.Price.Int64()
	rating, err := bookReqeuest.Rating.Int64()
	discount, err := bookReqeuest.Rating.Int64()
	if err != nil {
		fmt.Println("Error parsing price")
	}
	book := Book{
		Title:    bookReqeuest.Title,
		Price:    int(price),
		Rating:   int(rating),
		Discount: int(discount),
	}

	newBook, err := s.repository.Create(book)

	return newBook, err
}
