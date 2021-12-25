package filesystem

import "os"

func Touch(path string, opts ...TouchOption) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, opt := range opts {
		if err := opt(f); err != nil {
			return err
		}
	}

	return nil
}

func WithContentsOfString(s string) TouchOption {
	return func(f *os.File) error {
		_, err := f.WriteString(s)

		return err
	}
}

type TouchOption func(f *os.File) error
