package entity

type (
	Client struct {
		UUID      string `json:"uuid,omitempty"`
		Name      string `json:"name,omitempty"`
		CPF       string `json:"cpf,omitempty"`
		Email     string `json:"email,omitempty"`
		CreatedAt string ` json:"createdAt,omitempty"`
	}

	Payment struct {
		Price  float64 `json:"price,omitempty"`
		Client Client  `json:"payment,omitempty"`
	}
)
