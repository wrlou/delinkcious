package news_manager

import (
	om "github.com/wrlou/delinkcious/pkg/object_model"
)

type Store interface {
	GetNews(username string, startIndex int) (events []*om.LinkManagerEvent, nextIndex int, err error)
	AddEvent(username string, event *om.LinkManagerEvent) (err error)
}
