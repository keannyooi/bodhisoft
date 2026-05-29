package medicine_test

import (
	"bodhisoft-backend/internal/medicine"
	"errors"
	"log/slog"
	"testing"
)

func TestMedicineService_CreateMedicine(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		req     medicine.CreateMedicineRequest
		want    medicine.Medicine // TODO: replace with comparison of selected fields only
		wantErr error
	}{
		{
			name: "pos - with description",
			req: medicine.CreateMedicineRequest{
				Name:          "Paracetamol",
				Type:          medicine.TypeTablet,
				StrengthValue: 500,
				StrengthUnit:  medicine.UnitMg,
				Description:   "Used to treat pain and fever",
			},
			want: medicine.Medicine{
				Name:          "Paracetamol",
				Type:          medicine.TypeTablet,
				StrengthValue: 500,
				StrengthUnit:  medicine.UnitMg,
				Description:   "Used to treat pain and fever",
			},
			wantErr: nil,
		},
		{
			name: "neg - missing fields",
			req: medicine.CreateMedicineRequest{
				Name:          "Paracetamol",
				Type:          medicine.TypeTablet,
				StrengthValue: 0,
				StrengthUnit:  medicine.UnitMg,
			},
			want:    medicine.Medicine{},
			wantErr: medicine.ErrMissingRequiredFields,
		},
		{
			name: "neg - invalid strength unit (1)",
			req: medicine.CreateMedicineRequest{
				Name:          "Paracetamol",
				Type:          medicine.TypeSyrup,
				StrengthValue: 500,
				StrengthUnit:  "InvalidUnit",
			},
			want:    medicine.Medicine{},
			wantErr: medicine.ErrInvalidStrengthUnit,
		},
		{
			name: "neg - invalid strength unit (2)",
			req: medicine.CreateMedicineRequest{
				Name:          "Paracetamol",
				Type:          medicine.TypeTablet,
				StrengthValue: 500,
				StrengthUnit:  medicine.UnitMgml,
			},
			want:    medicine.Medicine{},
			wantErr: medicine.ErrInvalidStrengthUnit,
		},
		{
			name: "neg - invalid medicine type",
			req: medicine.CreateMedicineRequest{
				Name:          "Paracetamol",
				Type:          "InvalidType",
				StrengthValue: 500,
				StrengthUnit:  medicine.UnitMg,
			},
			want:    medicine.Medicine{},
			wantErr: medicine.ErrInvalidMedicineType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := medicine.NewRepo() // to prevent the results of unit tests from affecting each other
			s := medicine.NewService(repo)
			got, gotErr := s.CreateMedicine(tt.req)

			if !errors.Is(tt.wantErr, gotErr) {
				t.Errorf("CreateMedicine() failed: %v, expected %v", gotErr.Error(), tt.wantErr.Error())
			} else {
				nameMatch := tt.want.Name == got.Name
				typeMatch := tt.want.Type == got.Type
				strengthValueMatch := tt.want.StrengthValue == got.StrengthValue
				strengthUnitMatch := tt.want.StrengthUnit == got.StrengthUnit
				descMatch := tt.want.Description == got.Description

				if !nameMatch || !typeMatch || !strengthValueMatch || !strengthUnitMatch || !descMatch {
					t.Errorf("CreateMedicine() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestMedicineService_UpdateMedicine(t *testing.T) {
	updateName := "Strepsils"
	updateDesc := ""
	var updateInvalidType medicine.MedicineType = "InvalidType"
	updateInvalidStrengthUnit := medicine.UnitMgml
	var updateInvalidStatus medicine.MedicineStatus = "InvalidStatus"

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		code    string
		req     medicine.UpdateMedicineRequest
		want    medicine.Medicine
		wantErr error
	}{
		{
			name: "pos - multiple changes",
			code: "MED00001",
			req: medicine.UpdateMedicineRequest{
				Name:        &updateName,
				Description: &updateDesc,
			},
			want: medicine.Medicine{
				Name:        updateName,
				Description: updateDesc,
			},
			wantErr: nil,
		},
		{
			name: "neg - nonexistent medicine",
			code: "MED00003",
			req: medicine.UpdateMedicineRequest{
				Type: &updateInvalidType,
			},
			want:    medicine.Medicine{},
			wantErr: medicine.ErrMedicineNotFound,
		},
		{
			name: "neg - invalid medicine code",
			code: "ED00001",
			req: medicine.UpdateMedicineRequest{
				Type: &updateInvalidType,
			},
			want:    medicine.Medicine{},
			wantErr: medicine.ErrInvalidMedicineCode,
		},
		{
			name: "neg - invalid medicine type",
			code: "MED00001",
			req: medicine.UpdateMedicineRequest{
				Type: &updateInvalidType,
			},
			want:    medicine.Medicine{},
			wantErr: medicine.ErrInvalidMedicineType,
		},
		{
			name: "neg - invalid strength unit",
			code: "MED00001",
			req: medicine.UpdateMedicineRequest{
				StrengthUnit: &updateInvalidStrengthUnit,
			},
			want:    medicine.Medicine{},
			wantErr: medicine.ErrInvalidStrengthUnit,
		},
		{
			name: "neg - invalid medicine status",
			code: "MED00001",
			req: medicine.UpdateMedicineRequest{
				Status: &updateInvalidStatus,
			},
			want:    medicine.Medicine{},
			wantErr: medicine.ErrInvalidMedicineStatus,
		},
	}

	// setup
	repo := medicine.NewRepo() // to prevent the results of unit tests from affecting each other
	s := medicine.NewService(repo)
	_, err := s.CreateMedicine(medicine.CreateMedicineRequest{
		Name:          "Paracetamol",
		Type:          medicine.TypeTablet,
		StrengthValue: 500,
		StrengthUnit:  medicine.UnitMg,
		Description:   "Used to treat pain and fever",
	})
	if err != nil {
		slog.Error("CreateMedicine() failed during test environment setup. Run TestMedicineService_CreateMedicine for more information.")
		return
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// print(tt.req.StrengthUnit)
			got, gotErr := s.UpdateMedicine(tt.code, tt.req)
			if !errors.Is(tt.wantErr, gotErr) {
				t.Errorf("UpdateMedicine() failed: %v, expected %v", gotErr.Error(), tt.wantErr.Error())
			} else {
				nameMatch, typeMatch, strengthValueMatch := true, true, true
				strengthUnitMatch, descMatch := true, true
				if tt.req.Name != nil {
					nameMatch = tt.want.Name == got.Name
				}
				if tt.req.Type != nil {
					typeMatch = tt.want.Type == got.Type
				}
				if tt.req.StrengthValue != nil {
					strengthValueMatch = tt.want.StrengthValue == got.StrengthValue
				}
				if tt.req.StrengthUnit != nil {
					strengthUnitMatch = tt.want.StrengthUnit == got.StrengthUnit
				}
				if tt.req.Description != nil {
					descMatch = tt.want.Description == got.Description
				}

				if !nameMatch || !typeMatch || !strengthValueMatch || !strengthUnitMatch || !descMatch {
					t.Errorf("UpdateMedicine() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestMedicineService_DeleteMedicine(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		// Named input parameters for target function.
		code    string
		wantErr error
	}{
		{
			name:    "pos - delete existing medicine by code",
			code:    "MED00001",
			wantErr: nil,
		},
		{
			name:    "neg - invalid medicine code",
			code:    "MD00001",
			wantErr: medicine.ErrInvalidMedicineCode,
		},
	}

	for _, tt := range tests {
		// setup
		repo := medicine.NewRepo() // to prevent the results of unit tests from affecting each other
		s := medicine.NewService(repo)
		_, err := s.CreateMedicine(medicine.CreateMedicineRequest{
			Name:          "Paracetamol",
			Type:          medicine.TypeTablet,
			StrengthValue: 500,
			StrengthUnit:  medicine.UnitMg,
			Description:   "Used to treat pain and fever",
		})
		if err != nil {
			slog.Error("CreateMedicine() failed during test environment setup. Run TestMedicineService_CreateMedicine for more information.")
			return
		}

		t.Run(tt.name, func(t *testing.T) {
			gotErr := s.DeleteMedicine(tt.code)
			if !errors.Is(tt.wantErr, gotErr) {
				t.Errorf("DeleteMedicine() failed: %v, expected %v", gotErr.Error(), tt.wantErr.Error())
			}

			// test whether or not the medicine is actually deleted

		})
	}
}
