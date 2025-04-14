package search

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/olivere/elastic/v7"
)

type Repository interface {
	Close()
	PutPost(ctx context.Context, p Post) error
	GetPostByID(ctx context.Context, id string) (*Post, error)
	ListPosts(ctx context.Context, skip uint64, take uint64) ([]Post, error)
	ListPostsWithIDS(ctx context.Context, ids []string) ([]Post, error)
	SearchPosts(ctx context.Context, query string, skip uint64, take uint64) ([]Post, error)
}

type elasticRepository struct {
	client *elastic.Client
}

var ErrNotFound = errors.New("post not found")

type postDocument struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewElasticRepository(url string) (Repository, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	return &elasticRepository{client: client}, nil
}

func (r *elasticRepository) Close() {

}

func (r *elasticRepository) PutPost(ctx context.Context, p Post) error {
	_, err := r.client.Index().
		Index("posts").
		Id(p.ID).
		BodyJson(postDocument{
			Name:        p.Name,
			Description: p.Description,
		}).
		Do(ctx)
	return err
}

func (r *elasticRepository) GetPostByID(ctx context.Context, id string) (*Post, error) {
	res, err := r.client.Get().
		Index("posts").
		Id(id).
		Do(ctx)

	if err != nil {
		return nil, err

	}
	if !res.Found {
		return nil, ErrNotFound
	}
	p := postDocument{}
	if err = json.Unmarshal(res.Source, &p); err != nil {
		return nil, err
	}
	return &Post{
		ID:          id,
		Name:        p.Name,
		Description: p.Description,
	}, err
}

func (r *elasticRepository) ListPosts(ctx context.Context, skip uint64, take uint64) ([]Post, error) {
	res, err := r.client.Search().
		Index("posts").
		Query(elastic.NewMatchAllQuery()).
		From(int(skip)).Size(int(take)).
		Do(ctx)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	posts := []Post{}
	for _, hit := range res.Hits.Hits {
		p := postDocument{}
		if err := json.Unmarshal(hit.Source, &p); err == nil {
			posts = append(posts, Post{
				ID:          hit.Id,
				Name:        p.Name,
				Description: p.Description,
			})
		}
	}
	return posts, err
}
func (r *elasticRepository) ListPostsWithIDS(ctx context.Context, ids []string) ([]Post, error) {
	items := []*elastic.MultiGetItem{}
	for _, id := range ids {
		items = append(items, elastic.NewMultiGetItem().
			Index("posts").
			Id(id))
	}
	res, err := r.client.MultiGet().
		Add(items...).
		Do(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	posts := []Post{}
	for _, doc := range res.Docs {
		p := postDocument{}
		if err := json.Unmarshal(doc.Source, &p); err == nil {
			posts = append(posts, Post{
				ID:          doc.Id,
				Name:        p.Name,
				Description: p.Description,
			})
		}
	}
	return posts, nil
}

func (r *elasticRepository) SearchPosts(ctx context.Context, query string, skip uint64, take uint64) ([]Post, error) {
	res, err := r.client.Search().
		Index("posts").
		Query(elastic.NewMultiMatchQuery(query, "name", "description")).
		From(int(skip)).Size(int(take)).
		Do(ctx)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	posts := []Post{}
	for _, hit := range res.Hits.Hits {
		p := postDocument{}
		if err := json.Unmarshal(hit.Source, &p); err == nil {
			posts = append(posts, Post{
				ID:          hit.Id,
				Name:        p.Name,
				Description: p.Description,
			})
		}
	}
	return posts, err
}
