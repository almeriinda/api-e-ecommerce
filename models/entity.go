package models

import "github.com/gofiber/fiber/v2"

type Entity interface {
	create() error
	update() error
	delete() error
	findBy() error
	findAll() (interface{}, error)
	validations(*fiber.Ctx) error
	ifNotExists(*fiber.Ctx) bool
}

func CreateUsedJob(entity Entity) error {
	return entity.create()
}

func Create(entity Entity, c *fiber.Ctx) error {
	if err := entity.validations(c); err != nil {
		return err
	}
	return entity.create()
}

func Update(entity Entity) error {
	//entity.validations()
	return entity.update()
}

func Delete(entity Entity) error {
	//entity.validations()
	return entity.delete()
}

func FindBy(entity Entity) error {
	return entity.findBy()
}

func IfNotExists(entity Entity, c *fiber.Ctx) bool {
	return entity.ifNotExists(c)
}

func FindAll(entity Entity) (interface{}, error) {
	return entity.findAll()
}
