package stops

type StopRequest struct {
	InitialDate string `json:"data_inicial"`
	FinalDate   string `json:"data_final"`
}

type StopResponse []struct {
	Placa       string      `json:"placa"`
	DataInicial DataInicial `json:"data_inicial"`
	DataFinal   DataFinal   `json:"data_final"`
	Tempo       int         `json:"tempo"`
	Latitude    string      `json:"latitude"`
	Longitude   string      `json:"longitude"`
	Address     string      `json:"address"`
}

type DataInicial struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}

type DataFinal struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}
