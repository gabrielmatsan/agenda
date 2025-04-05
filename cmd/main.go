package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gabrielmatsan/agenda/internal/models"
	inmemory "github.com/gabrielmatsan/agenda/internal/repositories/in-memory"
	repositories "github.com/gabrielmatsan/agenda/internal/repositories/interface"
)

func main() {
	// Inicializa o repositório
	repo := inmemory.NewInMemoryContatosRepository()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		exibirMenu()
		opcao := lerOpcao(scanner)

		switch opcao {
		case 1:
			adicionarContato(scanner, repo)
		case 2:
			buscarContatoPorID(scanner, repo)
		case 3:
			listarTodosContatos(repo)
		case 4:
			atualizarContato(scanner, repo)
		case 5:
			deletarContato(scanner, repo)
		case 6:
			buscarContatoPorNome(scanner, repo)
		case 0:
			fmt.Println("\nSaindo do programa. Até logo!")
			return
		default:
			fmt.Println("\nOpção inválida! Por favor, tente novamente.")
		}

		fmt.Println("\nPressione ENTER para continuar...")
		scanner.Scan()
	}
}

func exibirMenu() {
	fmt.Println("\n=== AGENDA DE CONTATOS ===")
	fmt.Println("1. Adicionar contato")
	fmt.Println("2. Buscar contato por ID")
	fmt.Println("3. Listar todos os contatos")
	fmt.Println("4. Atualizar contato")
	fmt.Println("5. Deletar contato")
	fmt.Println("6. Buscar contato por nome")
	fmt.Println("0. Sair")
	fmt.Print("\nEscolha uma opção: ")
}

func lerOpcao(scanner *bufio.Scanner) int {
	scanner.Scan()
	opcao, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return -1 // Retorna opção inválida em caso de erro
	}
	return opcao
}

func adicionarContato(scanner *bufio.Scanner, repo repositories.ContatosRepository) {
	contato := models.ContatoDetran{}

	fmt.Println("\n=== ADICIONAR NOVO CONTATO ===")

	fmt.Print("Nome: ")
	scanner.Scan()
	contato.Nome = scanner.Text()

	fmt.Print("Telefone: ")
	scanner.Scan()
	contato.Telefone = scanner.Text()

	fmt.Print("Email: ")
	scanner.Scan()
	contato.Email = scanner.Text()

	fmt.Print("Data de Nascimento (DD/MM/AAAA): ")
	scanner.Scan()
	contato.DataNascimento = scanner.Text()

	fmt.Print("CPF: ")
	scanner.Scan()
	contato.CpfOrCNPJ = scanner.Text()

	fmt.Print("CEP: ")
	scanner.Scan()
	contato.Cep = scanner.Text()

	fmt.Print("Placa de Carro: ")
	scanner.Scan()
	contato.PlacaDeCarro = scanner.Text()

	fmt.Print("Senha: ")
	scanner.Scan()
	contato.Senha = scanner.Text()

	novoContato, err := repo.AdicionarContato(&contato)
	if err != nil {
		fmt.Printf("\nErro ao adicionar contato: %v\n", err)
		return
	}

	fmt.Printf("\nContato adicionado com sucesso! ID: %d\n", novoContato.ID)
}

func buscarContatoPorID(scanner *bufio.Scanner, repo repositories.ContatosRepository) {
	fmt.Println("\n=== BUSCAR CONTATO POR ID ===")
	fmt.Print("Digite o ID do contato: ")
	scanner.Scan()

	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("ID inválido!")
		return
	}

	contato, err := repo.BuscarContatoPorID(id)
	if err != nil {
		fmt.Printf("Erro ao buscar contato: %v\n", err)
		return
	}

	if contato == nil {
		fmt.Println("Contato não encontrado!")
		return
	}

	exibirContato(contato)
}

func listarTodosContatos(repo repositories.ContatosRepository) {
	fmt.Println("\n=== LISTA DE CONTATOS ===")

	contatos, err := repo.BuscarTodosContatos()
	if err != nil {
		fmt.Printf("Erro ao listar contatos: %v\n", err)
		return
	}

	if len(contatos) == 0 {
		fmt.Println("Nenhum contato cadastrado!")
		return
	}

	for _, contato := range contatos {
		fmt.Printf("ID: %d | Nome: %s | Telefone: %s\n",
			contato.ID, contato.Nome, contato.Telefone)
	}

	fmt.Printf("\nTotal: %d contato(s)\n", len(contatos))
}

