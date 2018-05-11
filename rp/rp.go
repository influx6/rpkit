package rp

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/gokit/rpkit/static"

	"github.com/influx6/moz/ast"
	"github.com/influx6/moz/gen"
	"runtime"
)

// InterfaceRP implements code generation for creating RPC based API which work with either
// HTTP 1.0 or other underline transport system.
func InterfaceRP(toPackage string, an ast.AnnotationDeclaration, in ast.InterfaceDeclaration, declr ast.PackageDeclaration, pkg ast.Package) ([]gen.WriteDirective, error) {
	packageName := fmt.Sprintf("%srp", strings.ToLower(in.Object.Name.Name))
	packageFileName := fmt.Sprintf("%s.rp.go", strings.ToLower(in.Object.Name.Name))
	packagePath := join(toPackage, packageName)

	sections := strings.Split(strings.TrimSuffix(declr.Path, "/"), "/")
	sections = sections[1:]
	serviceSection := sections[len(sections)-1:]

	var usesInternalPackage bool
	var noArgNoReturnMethods,
		onlyErrorMethods,
		outputWithNoErrorMethods,
		outputWithErrorMethods,
		inputWithOnlyErrorMethods,
		inputWithOutputOnlyMethods,
		inputWithOutputWithErrorMethods []ast.FunctionDefinition

	packagePrefix := fmt.Sprintf("%s.", declr.Package)

	methods := in.Methods(&declr)
methodLoop:
	for _, method := range methods {
		for _, arg := range method.Args {
			if strings.HasPrefix(arg.ExType, packagePrefix) {
				usesInternalPackage = true
				break methodLoop
			}
		}
		for _, arg := range method.Returns {
			if strings.HasPrefix(arg.ExType, packagePrefix) {
				usesInternalPackage = true
				break methodLoop
			}
		}
	}

	encodingTypes := map[string]ast.ArgType{}

	imports := in.GetImports(&declr, false)
	for _, method := range methods {
		if !method.HasReturns() && !method.HasArgs() {
			noArgNoReturnMethods = append(noArgNoReturnMethods, method)
			continue
		}

		if method.TotalArgs() == 1 && method.HasArgType("context.Context") && !method.HasReturns() {
			noArgNoReturnMethods = append(noArgNoReturnMethods, method)
			continue
		}

		if !method.HasArgs() && method.HasReturnType("error") && method.TotalReturns() == 1 {
			onlyErrorMethods = append(onlyErrorMethods, method)
			continue
		}

		if method.TotalArgs() == 1 && method.HasArgType("context.Context") &&
			method.HasReturnType("error") && method.TotalReturns() == 1 {
			onlyErrorMethods = append(onlyErrorMethods, method)
			continue
		}

		if !method.HasArgs() && !method.HasReturnType("error") && method.TotalReturns() == 1 {
			arg := method.Returns[0]
			if skipArg(arg) {
				continue
			}
			encodingTypes[arg.ExType] = arg
			outputWithNoErrorMethods = append(outputWithNoErrorMethods, method)
			continue
		}

		if method.TotalArgs() == 1 && method.HasArgType("context.Context") &&
			!method.HasReturnType("error") && method.TotalReturns() == 1 {
			arg := method.Returns[0]
			if skipArg(arg) {
				continue
			}
			encodingTypes[arg.ExType] = arg
			outputWithNoErrorMethods = append(outputWithNoErrorMethods, method)
			continue
		}

		if !method.HasArgs() && method.HasReturnType("error") &&
			method.TotalReturns() == 2 && method.ReturnTypePos("error") == 1 {
			arg := method.Returns[0]
			if skipArg(arg) {
				continue
			}
			encodingTypes[arg.ExType] = arg
			outputWithErrorMethods = append(outputWithErrorMethods, method)
			continue
		}

		if method.TotalArgs() == 1 && method.HasArgType("context.Context") &&
			method.HasReturnType("error") &&
			method.TotalReturns() == 2 && method.ReturnTypePos("error") == 1 {
			arg := method.Returns[0]
			if skipArg(arg) {
				continue
			}
			encodingTypes[arg.ExType] = arg
			outputWithErrorMethods = append(outputWithErrorMethods, method)
			continue
		}

		if method.TotalArgs() == 1 && method.TotalReturns() == 1 &&
			!method.HasArgType("context.Context") && method.HasReturnType("error") {
			arg := method.Args[0]
			if skipArg(arg) {
				continue
			}
			encodingTypes[arg.ExType] = arg
			inputWithOnlyErrorMethods = append(inputWithOnlyErrorMethods, method)
			continue
		}

		if method.TotalArgs() == 2 && method.TotalReturns() == 1 &&
			method.ArgTypePos("context.Context") == 0 && method.HasReturnType("error") {
			arg := method.Args[1]
			if skipArg(arg) {
				continue
			}
			encodingTypes[arg.ExType] = arg
			inputWithOnlyErrorMethods = append(inputWithOnlyErrorMethods, method)
			continue
		}

		if method.TotalArgs() == 1 && method.TotalReturns() == 1 &&
			!method.HasArgType("context.Context") && !method.HasReturnType("error") {
			arg := method.Args[0]
			if skipArg(arg) {
				continue
			}

			encodingTypes[arg.ExType] = arg
			arg2 := method.Returns[0]
			if skipArg(arg2) {
				continue
			}
			encodingTypes[arg2.ExType] = arg2
			inputWithOutputOnlyMethods = append(inputWithOutputOnlyMethods, method)
			continue
		}

		if method.TotalArgs() == 2 && method.TotalReturns() == 1 &&
			method.ArgTypePos("context.Context") == 0 && !method.HasReturnType("error") {

			// if both return types are error skip, this is bad signature for now.
			if method.CountOfReturnType("error") > 1 {
				continue
			}

			arg := method.Args[1]
			if skipArg(arg) {
				continue
			}
			encodingTypes[arg.ExType] = arg

			arg2 := method.Returns[0]
			if skipArg(arg2) {
				continue
			}
			encodingTypes[arg2.ExType] = arg2
			inputWithOutputOnlyMethods = append(inputWithOutputOnlyMethods, method)
			continue
		}

		if method.TotalArgs() == 1 && method.TotalReturns() == 2 &&
			!method.HasArgType("context.Context") && method.ReturnTypePos("error") == 1 {

			// if both return types are error skip, this is bad signature for now.
			if method.CountOfReturnType("error") > 1 {
				continue
			}

			arg := method.Args[0]
			if skipArg(arg) {
				continue
			}

			encodingTypes[arg.ExType] = arg
			arg2 := method.Returns[0]
			if skipArg(arg2) {
				continue
			}

			encodingTypes[arg2.ExType] = arg2
			inputWithOutputWithErrorMethods = append(inputWithOutputWithErrorMethods, method)
			continue
		}

		if method.TotalArgs() == 2 && method.TotalReturns() == 2 &&
			method.ArgTypePos("context.Context") == 0 && method.ReturnTypePos("error") == 1 {

			// if both return types are error skip, this is bad signature for now.
			if method.CountOfArgType("context.Context") > 1 {
				continue
			}

			// if both return types are error skip, this is bad signature for now.
			if method.CountOfReturnType("error") > 1 {
				continue
			}

			arg := method.Args[1]
			if skipArg(arg) {
				continue
			}

			encodingTypes[arg.ExType] = arg
			arg2 := method.Returns[0]
			if skipArg(arg2) {
				continue
			}

			encodingTypes[arg2.ExType] = arg2
			inputWithOutputWithErrorMethods = append(inputWithOutputWithErrorMethods, method)
			continue
		}
	}

	encodingGen := gen.Package(
		gen.Name(packageName),
		gen.Imports(),
		gen.Block(
			gen.SourceTextWith(
				"rpkit:encoding_rpc",
				string(static.MustReadFile("encoding_rpc.tml", true)),
				gen.ToTemplateFuncs(
					ast.ASTTemplatFuncs,
					template.FuncMap{
						"getTypeName": func(val string) string {
							if parts := strings.Split(val, "."); len(parts) > 1 {
								val = parts[1]
							}

							if len(val) == 1 {
								if val == strings.ToLower(val) {
									return strings.ToUpper(val)
								}
								return val
							}

							if val[:1] == strings.ToLower(val[:1]) {
								return strings.ToUpper(val[:1]) + val[1:]
							}
							return val
						},
					},
				),
				struct {
					ServiceName                    string
					IServiceName                   string
					TargetPackage                  string
					ImplPackageName                string
					UsesInternal                   bool
					Imports                        map[string]string
					EncodingArgs                   map[string]ast.ArgType
					An                             ast.AnnotationDeclaration
					Itr                            ast.InterfaceDeclaration
					Pkg                            ast.PackageDeclaration
					NoArgAndReturns                []ast.FunctionDefinition
					OnlyErrorMethods               []ast.FunctionDefinition
					OutputWithErrorMethods         []ast.FunctionDefinition
					OutputWithNoErrorMethods       []ast.FunctionDefinition
					InputWithErrorMethods          []ast.FunctionDefinition
					InputAndOutputMethods          []ast.FunctionDefinition
					InputAndOutputWithErrorMethods []ast.FunctionDefinition
				}{
					An:                             an,
					Itr:                            in,
					Pkg:                            declr,
					Imports:                        imports,
					EncodingArgs:                   encodingTypes,
					ImplPackageName:                packageName,
					TargetPackage:                  packagePath,
					OnlyErrorMethods:               onlyErrorMethods,
					UsesInternal:                   usesInternalPackage,
					NoArgAndReturns:                noArgNoReturnMethods,
					OutputWithNoErrorMethods:       outputWithNoErrorMethods,
					OutputWithErrorMethods:         outputWithErrorMethods,
					InputWithErrorMethods:          inputWithOnlyErrorMethods,
					InputAndOutputMethods:          inputWithOutputOnlyMethods,
					InputAndOutputWithErrorMethods: inputWithOutputWithErrorMethods,
					ServiceName:                    strings.Join(serviceSection, "."),
					IServiceName:                   strings.Join(serviceSection, "/"),
				},
			),
		),
	)

	iGen := gen.Package(
		gen.Name(packageName),
		gen.Imports(),
		gen.Block(
			gen.SourceTextWith(
				"rpkit:interface_rpc",
				string(static.MustReadFile("interface_rpc.tml", true)),
				gen.ToTemplateFuncs(
					ast.ASTTemplatFuncs,
					template.FuncMap{
						"getTypeName": func(val string) string {
							if parts := strings.Split(val, "."); len(parts) > 1 {
								val = parts[1]
							}

							if len(val) == 1 {
								if val == strings.ToLower(val) {
									return strings.ToUpper(val)
								}
								return val
							}

							if val[:1] == strings.ToLower(val[:1]) {
								return strings.ToUpper(val[:1]) + val[1:]
							}
							return val
						},
					},
				),
				struct {
					ServiceName                    string
					IServiceName                   string
					TargetPackage                  string
					ImplPackageName                string
					UsesInternal                   bool
					EncodingArgs                   map[string]ast.ArgType
					Imports                        map[string]string
					An                             ast.AnnotationDeclaration
					Itr                            ast.InterfaceDeclaration
					Pkg                            ast.PackageDeclaration
					NoArgAndReturns                []ast.FunctionDefinition
					OnlyErrorMethods               []ast.FunctionDefinition
					OutputWithErrorMethods         []ast.FunctionDefinition
					OutputWithNoErrorMethods       []ast.FunctionDefinition
					InputWithErrorMethods          []ast.FunctionDefinition
					InputAndOutputMethods          []ast.FunctionDefinition
					InputAndOutputWithErrorMethods []ast.FunctionDefinition
				}{
					An:                             an,
					Itr:                            in,
					Pkg:                            declr,
					Imports:                        imports,
					EncodingArgs:                   encodingTypes,
					ImplPackageName:                packageName,
					TargetPackage:                  packagePath,
					OnlyErrorMethods:               onlyErrorMethods,
					UsesInternal:                   usesInternalPackage,
					NoArgAndReturns:                noArgNoReturnMethods,
					OutputWithNoErrorMethods:       outputWithNoErrorMethods,
					OutputWithErrorMethods:         outputWithErrorMethods,
					InputWithErrorMethods:          inputWithOnlyErrorMethods,
					InputAndOutputMethods:          inputWithOutputOnlyMethods,
					InputAndOutputWithErrorMethods: inputWithOutputWithErrorMethods,
					ServiceName:                    strings.Join(serviceSection, "."),
					IServiceName:                   strings.Join(serviceSection, "/"),
				},
			),
		),
	)

	return []gen.WriteDirective{
		{
			Dir:      packageName,
			FileName: packageFileName,
			Writer:   iGen,
		},
		{
			Dir:      packageName,
			FileName: "encoding.rp.go",
			Writer:   encodingGen,
		},
	}, nil
}

func skipArg(arg ast.ArgType) bool {
	if arg.ChanType != nil {
		return true
	}

	return false
}

func join(s ...string) string {
	ss := filepath.Join(s...)
	if runtime.GOOS == "windows" {
		return filepath.ToSlash(ss)
	}
	return ss
}
