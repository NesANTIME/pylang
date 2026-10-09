package function

import "pylang/config"

func ReadConfig(pathWork string) (config.Configs, error) {
	var cfg config.Configs

	// --
	excluidePredetermined := []string{".git", "pylang-cache", "__pycache__", ".env", "pylang-config.conf"}
	cfg.ExcluideCompiler = append(cfg.ExcluideCompiler, excluidePredetermined...)
	// --

	return cfg, nil
}
