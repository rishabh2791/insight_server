package repository

import "insight/pkg/domain/entity"

type DeviceRepository interface {
	Create(*entity.Device) (*entity.Device, error)
	Get(string) (*entity.Device, error)
	List(string) ([]entity.Device, error)
	Update(string, *entity.Device) (*entity.Device, error)
}
