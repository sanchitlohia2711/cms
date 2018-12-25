package create

import (
	"github.com/ko/cms-db/cms-es/config"
	"github.com/ko/cms-db/cms-es/es"
)

var (
	esClient = es.GetClient()
	conf     = config.New()
)
