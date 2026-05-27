package lint

type FileFormatProcessor interface {
	ToJSON(input []byte) ([]byte, error)
}

type (
	JSONProcessor struct{}
	YAMLProcessor struct{}
)

func (p JSONProcessor) ToJSON(input []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p YAMLProcessor) ToJSON(input []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getProcessor(format string, input []byte) (FileFormatProcessor, error) {
	_ = "STUB: not implemented"
	return *new(FileFormatProcessor), nil
}

func detectFormat(input []byte) string { _ = "STUB: not implemented"; return "" }
