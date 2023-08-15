package dto

import "github.com/google/uuid"

type PersonRequestDTO struct {
	Name      string   `json:"nome"`
	Nickname  string   `json:"apelido"`
	Birthdate string   `json:"nascimento"`
	Stack     []string `json:"stack"`
}

type PersonResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"nome"`
	Nickname  string    `json:"apelido"`
	Birthdate string    `json:"nascimento"`
	Stack     []string  `json:"stack"`
}
