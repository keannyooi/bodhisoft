// TODO: replace with actual SQL database handling
package medicine

import (
	"errors"
	"strconv"
	"sync"
)

type Store struct {
	mu        sync.Mutex
	medicines map[string]Medicine
}

func NewStore() *Store {
	return &Store{
		medicines: make(map[string]Medicine),
	}
}

func (s *Store) GetAll() []Medicine {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := []Medicine{}
	for _, medicine := range s.medicines {
		result = append(result, medicine)
	}

	return result
}

func (s *Store) Create(request CreateMedicineRequest) Medicine {
	s.mu.Lock()
	defer s.mu.Unlock()

	numericId := uint(len(s.medicines) + 1) // simple auto-incrementing ID
	medicine := Medicine{
		ID:            numericId,
		Code:          "MED" + strconv.FormatUint(uint64(numericId), 10), // param2 is the base for conversion, 10 means decimal
		Name:          request.Name,
		Type:          request.Type,
		StrengthValue: request.StrengthValue,
		StrengthUnit:  request.StrengthUnit,
		Description:   request.Description,
		Status:        "Active",
	}

	s.medicines[medicine.Code] = medicine
	return medicine
}

func (s *Store) Update(code string, request UpdateMedicineRequest) (Medicine, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	medicine, ok := s.medicines[code]
	if !ok {
		return Medicine{}, errors.New("medicine not found")
	}

	// for every field in the request, if it's not nil, update the corresponding field in the medicine
	for request.Name != nil {
		medicine.Name = *request.Name
		break
	}

	medicine.Name = *request.Name
	medicine.Type = *request.Type
	medicine.StrengthValue = *request.StrengthValue
	medicine.StrengthUnit = *request.StrengthUnit
	medicine.Description = *request.Description
	medicine.Status = *request.Status
	s.medicines[code] = medicine

	return medicine, nil
}

func (s *Store) Delete(code string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.medicines[code]; !ok {
		return errors.New("medicine not found")
	}

	delete(s.medicines, code)
	return nil
}
