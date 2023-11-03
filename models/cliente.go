package models

type Cliente struct {
	ID      uint   `json:"id" gorm:"primary_key;column:IdCliente"`
	Name    string `json:"name" gorm:"column:NomeLongo"`
	Address string `json:"address" gorm:"column:EndCidade" `
}

type Tabler interface {
	TableName() string
  }
  
  // TableName overrides the table name used by Cliente to `Cad_ClientesFornecedores`
  func (Cliente) TableName() string {
	return "Cad_ClientesFornecedores"
  }