package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/rohankapoor/k8s-guardian/pkg/config"
	"github.com/rohankapoor/k8s-guardian/pkg/gates"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

// Validate mode constants
const (
	ValidateModeWarn  = "warn"
	ValidateModeBlock = "block"
	ValidateModeStrict = "strict"
)

type validateOptions struct {
	files       []string
	strict      bool
	output      string
	exitOnError bool
}

var validateOpts validateOptions

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate Kubernetes manifests against safety gates",
	Long: `Validate manifests before deployment to check compliance with safety policies.
Exits with non-zero code if validation fails (unless --strict is false).`,
	Example: `  # Validate a single file
  guardian validate -f deployment.yaml

  # Validate all YAML files in a directory
  guardian validate -f k8s/

  # Strict mode - fails on warnings
  guardian validate -f deployment.yaml --strict`,
	RunE: runValidate,
}

func init() {
	rootCmd.AddCommand(validateCmd)
	
	validateCmd.Flags().StringArrayVarP(&validateOpts.files, "file", "f", []string{}, "manifest file or directory to validate")
	validateCmd.Flags().BoolVarP(&validateOpts.strict, "strict", "s", false, "treat warnings as errors")
	validateCmd.Flags().StringVarP(&validateOpts.output, "output", "o", "text", "output format (text|json|yaml)")
	_ = validateCmd.MarkFlagRequired("file")
}

func runValidate(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load(viper.ConfigFileUsed())
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	validator := gates.NewValidator(cfg.Gates)

	var files []string
	for _, f := range validateOpts.files {
		expanded, err := expandFiles(f)
		if err != nil {
			return fmt.Errorf("failed to expand %s: %w", f, err)
		}
		files = append(files, expanded...)
	}

	if len(files) == 0 {
		return fmt.Errorf("no manifest files found")
	}

	var results []*gates.ValidationResult
	var hasErrors, hasWarnings bool

	for _, file := range files {
		result, err := validateFile(file, validator, cfg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error validating %s: %v\n", file, err)
			hasErrors = true
			continue
		}
		results = append(results, result)
		
		if result.HasErrors() {
			hasErrors = true
		}
		if result.HasWarnings() {
			hasWarnings = true
		}
	}

	printResults(results, validateOpts.output)

	if hasErrors || (validateOpts.strict && hasWarnings) {
		return fmt.Errorf("validation failed")
	}

	fmt.Println(color.GreenString("✓ All validations passed"))
	return nil
}

func expandFiles(path string) ([]string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return []string{path}, nil
	}

	var files []string
	err = filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && (filepath.Ext(p) == ".yaml" || filepath.Ext(p) == ".yml") {
			files = append(files, p)
		}
		return nil
	})

	return files, err
}

func validateFile(file string, validator *gates.Validator, cfg *config.Config) (*gates.ValidationResult, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Parse YAML into K8s objects
	objects, err := parseManifests(data)
	if err != nil {
		return nil, fmt.Errorf("parse error: %w", err)
	}

	result := &gates.ValidationResult{
		File:    file,
		Objects: make([]gates.ObjectResult, 0),
	}

	for _, obj := range objects {
		r := validator.Validate(obj, cfg.Gates)
		result.Objects = append(result.Objects, r...)
	}

	return result, nil
}

func parseManifests(data []byte) ([]runtime.Object, error) {
	// Simplified - in actual implementation would use k8s.io/apimachinery/pkg/util/yaml
	// For this structure, returning placeholder
	return []runtime.Object{}, nil
}

func printResults(results []*gates.ValidationResult, format string) {
	if format == "json" {
		// JSON output
		fmt.Println("{")
		for i, r := range results {
			if i > 0 {
				fmt.Print(",")
			}
			fmt.Printf("  %q: { ... }\n", r.File)
		}
		fmt.Println("}")
		return
	}

	// Text output
	for _, r := range results {
		fmt.Printf("\n📄 %s\n", r.File)
		if len(r.Objects) == 0 {
			fmt.Println("  No Kubernetes objects found")
			continue
		}
		
		for _, obj := range r.Objects {
			if len(obj.Issues) == 0 {
				fmt.Printf("  ✓ %s/%s\n", obj.Kind, obj.Name)
			} else {
				fmt.Printf("  ✗ %s/%s\n", obj.Kind, obj.Name)
				for _, issue := range obj.Issues {
					severity := "⚠️"
					if issue.Severity == gates.SeverityError {
						severity = "❌"
					}
					fmt.Printf("    %s [%s] %s\n", severity, issue.Gate, issue.Message)
				}
			}
		}
	}
}