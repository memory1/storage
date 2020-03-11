package storage

import (
	"fmt"
)

const BasicBlockSize = 1024

type Type int

const (
	BlockStorageType  Type = iota
	FileStorageType
	ObjectStorageType
)

type Storage struct {
	capacity    int
	storageType Type
}

type Lease struct {
	months  int
	storage Storage
}

type Tenant []Lease

func (tenant Tenant) charge() (total int, levels int) {
	index := 0
	price := 0
	exceed := 0

	for ; index < len(tenant); index++ {
		price = 0

		switch tenant[index].storage.storageType {
		case BlockStorageType:
			price += 40
			if tenant[index].storage.capacity > BasicBlockSize {
				exceed = tenant[index].storage.capacity - BasicBlockSize
				price += exceed * tenant[index].months * 3
			}
		case FileStorageType:
			price += 20
			if tenant[index].months > 2 {
				exceed = tenant[index].months - 2
				price += exceed * tenant[index].storage.capacity
			}
		case ObjectStorageType:
			price += 10
			if tenant[index].months > 3 {
				exceed = tenant[index].months - 3
				price += exceed * tenant[index].storage.capacity * 2
			}
		default:
		}

		total += price

		if tenant[index].storage.storageType == ObjectStorageType && tenant[index].months > 12 {
			levels += 1
		}
	}
	fmt.Println(total, ", levels:", levels)
	return total, levels
}
