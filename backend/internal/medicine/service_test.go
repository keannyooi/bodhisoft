package medicine_test

import (
	"bodhisoft-backend/internal/medicine"
	"database/sql"
	"errors"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestMedicineService_CreateMedicine(t *testing.T) {
	tests := []struct {
		name    string
		req     medicine.CreateMedicineRequest
		want    medicine.Medicine
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
			dbMock, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("failed to open sqlmock database: %v", err)
			}
			defer dbMock.Close()

			repo := medicine.NewRepo(dbMock)
			s := medicine.NewService(repo)

			if strings.Contains(tt.name, "pos - ") {
				mock.ExpectExec("^INSERT INTO medicine").WithArgs(tt.req.Name, tt.req.Type, tt.req.StrengthValue, tt.req.StrengthUnit, tt.req.Description).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("^UPDATE medicine SET medicineCode =").WithArgs(sqlmock.AnyArg(), int64(1)).WillReturnResult(sqlmock.NewResult(0, 1))
			}

			gotCode, gotErr := s.CreateMedicine(tt.req)

			if !errors.Is(tt.wantErr, gotErr) {
				t.Fatalf("CreateMedicine() failed: got %v, want %v", gotErr, tt.wantErr)
			}

			if strings.Contains(tt.name, "pos - ") {
				if err := mock.ExpectationsWereMet(); err != nil {
					t.Fatalf("unfulfilled expectations: %v", err)
				}
				if gotCode == "" {
					t.Fatal("CreateMedicine() returned empty medicine code")
				}
			}
		})
	}
}

func TestMedicineService_UpdateMedicine(t *testing.T) {
	updateName := "Strepsils"
	updateType := medicine.TypeSyrup
	updateStrengthUnit := medicine.UnitMgml
	updateStatus := medicine.StatusDiscontinued
	updateDesc := "uvuvwevwevwe onyetenyevwe ugwemubwem ossas"

	var updateInvalidType medicine.MedicineType = "InvalidType"
	var updateInvalidStrengthUnit medicine.StrengthUnit = "InvalidUnit"
	var updateInvalidStatus medicine.MedicineStatus = "InvalidStatus"

	tests := []struct {
		name    string
		code    string
		req     medicine.UpdateMedicineRequest
		want    medicine.Medicine
		wantErr error
	}{
		{
			name: "pos - multiple changes",
			code: "MED00001",
			req: medicine.UpdateMedicineRequest{
				Name:         &updateName,
				Description:  &updateDesc,
				Type:         &updateType,
				StrengthUnit: &updateStrengthUnit,
				Status:       &updateStatus,
			},
			want: medicine.Medicine{
				Name:         updateName,
				Description:  updateDesc,
				Type:         updateType,
				StrengthUnit: updateStrengthUnit,
				Status:       updateStatus,
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
			name: "neg - invalid strength unit (1)",
			code: "MED00001",
			req: medicine.UpdateMedicineRequest{
				StrengthUnit: &updateInvalidStrengthUnit,
			},
			want:    medicine.Medicine{},
			wantErr: medicine.ErrInvalidStrengthUnit,
		},
		{
			name: "neg - invalid strength unit (2)",
			code: "MED00001",
			req: medicine.UpdateMedicineRequest{
				Type: &updateType,
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbMock, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("failed to open sqlmock database: %v", err)
			}
			defer dbMock.Close()

			repo := medicine.NewRepo(dbMock)
			s := medicine.NewService(repo)

			if strings.Contains(tt.name, "pos - ") || tt.wantErr == medicine.ErrInvalidMedicineType || tt.wantErr == medicine.ErrInvalidStrengthUnit || tt.wantErr == medicine.ErrInvalidMedicineStatus {
				rows := sqlmock.NewRows([]string{"medicineId", "medicineCode", "name", "type", "strengthValue", "strengthUnit", "description", "status"}).AddRow(1, tt.code, "Paracetamol", medicine.TypeTablet, 500, medicine.UnitMg, "Used to treat pain and fever", "Available")
				mock.ExpectQuery("^SELECT \\* FROM medicine WHERE medicineCode = \\?;$").WithArgs(tt.code).WillReturnRows(rows)
			}

			if tt.wantErr == medicine.ErrMedicineNotFound {
				mock.ExpectQuery("^SELECT \\* FROM medicine WHERE medicineCode = \\?;$").WithArgs(tt.code).WillReturnError(sql.ErrNoRows)
			}

			strengthValue := uint(500)
			if tt.req.StrengthValue != nil {
				strengthValue = *tt.req.StrengthValue
			}

			if strings.Contains(tt.name, "pos - ") {
				mock.ExpectExec("^UPDATE medicine").WithArgs(updateName, updateType, strengthValue, updateStrengthUnit, updateDesc, updateStatus, int64(1)).WillReturnResult(sqlmock.NewResult(0, 1))
			}

			got, gotErr := s.UpdateMedicine(tt.code, tt.req)
			if !errors.Is(tt.wantErr, gotErr) {
				t.Fatalf("UpdateMedicine() failed: got %v, want %v", gotErr, tt.wantErr)
			}

			if tt.wantErr == nil && strings.Contains(tt.name, "pos - ") {
				if err := mock.ExpectationsWereMet(); err != nil {
					t.Fatalf("unfulfilled expectations: %v", err)
				}
				if got.Name != tt.want.Name || got.Type != tt.want.Type || got.StrengthUnit != tt.want.StrengthUnit || got.Description != tt.want.Description {
					t.Fatalf("UpdateMedicine() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestMedicineService_DeleteMedicine(t *testing.T) {
	tests := []struct {
		name    string
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
		t.Run(tt.name, func(t *testing.T) {
			dbMock, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("failed to open sqlmock database: %v", err)
			}
			defer dbMock.Close()

			repo := medicine.NewRepo(dbMock)
			s := medicine.NewService(repo)

			if strings.Contains(tt.name, "pos - ") {
				mock.ExpectExec("^DELETE FROM medicine").WithArgs(tt.code).WillReturnResult(sqlmock.NewResult(0, 1))
			}

			gotErr := s.DeleteMedicine(tt.code)
			if !errors.Is(tt.wantErr, gotErr) {
				t.Fatalf("DeleteMedicine() failed: got %v, want %v", gotErr, tt.wantErr)
			}

			if strings.Contains(tt.name, "pos - ") {
				if err := mock.ExpectationsWereMet(); err != nil {
					t.Fatalf("unfulfilled expectations: %v", err)
				}
			}
		})
	}
}
