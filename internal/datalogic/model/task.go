package model

import (
	outboundModel "base/internal/outbound/db/model"
	usecaseModel "base/internal/usecase/model"
)

type Task struct {
	ID     int64
	Title  string
	Status int
}

func (data Task) MapToUsecase() usecaseModel.Task {
	return usecaseModel.Task{
		ID:     data.ID,
		Title:  data.Title,
		Status: data.Status,
	}
}

func (data *Task) MapFromOutbound(outboundData outboundModel.Task) {
	*data = Task{
		ID:     outboundData.ID,
		Title:  outboundData.Title,
		Status: outboundData.Status,
	}
}