func atualizarContato(scanner *bufio.Scanner, repo repositories.ContatosRepository) {
	fmt.Println("\n=== ATUALIZAR CONTATO ===")
	fmt.Print("Digite o ID do contato a ser atualizado: ")
	scanner.Scan()

	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("ID inválido!")
		return
	}

	// Busca o contato atual
	contatoAtual, err := repo.BuscarContatoPorID(id)
	if err != nil || contatoAtual == nil {
		fmt.Println("Contato não encontrado!")
		return
	}

	// Cria uma cópia do contato atual
	novoContato := *contatoAtual

	fmt.Println("\nDigite os novos dados (ou deixe em branco para manter o valor atual):")

	fmt.Printf("Nome [%s]: ", novoContato.Nome)
	scanner.Scan()
	if texto := scanner.Text(); texto != "" {
		novoContato.Nome = texto
	}

	fmt.Printf("Telefone [%s]: ", novoContato.Telefone)
	scanner.Scan()
	if texto := scanner.Text(); texto != "" {
		novoContato.Telefone = texto
	}

	fmt.Printf("Email [%s]: ", novoContato.Email)
	scanner.Scan()
	if texto := scanner.Text(); texto != "" {
		novoContato.Email = texto
	}

	fmt.Printf("Data de Nascimento [%s]: ", novoContato.DataNascimento)
	scanner.Scan()
	if texto := scanner.Text(); texto != "" {
		novoContato.DataNascimento = texto
	}

	fmt.Printf("CPF [%s]: ", novoContato.CpfOrCNPJ)
	scanner.Scan()
	if texto := scanner.Text(); texto != "" {
		novoContato.CpfOrCNPJ = texto
	}

	fmt.Printf("CEP [%s]: ", novoContato.Cep)
	scanner.Scan()
	if texto := scanner.Text(); texto != "" {
		novoContato.Cep = texto
	}

	fmt.Printf("Placa de Carro [%s]: ", novoContato.PlacaDeCarro)
	scanner.Scan()
	if texto := scanner.Text(); texto != "" {
		novoContato.PlacaDeCarro = texto
	}

	fmt.Print("Nova Senha (deixe em branco para manter a atual): ")
	scanner.Scan()
	if texto := scanner.Text(); texto != "" {
		novoContato.Senha = texto
	}

	// Atualiza o contato
	contatoAtualizado, err := repo.AtualizarContato(id, &novoContato)
	if err != nil {
		fmt.Printf("Erro ao atualizar contato: %v\n", err)
		return
	}

	fmt.Println("\nContato atualizado com sucesso!")
	exibirContato(contatoAtualizado)
}

func deletarContato(scanner *bufio.Scanner, repo repositories.ContatosRepository) {
	fmt.Println("\n=== DELETAR CONTATO ===")
	fmt.Print("Digite o ID do contato a ser deletado: ")
	scanner.Scan()

	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("ID inválido!")
		return
	}

	// Confirma a exclusão
	fmt.Print("Tem certeza que deseja deletar este contato? (s/n): ")
	scanner.Scan()
	confirmacao := strings.ToLower(scanner.Text())

	if confirmacao != "s" && confirmacao != "sim" {
		fmt.Println("Operação cancelada!")
		return
	}

	err = repo.DeletarContato(id)
	if err != nil {
		fmt.Printf("Erro ao deletar contato: %v\n", err)
		return
	}

	fmt.Println("Contato deletado com sucesso!")
}

func buscarContatoPorNome(scanner *bufio.Scanner, repo repositories.ContatosRepository) {
	fmt.Println("\n=== BUSCAR CONTATO POR NOME ===")
	fmt.Print("Digite o nome (ou parte do nome) para buscar: ")
	scanner.Scan()

	nome := scanner.Text()
	if nome == "" {
		fmt.Println("Nome de busca inválido!")
		return
	}

	contatos, err := repo.BuscarContatoPorNome(nome)
	if err != nil {
		fmt.Printf("Erro ao buscar contatos: %v\n", err)
		return
	}

	if len(contatos) == 0 {
		fmt.Println("Nenhum contato encontrado!")
		return
	}

	fmt.Printf("\nEncontrados %d contato(s):\n", len(contatos))
	for _, contato := range contatos {
		fmt.Printf("ID: %d | Nome: %s | Telefone: %s\n",
			contato.ID, contato.Nome, contato.Telefone)
	}
}

func exibirContato(contato *models.ContatoDetran) {
	fmt.Println("\n=== DETALHES DO CONTATO ===")
	fmt.Printf("ID: %d\n", contato.ID)
	fmt.Printf("Nome: %s\n", contato.Nome)
	fmt.Printf("Telefone: %s\n", contato.Telefone)
	fmt.Printf("Email: %s\n", contato.Email)
	fmt.Printf("Data de Nascimento: %s\n", contato.DataNascimento)
	fmt.Printf("CPF: %s\n", contato.CpfOrCNPJ)
	fmt.Printf("CEP: %s\n", contato.Cep)
	fmt.Printf("Placa de Carro: %s\n", contato.PlacaDeCarro)
	fmt.Printf("Senha: ********\n") // Não exibe a senha real
}
