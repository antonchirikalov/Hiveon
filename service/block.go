package service

import (
	"hiveon-api/model"
	. "hiveon-api/repository"
)

type BlockService interface {
	GetBlockCount() model.BlockCount
}

type blockService struct {
	hiveosRepository IHiveosRepository
}

func NewBlockService() BlockService {
	return &blockService{hiveosRepository: NewHiveosRepository()}
}

func NewBlockServiceWithRepo(repo IHiveosRepository) BlockService {
	return &blockService{hiveosRepository: repo}
}

func (b *blockService) GetBlockCount() model.BlockCount {
	blockData := model.BlockCount{Code:200}
	blockData.Data.Uncles = b.hiveosRepository.GetBlock24Uncle()
	blockData.Data.Blocks = b.hiveosRepository.GetBlock24NotUnckle()
	return blockData
}


