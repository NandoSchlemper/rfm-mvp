package stops

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"rfmtransportes-api/services/auth"
	"time"
)

type IService interface {
	GetStops(initial_date, final_date string) (StopResponse, error)
}

type service struct {
	authService auth.IService
	baseUrl     string
	httpClient  *http.Client
}

// GetVehiclesByTime implements [IService].
func (s *service) GetStops(initial_date string, final_date string) (StopResponse, error) {
	response, err := s.authService.Login()
	if err != nil {
		return nil, fmt.Errorf("Erro ao realizar o login: %s", err.Error())
	}

	parameters := url.Values{}
	parameters.Add("data_inicial", initial_date)
	parameters.Add("data_final", final_date)
	parameters.Add("token", response.Token)

	fullUrl := fmt.Sprintf("%s?%s", s.baseUrl+"/paradas", parameters.Encode())

	request, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("Erro ao montar a requisição: %s", err.Error())
	}

	request.Header.Add("Authorization", "Bearer "+response.Token)

	resp, err := s.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Erro ao realizar a requisição: %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Erro ao ler os dados do corpo da resposta: %s", err.Error())
	}

	var jsonResponse StopResponse
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return nil, fmt.Errorf("Erro ao converter resposta para JSON: %s", err.Error())
	}

	filtered := make(StopResponse, 0, len(jsonResponse))
	for _, stop := range jsonResponse {
		if stop.Tempo != 0 {
			filtered = append(filtered, stop)
		}
	}

	return filtered, nil
}

func NewStopService() IService {
	authInstance := auth.NewAuthService()
	baseUrl := os.Getenv("URL_DARWIN_API")
	client := &http.Client{Timeout: 10 * time.Second}

	return &service{
		authService: authInstance,
		httpClient:  client,
		baseUrl:     baseUrl,
	}
}
