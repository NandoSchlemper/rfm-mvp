package auth

type AuthResponse struct {
	Token string `json:"token"`
}

type AuthRequest struct {
	CodigoCliente int    `json:"cod_cliente"`
	Login         string `json:"login"`
	Password      string `json:"senha"`
}
