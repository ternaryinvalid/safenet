package entity

type Account struct {
	Name       string `json:"name"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	Address    string `json:"address"`
}

type AccountCreateDTO struct {
	Name       string  `json:"name"`
	PrivateKey *string `json:"private_key"`
}

type AccountResponseDTO struct {
	Name      string `json:"name"`
	PublicKey string `json:"public_key"`
	Address   string `json:"address"`
}
