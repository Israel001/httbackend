package graph

import "htt/httbackend/graph/model"

type Resolver struct {
	SermonStore map[string]model.Sermon
}