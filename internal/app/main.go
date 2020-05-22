package app

import (
	"github.com/hotstone-seo/hotstone-seo/internal/urlstore"
)

// Main function to run server
func Main(s server, m urlstore.Worker) (err error) {
	if err = m.Start(); err != nil {
		return
	}
	if err = startServer(s); err != nil {
		return
	}
	return
}
