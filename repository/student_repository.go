package repository

import (
	"go-mongodb/db"
	"go-mongodb/model"
	"go-mongodb/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IStudentRepository interface {
	GetAll() ([]model.Student, error)
	GetOneByUsername(name string) (*model.Student, error)
	CreateOne(student model.Student) (*model.Student, error)
	GetWithPage(skip int64, limit int64) ([]model.Student, error)
}

type StudentRepository struct {
	repo *mongo.Collection
}

func (s *StudentRepository) GetAll() ([]model.Student, error) {
	var students []model.Student
	ctx, cancel := utils.InitContext()
	defer cancel()

	cursor, err := s.repo.Find(ctx, bson.D{})
	defer cursor.Close(ctx)

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var student model.Student
		err = cursor.Decode(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (s *StudentRepository) GetOneByUsername(name string) (*model.Student, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	var student model.Student
	err := s.repo.FindOne(ctx, bson.D{{"name", name}}).Decode(&student)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &student, nil
}

func (s *StudentRepository) CreateOne(student model.Student) (*model.Student, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()

	student.Id = primitive.NewObjectID()
	_, err := s.repo.InsertOne(ctx, student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *StudentRepository) GetWithPage(skip int64, limit int64) ([]model.Student, error) {
	var students []model.Student
	ctx, cancel := utils.InitContext()
	defer cancel()

	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}

	cursor, err := s.repo.Find(ctx, bson.D{}, &opts)

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var student model.Student
		err = cursor.Decode(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func NewStudentRepository(resource *db.Resource) IStudentRepository {
	studentCollection := resource.Db.Collection("students")
	studentRepository := &StudentRepository{repo: studentCollection}
	return studentRepository
}
