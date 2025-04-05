package inmemory

import (
	"errors"
	"strings"
	"sync"

	"github.com/gabrielmatsan/agenda/internal/models"
)

var (
	ErrContatoNil           = errors.New("contato não pode ser nulo")
	ErrContatoNaoEncontrado = errors.New("contato não encontrado")
)

type InMemoryContatosRepository struct {
	contatos []*models.ContatoDetran
	mutex    sync.RWMutex
}

func NewInMemoryContatosRepository() *InMemoryContatosRepository {
	return &InMemoryContatosRepository{
		contatos: []*models.ContatoDetran{},
	}
}

func (r *InMemoryContatosRepository) AdicionarContato(contato *models.ContatoDetran) (*models.ContatoDetran, error) {

	if contato == nil {
		return nil, ErrContatoNil
	}

	if _, err := models.ValidarContato(contato); err != nil {
		return nil, err
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	novoContato := new(models.ContatoDetran)
	*novoContato = *contato

	r.contatos = append(r.contatos, novoContato)

	return novoContato, nil
}

func (r *InMemoryContatosRepository) BuscarContatoPorID(id int) (*models.ContatoDetran, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, contato := range r.contatos {
		if contato.ID == id {

			// Retorna uma cópia defensiva do contato
			resultado := *contato

			return &resultado, nil
		}
	}
	return nil, ErrContatoNaoEncontrado
}

func (r *InMemoryContatosRepository) BuscarTodosContatos() ([]*models.ContatoDetran, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return r.contatos, nil
}

func (r *InMemoryContatosRepository) AtualizarContato(id int, contato *models.ContatoDetran) (*models.ContatoDetran, error) {

	if contato == nil {
		return nil, ErrContatoNil
	}

	if _, err := models.ValidarContato(contato); err != nil {
		return nil, err
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, c := range r.contatos {
		if c.ID == id {

			novoContato := new(models.ContatoDetran)
			*novoContato = *contato
			novoContato.ID = id

			r.contatos[i] = novoContato
			// Retorna uma cópia defensiva do contato atualizado
			resultado := *novoContato
			return &resultado, nil
		}
	}
	return nil, errors.New("contato não encontrado")
}

func (r *InMemoryContatosRepository) DeletarContato(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, c := range r.contatos {
		if c.ID == id {
			r.contatos = append(r.contatos[:i], r.contatos[i+1:]...)
			return nil
		}
	}

	return errors.New("contato não encontrado")
}

func (r *InMemoryContatosRepository) BuscarContatoPorNome(nome string) ([]*models.ContatoDetran, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var contatosEncontrados []*models.ContatoDetran

	for _, contato := range r.contatos {
		if strings.EqualFold(contato.Nome, nome) {
			copiaContato := *contato
			contatosEncontrados = append(contatosEncontrados, &copiaContato)
		}
	}

	if len(contatosEncontrados) == 0 {
		return nil, ErrContatoNaoEncontrado
	}

	return contatosEncontrados, nil
}
