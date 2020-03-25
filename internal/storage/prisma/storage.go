package prisma

import (
	"os"
)

type Storage struct {
	Prisma *Client
}

func NewStorage() (*Storage, error) {
	s := new(Storage)
	c := New(&Options{
		Secret:   os.Getenv("PRISMA_SECRET"),
		Endpoint: "http://" + os.Getenv("PRISMA_HOST") + ":4466/dudutalk/" + os.Getenv("PRISMA_STAGE"),
	})
	s.Prisma = c
	return s, nil
}
