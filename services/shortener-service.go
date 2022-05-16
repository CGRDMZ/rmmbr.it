package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/CGRDMZ/rmmbrit-api/models"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type ShortenerService struct {
	Db *pgxpool.Pool
}


func (ss *ShortenerService) CreateNewUrlMap(ctx context.Context, shortUrl, longUrl string) (*models.UrlMap, error) {
	// if no short url is provided, generate one
	if strings.Trim(shortUrl, " ") == "" {
		s, err := gonanoid.New()
		if err != nil {
			return nil, fmt.Errorf("something happened while generating id: %w", err)
		}
		shortUrl = s
	}

	// TODO: refactor this to a repository ------
	tx, err := ss.Db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("something wrong happened while beginning transaction: %w", err)
	}

	_, err = tx.Exec(context.Background(), "INSERT INTO url_map (long_url, short_url) VALUES ($1, $2)", longUrl, shortUrl)
	if err != nil {
		return nil, fmt.Errorf("something happened while creating a new 'url map': %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not commit the transaction: %w", err)
	}
	// -------------------------------------------

	um, _, err2 := ss.findUrlMapByShortUrl(ctx, shortUrl)

	return um, err2

}

func (ss *ShortenerService) GetUrlMapInfo(ctx context.Context, shortUrl string) (*models.UrlMap, error) {
	um, _, err2 := ss.findUrlMapByShortUrl(ctx, shortUrl)

	return um, err2
}

func (ss *ShortenerService) GetUrlMapInfoAndIncrementVisit(ctx context.Context, shortUrl string) (*models.UrlMap, error) {
	um, _, err := ss.findUrlMapByShortUrl(ctx, shortUrl)

	if err != nil {
		return nil, err
	}
	go func() {
		ss.incrementVisited(context.Background(), shortUrl)
	}()
	return um, err
}

// TODO: this method will go to a repository
func (ss *ShortenerService) findUrlMapByShortUrl(ctx context.Context, shortUrl string) (um *models.UrlMap, found bool, err error) {
	sql := "SELECT id, um.short_url, um.long_url, um.visited_count, um.created_at, um.updated_at " +
		"FROM url_map um WHERE short_url = $1"

	urlMap := new(models.UrlMap)
	err = ss.Db.QueryRow(ctx, sql, shortUrl).Scan(
		&urlMap.Id,
		&urlMap.ShortUrl,
		&urlMap.LongUrl,
		&urlMap.VisitedCounter,
		&urlMap.CreatedAt,
		&urlMap.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, false, nil
		} else {
			return nil, false, err
		}
	}

	return urlMap, true, nil

}

func (ss *ShortenerService) incrementVisited(ctx context.Context, shortUrl string) error {
	sql := "UPDATE url_map SET visited_count = visited_count + 1 WHERE short_url = $1"

	_, err := ss.Db.Exec(ctx, sql, shortUrl)
	return err
}