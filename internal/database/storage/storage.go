package storage

import "context"

func (s *Storage) Set(ctx context.Context, key, value string) error {
	s.engine.Set(ctx, key, value)

	return nil
}

func (s *Storage) Get(ctx context.Context, key string) (string, error) {
	value, found := s.engine.Get(ctx, key)
	if !found {
		return "", errNotFound
	}

	return value, nil
}

func (s *Storage) Del(ctx context.Context, key string) error {
	s.engine.Del(ctx, key)

	return nil
}
