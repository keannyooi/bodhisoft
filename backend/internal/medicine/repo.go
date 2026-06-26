// TODO: replace with actual SQL database handling
package medicine

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrMedicineNotFound = errors.New("Medicine not found")

type Repo struct {
	db *sql.DB
}

func NewRepo(dbPointer *sql.DB) *Repo {
	return &Repo{
		db: dbPointer,
	}
}

func (s *Repo) GetByCode(code string) (Medicine, error) {
	var medicine Medicine

	row := s.db.QueryRow("SELECT * FROM medicine WHERE medicineCode = ?;", code)
	if err := row.Scan( // row is automatically closed after this function call
		&medicine.ID,
		&medicine.Code,
		&medicine.Name,
		&medicine.Type,
		&medicine.StrengthValue,
		&medicine.StrengthUnit,
		&medicine.Description,
		&medicine.Status,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return medicine, ErrMedicineNotFound
		}
		return medicine, err
	}
	return medicine, nil
}

func (s *Repo) GetAll() ([]Medicine, error) {
	var medicines []Medicine

	rows, err := s.db.Query("SELECT * FROM medicine;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var medicine Medicine
		if err := rows.Scan(
			&medicine.ID,
			&medicine.Code,
			&medicine.Name,
			&medicine.Type,
			&medicine.StrengthValue,
			&medicine.StrengthUnit,
			&medicine.Description,
			&medicine.Status,
		); err != nil {
			return nil, err
		}
		medicines = append(medicines, medicine)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return medicines, nil
}

func (s *Repo) Create(req CreateMedicineRequest) (string, error) {
	// numericId := uint(len(s.medicines) + 1) // simple auto-incrementing ID
	// medicine := Medicine{
	// 	ID:            numericId,
	// 	Code:          "MED" + fmt.Sprintf("%05d", numericId), // e.g. MED00001, MED00002, etc.
	// 	Name:          req.Name,
	// 	Type:          req.Type,
	// 	StrengthValue: req.StrengthValue,
	// 	StrengthUnit:  req.StrengthUnit,
	// 	Description:   req.Description,
	// 	Status:        "Available",
	// }

	// TODO: turn into transaction
	result, err := s.db.Exec(`
		INSERT INTO medicine (name, type, strengthValue, strengthUnit, description)
		VALUES (?, ?, ?, ?, ?);
	`, req.Name, req.Type, req.StrengthValue, req.StrengthUnit, req.Description)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	code := "MED" + fmt.Sprintf("%05d", id)
	_, err = s.db.Exec("UPDATE medicine SET medicineCode = ? WHERE medicineId = ?", code, id)
	if err != nil {
		return "", err
	}

	return code, nil
}

func (s *Repo) Update(medicine Medicine) error {
	_, err := s.db.Exec(`
		UPDATE medicine
		SET name = ?, type = ?, strengthValue = ?, strengthUnit = ?, description = ?, status = ?
		WHERE medicineId = ?;
	`, medicine.Name, medicine.Type, medicine.StrengthValue, medicine.StrengthUnit, medicine.Description, medicine.Status, medicine.ID)
	if err != nil {
		print(err)
		return err
	}

	return nil
}

func (s *Repo) Delete(code string) error {
	// check if medicine code exists
	// if err != nil {
	// 	return ErrMedicineNotFound
	// }

	// TODO: prevent user from deleting medicine if it's still referenced in prescriptions
	_, err := s.db.Exec("DELETE FROM medicine WHERE medicineCode = ?", code)
	if err != nil {
		return err
	}

	return nil
}
