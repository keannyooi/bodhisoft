package medicine

import (
	"errors"
	"strings"
)

type MedicineService struct {
	repo *Repo
}

func NewService(repo *Repo) *MedicineService {
	return &MedicineService{repo: repo}
}

func (s *MedicineService) CreateMedicine(req CreateMedicineRequest) (Medicine, error) {
	if req.StrengthValue <= 0 {
		return Medicine{}, errors.New("Strength must be positive")
	}
	if req.StrengthUnit == "" {
		return Medicine{}, errors.New("Invalid strength unit")
	}

	return s.repo.Create(req), nil
}

func (s *MedicineService) UpdateMedicine(code string, req UpdateMedicineRequest) (Medicine, error) {
	if req.Status != nil && *req.Status != "Active" && *req.Status != "Inactive" {
		return Medicine{}, errors.New("Invalid status")
	}

	return s.repo.Update(code, req)
}

func (s *MedicineService) DeleteMedicine(code string) error {
	if !strings.HasPrefix(code, "MED") {
		return errors.New("Invalid medicine code")
	}
	return s.repo.Delete(code)
}
