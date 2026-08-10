package repository

import (
	"context"
	"strings"
	"sync"
	"time"

	"day-30/domain"
)

type memoryUserRepository struct {
	mu     sync.RWMutex
	users  map[string]*domain.User
	emails map[string]string
	logger domain.Logger
}

func NewMemoryUserRepository(logger domain.Logger) domain.UserRepository {
	return &memoryUserRepository{
		users:  make(map[string]*domain.User),
		emails: make(map[string]string),
		logger: logger,
	}
}

func (r *memoryUserRepository) SaveUser(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.logger.Debug(ctx, "DB SaveUser requested", "email", user.Email)

	if existingID, exists := r.emails[user.Email]; exists && existingID != user.ID {
		r.logger.Warn(ctx, "DB Conflict: Email already registered", "email", user.Email)
		return domain.ErrEmailExists
	}

	r.users[user.ID] = user
	r.emails[user.Email] = user.ID
	r.logger.Info(ctx, "DB User saved successfully", "user_id", user.ID, "role", user.Role)
	return nil
}

func (r *memoryUserRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.logger.Debug(ctx, "DB FindUserByEmail requested", "email", email)

	id, exists := r.emails[email]
	if !exists {
		return nil, domain.ErrUserNotFound
	}
	return r.users[id], nil
}

func (r *memoryUserRepository) FindUserByID(ctx context.Context, id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, domain.ErrUserNotFound
	}
	return user, nil
}

type memoryStudentRepository struct {
	mu       sync.RWMutex
	students map[string]*domain.Student
	emails   map[string]string
	logger   domain.Logger
}

func NewMemoryStudentRepository(logger domain.Logger) domain.StudentRepository {
	return &memoryStudentRepository{
		students: make(map[string]*domain.Student),
		emails:   make(map[string]string),
		logger:   logger,
	}
}

func (r *memoryStudentRepository) Save(ctx context.Context, student *domain.Student) error {
	start := time.Now()
	r.mu.Lock()
	defer r.mu.Unlock()

	r.logger.Debug(ctx, "DB Save Student requested", "student_id", student.ID, "email", student.Email)

	if existingID, exists := r.emails[student.Email]; exists && existingID != student.ID {
		r.logger.Warn(ctx, "DB Conflict: Student email already exists", "email", student.Email)
		return domain.ErrEmailExists
	}

	r.students[student.ID] = student
	r.emails[student.Email] = student.ID

	r.logger.Info(ctx, "DB Student saved", "student_id", student.ID, "duration_ms", time.Since(start).Milliseconds())
	return nil
}

func (r *memoryStudentRepository) FindByID(ctx context.Context, id string) (*domain.Student, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.logger.Debug(ctx, "DB FindByID requested", "student_id", id)

	student, exists := r.students[id]
	if !exists {
		return nil, domain.ErrStudentNotFound
	}
	return student, nil
}

func (r *memoryStudentRepository) FindByEmail(ctx context.Context, email string) (*domain.Student, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	id, exists := r.emails[email]
	if !exists {
		return nil, domain.ErrStudentNotFound
	}
	return r.students[id], nil
}

func (r *memoryStudentRepository) FindAll(ctx context.Context, department string, status string) ([]*domain.Student, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.logger.Debug(ctx, "DB FindAll Students requested", "department_filter", department, "status_filter", status)

	result := make([]*domain.Student, 0)
	for _, s := range r.students {
		if department != "" && !strings.EqualFold(s.Department, department) {
			continue
		}
		if status != "" && !strings.EqualFold(s.Status, status) {
			continue
		}
		result = append(result, s)
	}

	r.logger.Info(ctx, "DB Students fetched", "count", len(result))
	return result, nil
}

func (r *memoryStudentRepository) Update(ctx context.Context, student *domain.Student) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.logger.Debug(ctx, "DB Update Student requested", "student_id", student.ID)

	if _, exists := r.students[student.ID]; !exists {
		return domain.ErrStudentNotFound
	}

	r.students[student.ID] = student
	r.logger.Info(ctx, "DB Student updated", "student_id", student.ID)
	return nil
}

func (r *memoryStudentRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.logger.Debug(ctx, "DB Delete Student requested", "student_id", id)

	student, exists := r.students[id]
	if !exists {
		return domain.ErrStudentNotFound
	}

	delete(r.emails, student.Email)
	delete(r.students, id)
	r.logger.Info(ctx, "DB Student deleted", "student_id", id)
	return nil
}
