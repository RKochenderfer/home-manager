package valueobjects

type ValueObject interface {
	eq(v *ValueObject) bool
}