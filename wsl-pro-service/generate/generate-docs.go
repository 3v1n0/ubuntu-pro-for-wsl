//go:build tools

package generate

//go:generate go run ../../tools/generate/generate_autocompletion_documentation.go update-readme generate.yaml
//go:generate go run ../../tools/generate/generate_autocompletion_documentation.go update-doc-cli-ref generate.yaml
