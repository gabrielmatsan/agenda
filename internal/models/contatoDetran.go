package models

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	// Expressões regulares
	regexNome           = regexp.MustCompile(`^[A-ZÀ-ÖØ-ÿ][a-zà-öø-ÿ]+([ ][A-ZÀ-ÖØ-ÿ][a-zà-öø-ÿ]+)*$`)
	regexTelefone       = regexp.MustCompile(`^(?:\([0-9]{2}\)|[0-9]{2})?[0-9]{5}-[0-9]{4}$`)
	regexEmail          = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	regexDataNascimento = regexp.MustCompile(`^(0[0-9]|[12][0-9]|3[01])/(0[0-9]|1[0-2])/\d{4}$`)
	regexCpfOrCnpj      = regexp.MustCompile(`([0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2})|([0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2})`)
	regexCEP            = regexp.MustCompile(`([0-9]{5}[-]?[0-9]{3})|([0-9]{8})`)
	regexPlacaDeCarro   = regexp.MustCompile(`^[A-Z]{3}[0-9]{4}$|^[A-Z]{3}-[0-9]{4}$|^[A-Z]{3}[0-9][A-Z][0-9]{2}$`)
)

// Erros comuns
var (
	ErrContatoNil           = errors.New("contato não pode ser nulo")
	ErrContatoNaoEncontrado = errors.New("contato não encontrado")

	// Erros específicos de validação
	ErrNomeInvalido           = errors.New("nome inválido: deve começar com letra maiúscula e não pode conter números ou caracteres especiais")
	ErrTelefoneInvalido       = errors.New("telefone inválido: deve estar no formato (XX) XXXXX-XXXX ou XX XXXXX-XXXX")
	ErrEmailInvalido          = errors.New("email inválido: verifique o formato do email")
	ErrDataNascimentoInvalida = errors.New("data de nascimento inválida: o formato deve ser DD/MM/AAAA")
	ErrCpfCnpjInvalido        = errors.New("CPF ou CNPJ inválido: formato incorreto")
	ErrCepInvalido            = errors.New("CEP inválido: deve estar no formato XXXXX-XXX ou XXXXXXXX")
	ErrPlacaCarroInvalida     = errors.New("placa de carro inválida: deve estar no formato ABC1234, ABC-1234 ou ABC1D23")
	ErrSenhaInvalida          = errors.New("senha inválida: deve conter pelo menos 8 caracteres, incluindo letras maiúsculas, minúsculas, números e caracteres especiais ($*&@#)")
)

var staticID = 0

type ContatoDetran struct {
	ID             int    `json:"id"`
	Nome           string `json:"nome"`
	Telefone       string `json:"telefone"`
	Email          string `json:"email"`
	DataNascimento string `json:"data_nascimento"`
	CpfOrCNPJ      string `json:"cpf"`
	Cep            string `json:"cep"`
	PlacaDeCarro   string `json:"placa_de_carro"`
	Senha          string `json:"senha"`
}

var ErrCamposInvalidos = errors.New("campos inválidos")

func ValidarContato(contato *ContatoDetran) (map[string]error, error) {
	if contato == nil {
		return nil, ErrContatoNil
	}

	errorsMap := make(map[string]error)

	if err := ValidarNome(contato.Nome); err != nil {
		errorsMap["nome"] = err
	}
	if err := ValidarTelefone(contato.Telefone); err != nil {
		errorsMap["telefone"] = err
	}
	if err := ValidarEmail(contato.Email); err != nil {
		errorsMap["email"] = err
	}
	if err := ValidarDataNascimento(contato.DataNascimento); err != nil {
		errorsMap["data_nascimento"] = err
	}
	if err := ValidarCpfOrCnpj(contato.CpfOrCNPJ); err != nil {
		errorsMap["cpfOrCnpj"] = err
	}
	if err := ValidarCEP(contato.Cep); err != nil {
		errorsMap["cep"] = err
	}
	if err := ValidarPlacaDeCarro(contato.PlacaDeCarro); err != nil {
		errorsMap["placa_de_carro"] = err
	}
	if err := ValidarSenha(contato.Senha); err != nil {
		errorsMap["senha"] = err
	}

	if len(errorsMap) > 0 {
		var campos []string
		for campo := range errorsMap {
			campos = append(campos, campo)
		}
		msg := fmt.Sprintf("Erro nos campos: %s", strings.Join(campos, ", "))
		return errorsMap, fmt.Errorf("%w: %s", ErrCamposInvalidos, msg)
	}

	return nil, nil
}

func NewContatoDetran(contato ContatoDetran) *ContatoDetran {
	return &ContatoDetran{
		ID:             GenerateID(),
		Nome:           contato.Nome,
		Telefone:       contato.Telefone,
		Email:          contato.Email,
		DataNascimento: contato.DataNascimento,
		CpfOrCNPJ:      contato.CpfOrCNPJ,
		Cep:            contato.Cep,
		PlacaDeCarro:   contato.PlacaDeCarro,
		Senha:          contato.Senha,
	}
}

func GenerateID() int {
	staticID++
	return staticID
}

func ValidarNome(nome string) error {
	if !regexNome.MatchString(nome) {
		return ErrNomeInvalido
	}
	return nil
}

func ValidarTelefone(telefone string) error {
	if !regexTelefone.MatchString(telefone) {
		return ErrTelefoneInvalido
	}
	return nil
}

func ValidarEmail(email string) error {
	if !regexEmail.MatchString(email) {
		return ErrEmailInvalido
	}
	return nil
}

func ValidarDataNascimento(dataNascimento string) error {
	if !regexDataNascimento.MatchString(dataNascimento) {
		return ErrDataNascimentoInvalida
	}
	return nil
}

func ValidarCpfOrCnpj(cpfOrCnpj string) error {
	if !regexCpfOrCnpj.MatchString(cpfOrCnpj) {
		return ErrCpfCnpjInvalido
	}
	return nil
}

func ValidarCEP(cep string) error {
	if !regexCEP.MatchString(cep) {
		return ErrCepInvalido
	}
	return nil
}

func ValidarPlacaDeCarro(placaDeCarro string) error {
	if !regexPlacaDeCarro.MatchString(placaDeCarro) {
		return ErrPlacaCarroInvalida
	}
	return nil
}

func ValidarSenha(senha string) error {
	if len(senha) < 8 {
		return errors.New("senha inválida: deve ter pelo menos 8 caracteres")
	}

	temNumero := regexp.MustCompile(`[0-9]`).MatchString(senha)
	if !temNumero {
		return errors.New("senha inválida: deve conter pelo menos um número")
	}

	temLetraMaiuscula := regexp.MustCompile(`[A-Z]`).MatchString(senha)
	if !temLetraMaiuscula {
		return errors.New("senha inválida: deve conter pelo menos uma letra maiúscula")
	}

	temLetraMinuscula := regexp.MustCompile(`[a-z]`).MatchString(senha)
	if !temLetraMinuscula {
		return errors.New("senha inválida: deve conter pelo menos uma letra minúscula")
	}

	temEspecial := regexp.MustCompile(`[!$*&@#]`).MatchString(senha)
	if !temEspecial {
		return errors.New("senha inválida: deve conter pelo menos um caractere especial ($*&@#)")
	}

	somentePermitidos := regexp.MustCompile(`^[0-9a-zA-Z!$*&@#]+$`).MatchString(senha)
	if !somentePermitidos {
		return errors.New("senha inválida: contém caracteres não permitidos")
	}

	return nil
}
