package config

type Application struct {
	Name   string
	Path   Paths
	Config Configs

	ContentByte ContentBytes
}

type Paths struct {
	Work        string
	Interpreter string

	PylangCompiler string
	PylangBuild    string
	PylangCache    string
}

type Configs struct {
	SO   string
	ARCH string

	ModelCompressed    string
	VersionInterpreter string

	ExcluideCompiler    []string
	ExcluideInterpreter []string
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
