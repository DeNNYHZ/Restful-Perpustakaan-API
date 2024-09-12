package repositories

import (
	"Restful-Perpustakaan-API/app/models"
	"database/sql"
	"errors"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// MemberRepository provides methods for interacting with member data in the database.
type MemberRepository interface {
	GetAllMembers() ([]models.Member, error)
	GetMemberByID(id int) (*models.Member, error)
	GetMemberByEmail(email string) (*models.Member, error)
	CreateMember(m *models.Member) error
	UpdateMember(m *models.Member) error
	DeleteMember(id int) error
}

type memberRepository struct {
	db *sql.DB
}

func (mr *memberRepository) GetAllMembers() ([]models.Member, error) {
	var members []models.Member
	rows, err := mr.db.Query("SELECT * FROM members")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var member models.Member
		err = rows.Scan(&member.ID, &member.Name, &member.Email)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
}

func (mr *memberRepository) GetMemberByID(id int) (*models.Member, error) {
	var member models.Member
	err := mr.db.QueryRow("SELECT * FROM members WHERE id = $1", id).Scan(&member.ID, &member.Name, &member.Email)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (mr *memberRepository) GetMemberByEmail(email string) (*models.Member, error) {
	var member models.Member
	err := mr.db.QueryRow("SELECT * FROM members WHERE email = $1", email).Scan(&member.ID, &member.Name, &member.Email)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (mr *memberRepository) CreateMember(m *models.Member) error {
	_, err := mr.db.Exec("INSERT INTO members (name, email) VALUES ($1, $2)", m.Name, m.Email)
	return err
}

func (mr *memberRepository) UpdateMember(m *models.Member) error {
	_, err := mr.db.Exec("UPDATE members SET name = $1, email = $2 WHERE id = $3", m.Name, m.Email, m.ID)
	return err
}

func (mr *memberRepository) DeleteMember(id int) error {
	_, err := mr.db.Exec("DELETE FROM members WHERE id = $1", id)
	return err
}

// NewMemberRepository creates a new MemberRepository instance.
func NewMemberRepository(db *sql.DB) MemberRepository {
	return &memberRepository{db: db}
}

// getAllMembers retrieves all members from the database.
func (mr *memberRepository) getAllMembers() ([]models.Member, error) {
	query := `SELECT id, name, gender, email, phone_number, address, registration_date, membership_type FROM members`

	rows, err := mr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []models.Member
	for rows.Next() {
		var member models.Member
		err := rows.Scan(&member.ID, &member.Name, &member.Gender, &member.Email, &member.PhoneNumber, &member.Address, &member.RegistrationDate, &member.MembershipType)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	return members, nil
}

// getMemberByID retrieves a member by their ID.
func (mr *memberRepository) getMemberByID(id int) (*models.Member, error) {
	query := `SELECT id, name, gender, email, phone_number, address, registration_date, membership_type FROM members WHERE id = $1`

	var member models.Member
	err := mr.db.QueryRow(query, id).Scan(&member.ID, &member.Name, &member.Gender, &member.Email, &member.PhoneNumber, &member.Address, &member.RegistrationDate, &member.MembershipType)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("member not found")
		}
		return nil, err
	}
	return &member, nil
}

// getMemberByEmail retrieves a member by their email.
func (mr *memberRepository) getMemberByEmail(email string) (*models.Member, error) {
	query := `SELECT id, name, gender, email, phone_number, address, registration_date, membership_type FROM members WHERE email = $1`

	var member models.Member
	err := mr.db.QueryRow(query, email).Scan(&member.ID, &member.Name, &member.Gender, &member.Email, &member.PhoneNumber, &member.Address, &member.RegistrationDate, &member.MembershipType)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Return nil if member not found
		}
		return nil, err
	}
	return &member, nil
}

// createMember creates a new member in the database.
func (mr *memberRepository) createMember(m *models.Member) error {
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

// updateMember updates an existing member in the database
func (mr *memberRepository) updateMember(m *models.Member) error {
	query := `
		UPDATE members
		SET name = $1, gender = $2, email = $3, phone_number = $4, address = $5, membership_type = $6
		WHERE id = $7
	`

	_, err := mr.db.Exec(query, m.Name, m.Gender, m.Email, m.PhoneNumber, m.Address, m.MembershipType, m.ID)
	return err
}

// deleteMember deletes a member from the database.
func (mr *memberRepository) deleteMember(id int) error {
	query := "DELETE FROM members WHERE id = $1"
	_, err := mr.db.Exec(query, id)
	return err
}

// getTotalMembers is not implemented yet
func (mr *memberRepository) getTotalMembers() (interface{}, interface{}) {
	return nil, nil
}
