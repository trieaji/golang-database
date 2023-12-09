package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang-database/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

// implementasi new repository
func NewCommentRepository(db *sql.DB) CommentRepository { //kasarannya sebagai connection ke db, jadi setelah memenuhi kontrak(interface), maka dilakukan connection ke db
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	fmt.Println("=== print ===")
	fmt.Println(ctx)
	script := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	fmt.Println("=== print2 ===")
	fmt.Println(ctx)
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	fmt.Println("=== print3 ===")
	fmt.Println(ctx)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("Id" + strconv.Itoa(int(id)) + " Not found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil
}
