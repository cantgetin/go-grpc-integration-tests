package bookrepository

import "context"

func (r *RepositoryImpl) GetBook(ctx context.Context, id int32) (*Book, error) {
	var book *Book

	err := r.db.First(&book, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *RepositoryImpl) ListBooks(ctx context.Context) ([]*Book, error) {
	var books []*Book

	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
