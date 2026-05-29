// TODO: replace with actual SQL database handling
package medicine

import (
	"errors"
	"fmt"
	"sync"
)

var ErrMedicineNotFound = errors.New("Medicine not found")

type Repo struct {
	mu        sync.Mutex
	medicines map[string]Medicine
}

func NewRepo() *Repo {
	return &Repo{
		medicines: make(map[string]Medicine),
	}
}

func (s *Repo) GetByCode(code string) (Medicine, error) {
	medicine, ok := s.medicines[code]
	if !ok {
		return Medicine{}, ErrMedicineNotFound
	}

	return medicine, nil
}

func (s *Repo) GetAll() []Medicine {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := []Medicine{}
	for _, medicine := range s.medicines {
		result = append(result, medicine)
	}

	return result
}

func (s *Repo) Create(request CreateMedicineRequest) Medicine {
	s.mu.Lock()
	defer s.mu.Unlock()

	numericId := uint(len(s.medicines) + 1) // simple auto-incrementing ID
	medicine := Medicine{
		ID:            numericId,
		Code:          "MED" + fmt.Sprintf("%05d", numericId), // e.g. MED00001, MED00002, etc.
		Name:          request.Name,
		Type:          request.Type,
		StrengthValue: request.StrengthValue,
		StrengthUnit:  request.StrengthUnit,
		Description:   request.Description,
		Status:        "Available",
	}

	s.medicines[medicine.Code] = medicine
	return medicine
}

func (s *Repo) Update(medicine Medicine) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.medicines[medicine.Code] = medicine

	return nil
}

func (s *Repo) Delete(code string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.medicines[code]; !ok {
		return ErrMedicineNotFound
	}
	// TODO: prevent user from deleting medicine if it's still referenced in prescriptions

	delete(s.medicines, code)
	return nil
}
