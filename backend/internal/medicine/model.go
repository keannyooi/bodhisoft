package medicine

// custom string types to simulate enums since go doesn't support them natively
type MedicineType string

const (
	TypeTablet MedicineType = "Tablet"
	TypePill   MedicineType = "Pill"
	TypeSyrup  MedicineType = "Syrup"
)

func (t MedicineType) IsValid() bool {
	switch t {
	case TypeTablet, TypePill, TypeSyrup:
		return true
	default:
		return false
	}
}

type StrengthUnit string

const (
	UnitMg   StrengthUnit = "mg"
	UnitG    StrengthUnit = "g"
	UnitMgml StrengthUnit = "mg/ml"
)

func (u StrengthUnit) IsValid() bool {
	switch u {
	case UnitMg, UnitG, UnitMgml:
		return true
	default:
		return false
	}
}

type MedicineStatus string

const (
	StatusAvailable    MedicineStatus = "Available"
	StatusDiscontinued MedicineStatus = "Discontinued"
)

func (s MedicineStatus) IsValid() bool {
	switch s {
	case StatusAvailable, StatusDiscontinued:
		return true
	default:
		return false
	}
}

// ================================================================================
// struct tags are for json encoding, in this case it essentially converts CamelCase into camelCase
// omitempty means that if the field is empty, it will be omitted from the JSON output
type Medicine struct {
	ID            uint           `json:"id"`
	Code          string         `json:"code"`
	Name          string         `json:"name"`
	Type          MedicineType   `json:"type"`
	StrengthValue uint           `json:"strengthValue"`
	StrengthUnit  StrengthUnit   `json:"strengthUnit"`
	Description   string         `json:"description,omitempty"`
	Status        MedicineStatus `json:"status"`
}

// all hail the universal setter
func (m *Medicine) Set(request UpdateMedicineRequest) {
	// apparently this is idiomatic go
	if request.Name != nil {
		m.Name = *request.Name
	}
	if request.Type != nil {
		m.Type = *request.Type
	}
	if request.StrengthValue != nil {
		m.StrengthValue = *request.StrengthValue
	}
	if request.StrengthUnit != nil {
		m.StrengthUnit = *request.StrengthUnit
	}
	if request.Description != nil {
		m.Description = *request.Description
	}
	if request.Status != nil {
		m.Status = *request.Status
	}
}

func (m Medicine) String() string {
	return m.Name
}

func (m Medicine) IsAvailable() bool {
	return m.Status == StatusAvailable
}

// ================================================================================
// requests structs
type CreateMedicineRequest struct {
	Name          string       `json:"name"`
	Type          MedicineType `json:"type"`
	StrengthValue uint         `json:"strengthValue"`
	StrengthUnit  StrengthUnit `json:"strengthUnit"`
	Description   string       `json:"description,omitempty"`
}

// all fields in UpdateMedicineRequest are pointers so that we can differentiate
// between "field not provided" and "field provided with zero value"
type UpdateMedicineRequest struct {
	Name          *string         `json:"name,omitempty"`
	Type          *MedicineType   `json:"type,omitempty"`
	StrengthValue *uint           `json:"strengthValue,omitempty"`
	StrengthUnit  *StrengthUnit   `json:"strengthUnit,omitempty"`
	Description   *string         `json:"description,omitempty"`
	Status        *MedicineStatus `json:"status,omitempty"`
}
