package services

type IndexService interface {
    
}

type indexService struct{}

func NewIndexService() IndexService {
    return &indexService{}
}