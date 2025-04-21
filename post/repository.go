package post

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type Repository interface {
	Close()
	PutPost(ctx context.Context, p Post) error
	GetPostByID(ctx context.Context, id string) (*Post, error)
	GetPostByName(ctx context.Context, name string) (*Post, error)
	GetPostByDescription(ctx context.Context, description string) (*Post, error)
	GetAllPosts(ctx context.Context) ([]*Post, error)
	DeletePost(ctx context.Context, id string) error
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*postgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() {
	r.db.Close()
}

func (r *postgresRepository) PutPost(ctx context.Context, p Post) (err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	_, err = tx.ExecContext(ctx, `
		INSERT INTO posts(id, created_at, updated_at, name, description, account_id)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		p.ID, p.CreatedAt, p.UpdatedAt, p.Name, p.Description, p.AccountID,
	)
	if err != nil {
		return err
	}
	if len(p.Likes) > 0 {
		stmtLikes, err := tx.PrepareContext(ctx, pq.CopyIn("likes", "id", "post_id", "account_id"))
		if err != nil {
			return err
		}
		defer stmtLikes.Close()

		for _, like := range p.Likes {
			_, err = stmtLikes.ExecContext(ctx, like.ID, like.PostID, like.AccountID)
			if err != nil {
				return err
			}
		}
		_, err = stmtLikes.ExecContext(ctx)
		if err != nil {
			return err
		}
	}

	if len(p.Comments) > 0 {
		stmtComments, err := tx.PrepareContext(ctx, pq.CopyIn("comments", "id", "post_id", "account_id", "content"))
		if err != nil {
			return err
		}
		defer stmtComments.Close()

		for _, c := range p.Comments {
			_, err = stmtComments.ExecContext(ctx, c.ID, c.PostID, c.AccountID, c.Content)
			if err != nil {
				return err
			}
		}
		_, err = stmtComments.ExecContext(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *postgresRepository) GetPostByID(ctx context.Context, id string) (*Post, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id,created_at,updated_at,name,description,account_id FROM posts WHERE id=$1", id)

	p := &Post{}
	err := row.Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt, &p.Name, &p.Description, &p.AccountID)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *postgresRepository) GetPostByName(ctx context.Context, name string) (*Post, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id,created_at,updated_at,name,description,account_id FROM posts WHERE name=$1", name)

	p := &Post{}
	err := row.Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt, &p.Name, &p.Description, &p.AccountID)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *postgresRepository) GetPostByDescription(ctx context.Context, description string) (*Post, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id,created_at,updated_at,name,description,account_id FROM posts WHERE description=$1", description)

	p := &Post{}
	err := row.Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt, &p.Name, &p.Description, &p.AccountID)
	if err != nil {
		return nil, err
	}
	return p, nil
}
func (r *postgresRepository) GetAllPosts(ctx context.Context) ([]*Post, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id,created_at,updated_at,name,description,account_id FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := []*Post{}
	for rows.Next() {
		p := &Post{}
		err = rows.Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt, &p.Name, &p.Description, &p.AccountID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postgresRepository) DeletePost(ctx context.Context, id string) error {
	row, err := r.db.ExecContext(ctx, "DELETE FROM posts WHERE id=$1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no post foudn with id %s", id)
	}
	return nil

}
