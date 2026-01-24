package value_object

type Tariff int

const (
	TariffUnknown Tariff = 0

	TariffBase Tariff = 1

	TariffMax Tariff = 2
)
