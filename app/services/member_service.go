package services

import (
	"Restful-Perpustakaan-API/app/models"
	"Restful-Perpustakaan-API/app/repositories"
)

// MemberService provides methods for managing members
type MemberService struct {
	memberRepository repositories.MemberRepository
}

// NewMemberService creates a new MemberService instance
func NewMemberService(memberRepository repositories.MemberRepository) *MemberService {
	return &MemberService{memberRepository: memberRepository}
}

// GetAllMembers mengambil semua anggota
func (ms *MemberService) GetAllMembers() ([]models.Member, error) {
	return ms.memberRepository.GetAllMembers()
}

// GetMemberByID mengambil anggota berdasarkan ID
func (ms *MemberService) GetMemberByID(id int) (*models.Member, error) {
	return ms.memberRepository.GetMemberByID(id)
}

// CreateMember membuat anggota baru
func (ms *MemberService) CreateMember(m *models.Member) error {
	// Anda dapat menambahkan logika validasi atau bisnis lainnya di sini sebelum menyimpan anggota ke database
	return ms.memberRepository.CreateMember(m)
}

// UpdateMember memperbarui anggota
func (ms *MemberService) UpdateMember(m *models.Member) error {
	// Anda dapat menambahkan logika validasi atau bisnis lainnya di sini sebelum memperbarui anggota di database
	return ms.memberRepository.UpdateMember(m)
}

// DeleteMember menghapus anggota
func (ms *MemberService) DeleteMember(id int) error {
	return ms.memberRepository.DeleteMember(id)
}

// GetMemberByEmail mengambil anggota berdasarkan email
func (ms *MemberService) GetMemberByEmail(email string) (*models.Member, error) {
	return ms.memberRepository.GetMemberByEmail(email)
}

// ... (fungsi lain yang mungkin Anda butuhkan, seperti SearchMembers, dll.)
