package bot

type FuncOrSlice interface {
	func() []string | []string
}
