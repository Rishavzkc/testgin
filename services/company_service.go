package services

import (
	"testcompany/interfaces"
	"testcompany/models"
)

type CompanyService struct {
	repository interfaces.CompanyRepository
}

func NewCompanyService(r interfaces.CompanyRepository) *CompanyService {
	return &CompanyService{
		repository: r,
	}
}

func (s *CompanyService) CreateCompany(data models.Company) (models.Company, error) {
	comp := models.NewCompany(data.ID, data.Name, data.Location)

	return s.repository.CreateCompany(*comp)
}

func (s *CompanyService) UpdateCompany(data models.Company) error {
	comp := models.NewCompanyWithID(data.ID, data.Name, data.Location)

	err := s.repository.UpdateCompany(*comp)

	if err != nil {
		return err
	}

	return nil
}

func (s *CompanyService) DeleteCompany(id int) (models.Company, error) {
	return s.repository.DeleteCompany(id)
}

func (s *CompanyService) ShowCompany(id int) (models.Company, error) {
	return s.repository.GetCompanyById(id)
}

func (s *CompanyService) ShowAllCompanies() ([]models.Company, error) {
	return s.repository.GetAllCompany()
}
