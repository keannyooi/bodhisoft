package medicine

import (
	"errors"
	"strings"
)

// sentinel errors cuz Go recommends 'em
var ErrInvalidMedicineCode = errors.New("Invalid medicine code")
var ErrInvalidMedicineType = errors.New("Invalid medicine type")
var ErrInvalidStrengthUnit = errors.New("Invalid strength unit")
var ErrInvalidMedicineStatus = errors.New("Invalid strength unit")
var ErrMissingRequiredFields = errors.New("Missing required fields")

// ================================================================================
type MedicineService struct {
	repo *Repo
}

func NewService(repo *Repo) *MedicineService {
	return &MedicineService{repo: repo}
}

func (s *MedicineService) CreateMedicine(req CreateMedicineRequest) (Medicine, error) {
	// check completeness of request body
	if req.Name == "" || req.Type == "" || req.StrengthValue == 0 || req.StrengthUnit == "" {
		return Medicine{}, ErrMissingRequiredFields
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

func (s *MedicineService) GetMedicine(code string) (Medicine, error) {
	if !strings.HasPrefix(code, "MED") {
		return Medicine{}, ErrInvalidMedicineCode
	}

	medicine, err := s.repo.GetByCode(code)
	if err != nil {
		return Medicine{}, err
	}

	return medicine, nil
}

func (s *MedicineService) GetMedicines() []Medicine {
	return s.repo.GetAll()
}

func (s *MedicineService) UpdateMedicine(code string, req UpdateMedicineRequest) (Medicine, error) {
	refMedicine, err := s.GetMedicine(code)
	if err != nil {
		return Medicine{}, err
	}

	refMedicineType := refMedicine.Type
	if req.Type != nil {
		err := validateMedicineType(*req.Type)
		if err != nil {
			return Medicine{}, err
		}

		refMedicineType = *req.Type
	}

	if req.StrengthUnit != nil {
		err := validateStrengthUnit(refMedicineType, *req.StrengthUnit)
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

	refMedicine.Set(req)
	return refMedicine, s.repo.Update(refMedicine) // TODO: refactor?
}

func (s *MedicineService) DeleteMedicine(code string) error {
	if !strings.HasPrefix(code, "MED") {
		return ErrInvalidMedicineCode
	}
	return s.repo.Delete(code)
}

// ================================================================================
// validation functions since i don't want to write two sets of functions for each type of request
func validateMedicineType(t MedicineType) error {
	if !MedicineType(t).IsValid() {
		return ErrInvalidMedicineType
	}
	return nil
}

func validateStrengthUnit(t MedicineType, u StrengthUnit) error {
	// only "Tablet" and "Pill" type can have "mg" or "g" strength units
	// and "Syrup" type can only have "mg/ml" strength unit
	switch t {
	case TypeTablet, TypePill:
		if u != UnitMg && u != UnitG {
			return ErrInvalidStrengthUnit
		}
	case TypeSyrup:
		if u != UnitMgml {
			return ErrInvalidStrengthUnit
		}
	default:
		// note: duplicated validation, beware
		return ErrInvalidMedicineType
	}
	return nil
}

func validateMedicineStatus(s MedicineStatus) error {
	if !MedicineStatus(s).IsValid() {
		return ErrInvalidMedicineStatus
	}
	return nil
}
