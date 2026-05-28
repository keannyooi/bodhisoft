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
	// check completeness of request body
	if req.Name == "" || req.Type == "" || req.StrengthValue == 0 || req.StrengthUnit == "" {
		return Medicine{}, errors.New("Missing required fields")
	}

	err := validateMedicineType(req.Type)
	if err != nil {
		return Medicine{}, err
	}

	err = validateStrengthUnit(req.Type, req.StrengthUnit)
	if err != nil {
		return Medicine{}, err
	}

	return s.repo.Create(req), nil
}

func (s *MedicineService) UpdateMedicine(code string, req UpdateMedicineRequest) (Medicine, error) {
	if !strings.HasPrefix(code, "MED") {
		return Medicine{}, errors.New("Invalid medicine code")
	}

	if req.Type != nil {
		err := validateMedicineType(*req.Type)
		if err != nil {
			return Medicine{}, err
		}
	}

	if req.StrengthUnit != nil {
		err := validateStrengthUnit(*req.Type, *req.StrengthUnit)
		if err != nil {
			return Medicine{}, err
		}
	}

	if req.Status != nil {
		err := validateMedicineStatus(*req.Status)
		if err != nil {
			return Medicine{}, err
		}
	}

	return s.repo.Update(code, req)
}

func (s *MedicineService) DeleteMedicine(code string) error {
	if !strings.HasPrefix(code, "MED") {
		return errors.New("Invalid medicine code")
	}
	return s.repo.Delete(code)
}

// validation functions since i don't want to write two sets of functions for each type of request
func validateMedicineType(t MedicineType) error {
	if !MedicineType(t).IsValid() {
		return errors.New("Invalid medicine type")
	}
	return nil
}

func validateStrengthUnit(t MedicineType, u StrengthUnit) error {
	// only "Tablet" and "Pill" type can have "mg" or "g" strength units
	// and "Syrup" type can only have "mg/ml" strength unit
	switch t {
	case TypeTablet, TypePill:
		if u != UnitMg && u != UnitG {
			return errors.New("Invalid strength unit for Tablet/Pill type")
		}
	case TypeSyrup:
		if u != UnitMgml {
			return errors.New("Invalid strength unit for Syrup type")
		}
	default:
		return errors.New("Invalid medicine type")
	}
	return nil
}

func validateMedicineStatus(s MedicineStatus) error {
	if !MedicineStatus(s).IsValid() {
		return errors.New("Invalid medicine status")
	}
	return nil
}
