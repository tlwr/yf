package traverser

type TraversedDocument struct {
	Content         string
	OriginalContent string

	Line int
}

func Traverse(path string, content string) (TraversedDocument, error) {
	return TraversedDocument{
		Content:         content,
		OriginalContent: content,
		Line:            0,
	}, nil
}
