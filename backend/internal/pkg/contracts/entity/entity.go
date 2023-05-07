package entity

const (
	ContractNameSmartDRM = "smartdrm"
)

type Contract struct {
	Name    string `db:"name"`
	Address string `db:"address"`
}
