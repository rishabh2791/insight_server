package repository

import "insight/pkg/domain/entity"

type DeviceDataRepository interface {
	Create(*entity.DeviceData) (*entity.DeviceData, error)
	Get(string) (*entity.DeviceData, error)
	List(string) ([]entity.DeviceData, error)
}
