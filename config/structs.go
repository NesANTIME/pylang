package config

type Application struct {
	Parameters Args
	Path       Paths

	ContentByte ContentBytes
}

type Args struct {
	NameFile string

	ModelCompressed         string
	ExcluideFoldersAndFiles []string
}

type Paths struct {
	Work string

	PylangCompiler string
	PylangBuild    string
}

type ContentBytes struct {
	Code    []byte
	Runtime []byte
}

// ---

type Log struct {
	Origin  string
	Message string

	IsAlert bool
	Err     error

	TimeStart float32
}

type Animations struct {
	Use       bool
	TimeSleep int
}
