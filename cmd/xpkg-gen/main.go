// Package main provides a tool to generate crossplane.yaml files from a template
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type PackageMetadata struct {
	Name               string
	Service            string
	ServiceDisplayName string
	ProviderName       string
	XpkgRegOrg         string
	Version            string
	DepConstraint      string
}

var serviceDisplayNames = map[string]string{
	"config":                  "Configuration",
	"compute":                 "Compute",
	"networking":              "Networking",
	"blockstorage":            "Block Storage",
	"networkconnectivity":     "Network Connectivity",
	"containerengine":         "Container Engine",
	"identity":                "Identity and Access Management",
	"objectstorage":           "Object Storage",
	"loadbalancer":            "Load Balancer",
	"networkloadbalancer":     "Network Load Balancer",
	"dns":                     "DNS",
	"kms":                     "Key Management",
	"functions":               "Functions",
	"logging":                 "Logging",
	"monitoring":              "Monitoring",
	"events":                  "Events",
	"streaming":               "Streaming",
	"filestorage":             "File Storage",
	"artifacts":               "Artifacts",
	"vault":                   "Vault",
	"ons":                     "Notifications",
	"certificatesmanagement":  "Certificates Management",
	"networkfirewall":         "Network Firewall",
	"servicemesh":             "Service Mesh",
	"healthchecks":            "Health Checks",
	"monolith":                "All Services",
}

func main() {
	var (
		service       = flag.String("service", "", "Service name (required)")
		templatePath  = flag.String("template", "package/crossplane.yaml.tmpl", "Template file path")
		outputDir     = flag.String("output", "_output/xpkg", "Output directory")
		version       = flag.String("version", "1.0.0", "Package version")
		xpkgRegOrg    = flag.String("registry", "xpkg.upbound.io/upbound", "Package registry organization")
		depConstraint = flag.String("dep-constraint", ">= 1.0.0", "Dependency constraint")
	)
	flag.Parse()

	if *service == "" {
		fmt.Fprintf(os.Stderr, "Error: -service is required\n")
		flag.Usage()
		os.Exit(1)
	}

	displayName, ok := serviceDisplayNames[*service]
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: Unknown service: %s\n", *service)
		os.Exit(1)
	}

	// Determine package name
	var packageName string
	if *service == "config" {
		packageName = "provider-family-oci"
	} else if *service == "monolith" {
		packageName = "provider-oci"
	} else {
		packageName = fmt.Sprintf("provider-oci-%s", *service)
	}

	// Read template
	tmplContent, err := os.ReadFile(*templatePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading template: %v\n", err)
		os.Exit(1)
	}

	// Parse and execute template
	tmpl, err := template.New("crossplane.yaml").Parse(string(tmplContent))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing template: %v\n", err)
		os.Exit(1)
	}

	metadata := PackageMetadata{
		Name:               packageName,
		Service:            *service,
		ServiceDisplayName: displayName,
		ProviderName:       "oci",
		XpkgRegOrg:         *xpkgRegOrg,
		Version:            *version,
		DepConstraint:      *depConstraint,
	}

	// Create output directory
	serviceOutputDir := filepath.Join(*outputDir, *service)
	if err := os.MkdirAll(serviceOutputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Write output file
	outputPath := filepath.Join(serviceOutputDir, "crossplane.yaml")
	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	if err := tmpl.Execute(outputFile, metadata); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing template: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s\n", outputPath)
}

// Helper function for template
func (p PackageMetadata) IsMonolith() bool {
	return p.Service == "monolith"
}

func (p PackageMetadata) IsConfig() bool {
	return p.Service == "config"
}

func (p PackageMetadata) NeedsFamily() bool {
	return p.Service != "config" && p.Service != "monolith"
}