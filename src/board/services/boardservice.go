package board

import (
	. "../model"
	"errors"
)

type BoardService struct {
	Model BoardModel
}

func NewBoardService(model BoardModel) *BoardService {
	return &BoardService{Model: *model.Fresh()}
}

func (service BoardService) FindByUUID(uuid string) (Board *BoardModel, err error) {
	Board = service.Model.Fresh()

	service.Model.Connection.Where("id = ?", uuid).Find(&Board)
	if Board.ID == "" {
		err = errors.New("board not found")
	}

	return Board, err
}
