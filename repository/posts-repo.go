package repository

import "github.com/luisabarbalho/go_cache_pragmatic_review/entity"

// PostRepository interface
type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
