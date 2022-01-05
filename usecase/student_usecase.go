package usecase

import (
	"go-mongodb/model"
	"go-mongodb/repository"
)

type IStudentUseCase interface {
	NewRegistration(student model.Student) (*model.Student, error)
	FindAllStudent() ([]model.Student, error)
	FindStudentInfoByName(name string) (*model.Student, error)
	FindAllStudentWithPagination(skip int64, limit int64) ([]model.Student, error)
}

type StudentUseCase struct {
	repo repository.IStudentRepository
}

func (s *StudentUseCase) NewRegistration(student model.Student) (*model.Student, error) {
	return s.repo.CreateOne(student)
}

func (s *StudentUseCase) FindAllStudent() ([]model.Student, error) {
	return s.repo.GetAll()
}

func (s *StudentUseCase) FindStudentInfoByName(name string) (*model.Student, error) {
	return s.repo.GetOneByUsername(name)
}

func (s *StudentUseCase) FindAllStudentWithPagination(skip int64, limit int64) ([]model.Student, error) {
	return s.repo.GetWithPage(skip, limit)
}

func NewStudentUseCase(studentRepository repository.IStudentRepository) IStudentUseCase {
	return &StudentUseCase{
		studentRepository,
	}
}
