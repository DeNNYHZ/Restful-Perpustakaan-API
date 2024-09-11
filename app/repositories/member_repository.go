package repositories

import (
	"Restful-Perpustakaan-API/app/models"
	"database/sql"
	"errors"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// MemberRepository provides methods for interacting with member data in the database.
type MemberRepository struct {
	db *sql.DB
}

// NewMemberRepository creates a new MemberRepository instance.
func NewMemberRepository(db *sql.DB) *MemberRepository {
	return &MemberRepository{db: db}
}

// GetAllMembers retrieves all members from the database.
func (mr *MemberRepository) GetAllMembers() ([]models.Member, error) {
	query := `SELECT id, name, gender, email, phone_number, address, registration_date, membership_type FROM members`

	rows, err := mr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []models.Member
	for rows.Next() {
		var m models.Member
		err := rows.Scan(&m.ID, &m.Name, &m.Gender, &m.Email, &m.PhoneNumber, &m.Address, &m.RegistrationDate, &m.MembershipType)
		if err != nil {
			return nil, err
		}
		members = append(members, m)
	}

	return members, nil
}

// GetMemberByID retrieves a member by their ID.
func (mr *MemberRepository) GetMemberByID(id int) (*models.Member, error) {
	query := `SELECT id, name, gender, email, phone_number, address, registration_date, membership_type FROM members WHERE id = $1`

	var m models.Member
	err := mr.db.QueryRow(query, id).Scan(&m.ID, &m.Name, &m.Gender, &m.Email, &m.PhoneNumber, &m.Address, &m.RegistrationDate, &m.MembershipType)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("member not found")
		}
		return nil, err
	}
	return &m, nil
}

// CreateMember creates a new member in the database.
func (mr *MemberRepository) CreateMember(m *models.Member) error {
	query := `
        INSERT INTO members (name, gender, email, phone_number, address, registration_date, membership_type) 
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id
    `

	err := mr.db.QueryRow(query, m.Name, m.Gender, m.Email, m.PhoneNumber, m.Address, time.Now(), m.MembershipType).Scan(&m.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateMember updates an existing member in the database
func (mr *MemberRepository) UpdateMember(m *models.Member) error {
	query := `
		UPDATE members
		SET name = $1, gender = $2, email = $3, phone_number = $4, address = $5, membership_type = $6
		WHERE id = $7
	`

	_, err := mr.db.Exec(query, m.Name, m.Gender, m.Email, m.PhoneNumber, m.Address, m.MembershipType, m.ID)
	return err
}

// DeleteMember deletes a member from the database.
func (mr *MemberRepository) DeleteMember(id int) error {
	query := "DELETE FROM members WHERE id = $1"
	_, err := mr.db.Exec(query, id)
	return err
}

// GetMemberByEmail retrieves a member by their email.
func (mr *MemberRepository) GetMemberByEmail(email string) (*models.Member, error) {
	query := `SELECT id, name, gender, email, phone_number, address, registration_date, membership_type FROM members WHERE email = $1`

	var m models.Member
	err := mr.db.QueryRow(query, email).Scan(&m.ID, &m.Name, &m.Gender, &m.Email, &m.PhoneNumber, &m.Address, &m.RegistrationDate, &m.MembershipType)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Return nil if member not found
		}
		return nil, err
	}
	return &m, nil
}

func (mr *MemberRepository) GetTotalMembers() (interface{}, interface{}) {
	// TODO: Implement this method

	return nil, nil

}

// ... (You can add more repository methods as needed)
