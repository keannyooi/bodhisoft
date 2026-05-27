package medicine

// struct tags are for json encoding, in this case it essentially converts CamelCase into camelCase
// omitempty means that if the field is empty, it will be omitted from the JSON output
type Medicine struct {
	ID            uint   `json:"id"`
	Code          string `json:"code"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	StrengthValue uint   `json:"strengthValue"`
	StrengthUnit  string `json:"strengthUnit"`
	Description   string `json:"description,omitempty"`
	Status        string `json:"status"`
}

type CreateMedicineRequest struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	StrengthValue uint   `json:"strengthValue"`
	StrengthUnit  string `json:"strengthUnit"`
	Description   string `json:"description,omitempty"`
}

// all fields in UpdateMedicineRequest are pointers so that we can differentiate
// between "field not provided" and "field provided with zero value"
type UpdateMedicineRequest struct {
	Name          *string `json:"name,omitempty"`
	Type          *string `json:"type,omitempty"`
	StrengthValue *uint   `json:"strengthValue,omitempty"`
	StrengthUnit  *string `json:"strengthUnit,omitempty"`
	Description   *string `json:"description,omitempty"`
	Status        *string `json:"status,omitempty"`
}
