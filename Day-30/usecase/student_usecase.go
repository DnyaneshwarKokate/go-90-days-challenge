package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"day-30/domain"

	"github.com/google/uuid"
)

type StudentUseCase struct {
	studentRepo domain.StudentRepository
	logger      domain.Logger
}

func NewStudentUseCase(studentRepo domain.StudentRepository, logger domain.Logger) *StudentUseCase {
	return &StudentUseCase{
		studentRepo: studentRepo,
		logger:      logger,
	}
}

func (u *StudentUseCase) CreateStudent(ctx context.Context, input domain.CreateStudentInput) (*domain.Student, error) {
	u.logger.Info(ctx, "Creating new student profile", "name", input.FullName, "email", input.Email)

	if strings.TrimSpace(input.FullName) == "" {
		u.logger.Warn(ctx, "Student validation failed: Empty full name")
		return nil, domain.ErrInvalidInput
	}

	email := strings.ToLower(strings.TrimSpace(input.Email))
	if !strings.Contains(email, "@") {
		u.logger.Warn(ctx, "Student validation failed: Invalid email format", "email", email)
		return nil, domain.ErrInvalidInput
	}

	existingStudent, err := u.studentRepo.FindByEmail(ctx, email)
	if err == nil && existingStudent != nil {
		u.logger.Warn(ctx, "Student creation rejected: Duplicate email", "email", email)
		return nil, domain.ErrEmailExists
	}

	userID, _ := ctx.Value(domain.UserIDKey).(string)

	now := time.Now()
	student := &domain.Student{
		ID:         fmt.Sprintf("std_%s", uuid.New().String()[:8]),
		UserID:     userID,
		FullName:   strings.TrimSpace(input.FullName),
		Email:      email,
		Department: strings.ToUpper(strings.TrimSpace(input.Department)),
		GPA:        input.GPA,
		Status:     "ACTIVE",
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	if err := u.studentRepo.Save(ctx, student); err != nil {
		u.logger.Error(ctx, "Failed to persist student record", "student_id", student.ID, "error", err)
		return nil, err
	}

	u.logger.Info(ctx, "Student record created successfully", "student_id", student.ID, "email", student.Email)
	return student, nil
}

func (u *StudentUseCase) GetStudentByID(ctx context.Context, id string) (*domain.Student, error) {
	u.logger.Debug(ctx, "Fetching student by ID", "student_id", id)

	student, err := u.studentRepo.FindByID(ctx, id)
	if err != nil {
		if err == domain.ErrStudentNotFound {
			u.logger.Warn(ctx, "Student profile requested but not found", "student_id", id)
		} else {
			u.logger.Error(ctx, "Unexpected error retrieving student", "student_id", id, "error", err)
		}
		return nil, err
	}

	return student, nil
}

func (u *StudentUseCase) ListStudents(ctx context.Context, department string, status string) ([]*domain.Student, error) {
	u.logger.Info(ctx, "Listing student profiles", "department", department, "status", status)
	return u.studentRepo.FindAll(ctx, department, status)
}

func (u *StudentUseCase) UpdateStudent(ctx context.Context, id string, input domain.UpdateStudentInput) (*domain.Student, error) {
	u.logger.Info(ctx, "Updating student profile", "student_id", id)

	student, err := u.studentRepo.FindByID(ctx, id)
	if err != nil {
		u.logger.Warn(ctx, "Student update failed: Record not found", "student_id", id)
		return nil, domain.ErrStudentNotFound
	}

	if strings.TrimSpace(input.FullName) != "" {
		student.FullName = strings.TrimSpace(input.FullName)
	}

	if strings.TrimSpace(input.Department) != "" {
		student.Department = strings.ToUpper(strings.TrimSpace(input.Department))
	}

	if input.GPA > 0 {
		student.GPA = input.GPA
	}

	if strings.TrimSpace(input.Status) != "" {
		student.Status = strings.ToUpper(strings.TrimSpace(input.Status))
	}

	student.UpdatedAt = time.Now()

	if err := u.studentRepo.Update(ctx, student); err != nil {
		u.logger.Error(ctx, "Failed to update student record", "student_id", id, "error", err)
		return nil, err
	}

	u.logger.Info(ctx, "Student record updated successfully", "student_id", id)
	return student, nil
}

func (u *StudentUseCase) DeleteStudent(ctx context.Context, id string) error {
	u.logger.Info(ctx, "Deleting student profile", "student_id", id)

	if err := u.studentRepo.Delete(ctx, id); err != nil {
		u.logger.Warn(ctx, "Student deletion failed", "student_id", id, "error", err)
		return err
	}

	u.logger.Info(ctx, "Student record deleted successfully", "student_id", id)
	return nil
}
