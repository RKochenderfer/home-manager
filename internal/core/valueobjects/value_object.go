package valueobjects

type ValueObject interface {
	PartialEq(v ValueObject) bool
}