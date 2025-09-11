package week1

type CarBuilder struct {
	car *Car
}

func (builder *CarBuilder) SetBrand(brand string) *CarBuilder {
	builder.car.brand = brand
	return builder
}

func (builder *CarBuilder) SetYear(year int) *CarBuilder {
	builder.car.year = year
	return builder
}

func (builder *CarBuilder) Build() *Car { return builder.car }

func NewCarBuilder() *CarBuilder { return &CarBuilder{car: &Car{}} }
