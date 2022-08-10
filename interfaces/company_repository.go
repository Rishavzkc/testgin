package interfaces

import "testcompany/models"

type CompanyRepository interface {
	GetAllCompany() ([]models.Company, error)
	CreateCompany(c models.Company) (models.Company, error)
	GetCompanyById(id int) (models.Company, error)
	DeleteCompany(id int) (models.Company, error)
	UpdateCompany(c models.Company) error
}
