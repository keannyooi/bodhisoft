package medicine_test

import (
	"bodhisoft-backend/internal/medicine"
	"errors"
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
			name: "neg - invalid strength unit",
			req: medicine.CreateMedicineRequest{
				Name:          "Paracetamol",
				Type:          medicine.TypeTablet,
				StrengthValue: 500,
				StrengthUnit:  "InvalidUnit",
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
				return
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
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		repo *medicine.Repo
		// Named input parameters for target function.
		code    string
		req     medicine.UpdateMedicineRequest
		want    medicine.Medicine
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := medicine.NewService(tt.repo)
			got, gotErr := s.UpdateMedicine(tt.code, tt.req)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UpdateMedicine() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UpdateMedicine() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("UpdateMedicine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMedicineService_DeleteMedicine(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		repo *medicine.Repo
		// Named input parameters for target function.
		code    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := medicine.NewService(tt.repo)
			gotErr := s.DeleteMedicine(tt.code)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("DeleteMedicine() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("DeleteMedicine() succeeded unexpectedly")
			}
		})
	}
}
