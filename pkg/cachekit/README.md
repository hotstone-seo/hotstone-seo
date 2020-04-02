# Cachekit

Example
```go
func (r *BookServiceImpl) FindOne(ctx context.Context, id int64) (book *repository.Book, err error) {
	book = new(repository.Book)
	var (
		cacheStore = cachekit.New(r.Redis)
		key        = fmt.Sprintf("BOOK:FIND_ONE:%d", id)
	)
	cacheStore.Retrieve(ctx, key, book, func() (interface{}, error) {
		book, err = r.BookRepo.FindOne(ctx, id)
		return book, err
	})
	return
}
```