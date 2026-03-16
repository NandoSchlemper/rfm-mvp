package auth

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type IService interface {
	Login() (*AuthResponse, error)
}

type service struct{}

// Login implements [IService].
func (s *service) Login() (*AuthResponse, error) {
	url := url.Values{}
	url.Add("cod_cliente", os.Getenv("DARWIN_COD_CLIENTE"))
	url.Add("login", os.Getenv("DARWIN_LOGIN"))
	url.Add("senha", os.Getenv("DARWIN_PASSWORD"))
	baseUrl := fmt.Sprintf("%s%s", os.Getenv("URL_DARWIN_API"), "/index")
	fullUrl := fmt.Sprintf("%s?%s", baseUrl, url.Encode())

	request, err := http.NewRequest("POST", fullUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("Erro ao montar a requisição.")
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Erro ao realizar a requisição POST.")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Erro ao ler o corpo da resposta.")
	}

	token := strings.Trim(string(body), "\"")
	if token == "" {
		return nil, fmt.Errorf("Token vazio na resposta")
	}

	return &AuthResponse{Token: token}, nil
}

func NewAuthService() IService {
	return &service{}
}
