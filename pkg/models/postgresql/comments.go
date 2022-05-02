package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.com/Yolo62442/Comments/pkg/models"
)

const (
	insertSql                 = "INSERT INTO comments (author, comment, parent_id) VALUES ($1,$2,$3) RETURNING id"
	getAll 					  = "SELECT id, author, comment, parent_id FROM comments "
	deleteSql 				  = "DELETE FROM comments WHERE id=$1"
	deleteSqlChildren 		  = "DELETE FROM comments WHERE parent_id=$1"
)

type CommentModel struct {
	Pool *pgxpool.Pool
}

func (m *CommentModel) Insert(author, comment string, parentID int) (int, error) {
	var id uint64
	row := m.Pool.QueryRow(context.Background(), insertSql, author, comment, parentID)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}


func (m *CommentModel) All() ([]*models.Comment, error) {
	comments := []*models.Comment{}
	rows, err := m.Pool.Query(context.Background(), getAll)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		s := &models.Comment{}
		err = rows.Scan(&s.ID, &s.Author, &s.Comments, &s.ParentID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func (m *CommentModel) Delete(id int) error {
	commTag, err := m.Pool.Exec(context.Background(), deleteSql, id)
	if err != nil {
		return err
	}
	if commTag.RowsAffected() == 0 {
		return models.ErrNoRecord
	}
	commTag, err = m.Pool.Exec(context.Background(), deleteSqlChildren, id)
	if err != nil {
		return err
	}
	return nil
}
