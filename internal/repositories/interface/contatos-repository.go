package repositories

import "github.com/gabrielmatsan/agenda/internal/models"

type ContatosRepository interface {
	AdicionarContato(contato *models.ContatoDetran) (*models.ContatoDetran, error)
	BuscarContatoPorID(id int) (*models.ContatoDetran, error)
	BuscarTodosContatos() ([]*models.ContatoDetran, error)
	AtualizarContato(id int, contato *models.ContatoDetran) (*models.ContatoDetran, error)
	DeletarContato(id int) error
	BuscarContatoPorNome(nome string) ([]*models.ContatoDetran, error)
}
