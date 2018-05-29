package rp

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/gokit/rpkit/static"

	"runtime"

	"github.com/influx6/moz/ast"
	"github.com/influx6/moz/gen"
)

// InterfaceRP implements code generation for creating RPC based API which work with either
// HTTP 1.0 or other underline transport system.
func InterfaceRP(toPackage string, an ast.AnnotationDeclaration, in ast.InterfaceDeclaration, declr ast.PackageDeclaration, pkg ast.Package) ([]gen.WriteDirective, error) {
	packageName := fmt.Sprintf("%srp", strings.ToLower(in.Object.Name.Name))
	jsPackageName := fmt.Sprintf("%srpjsc", strings.ToLower(in.Object.Name.Name))
	jsServerPackageName := fmt.Sprintf("%srpjss", strings.ToLower(in.Object.Name.Name))
	packageFileName := fmt.Sprintf("%s.rp.go", strings.ToLower(in.Object.Name.Name))
	//jsPackageFileName := fmt.Sprintf("%s.client.rp.js", strings.ToLower(in.Object.Name.Name))
	//jsServerPackageFileName := fmt.Sprintf("%s.server.rp.js", strings.ToLower(in.Object.Name.Name))
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

	templFuncs := gen.ToTemplateFuncs(
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
	)

	binding := rpBinding{
		An:                             an,
		Itr:                            in,
		Pkg:                            declr,
		Imports:                        imports,
		EncodingArgs:                   encodingTypes,
		ImplPackageName:                packageName,
		ImplJSClientPackageName:        jsPackageName,
		ImplJSServerPackageName:        jsServerPackageName,
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
	}

	encodingGen := gen.Package(
		gen.Name(packageName),
		gen.Imports(),
		gen.Block(
			gen.SourceTextWith(
				"rpkit:encoding_rpc",
				string(static.MustReadFile("encoding_rpc.tml", true)),
				templFuncs,
				binding,
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
				templFuncs,
				binding,
			),
		),
	)

	packageJSClient := gen.Block(
		gen.SourceTextWith(
			"rpkit:js_package",
			string(static.MustReadFile("interface_rp_js_client.tml", true)),
			templFuncs,
			binding,
		),
	)

	packageJSServer := gen.Block(
		gen.SourceTextWith(
			"rpkit:js_package",
			string(static.MustReadFile("interface_rp_js_server.tml", true)),
			templFuncs,
			binding,
		),
	)

	packageServerJSON := gen.Block(
		gen.SourceTextWith(
			"rpkit:js_package_json",
			string(static.MustReadFile("jsbundle/package.server.json", true)),
			templFuncs,
			binding,
		),
	)

	packageClientJSON := gen.Block(
		gen.SourceTextWith(
			"rpkit:js_package_json",
			string(static.MustReadFile("jsbundle/package.client.json", true)),
			templFuncs,
			binding,
		),
	)

	editorConf := gen.Block(
		gen.Text(
			string(static.MustReadFile("jsbundle/.editorconfig", true)),
		),
	)

	babelrc := gen.Block(
		gen.Text(
			string(static.MustReadFile("jsbundle/.babelrc", true)),
		),
	)

	eslinter := gen.Block(
		gen.Text(
			string(static.MustReadFile("jsbundle/.eslintrc.json", true)),
		),
	)

	wpConfig := gen.Block(
		gen.Text(
			string(static.MustReadFile("jsbundle/webpack.config.js", true)),
		),
	)

	wpConfigPre := gen.Block(
		gen.Text(
			string(static.MustReadFile("jsbundle/config/webpack.preconfig.js", true)),
		),
	)

	wpConfigWeb := gen.Block(
		gen.Text(
			string(static.MustReadFile("jsbundle/config/webpack.config.web.js", true)),
		),
	)
	
	wpServerConfigNode := gen.Block(
		gen.SourceTextWith(
			"rpkit:webpack_config_node_js",
			string(static.MustReadFile("jsbundle/config/webpack.config.node.js", true)),
			templFuncs,
			binding,
		),
	)

	wpClientConfigNode := gen.Block(
		gen.SourceTextWith(
			"rpkit:webpack_config_node_js",
			string(static.MustReadFile("jsbundle/config/webpack.config.node.client.js", true)),
			templFuncs,
			binding,
		),
	)

	jsServerDirectives := []gen.WriteDirective{
		{
			Dir:      jsServerPackageName,
			FileName: "package.json",
			Writer:   packageServerJSON,
		},
		{
			Dir:      jsServerPackageName,
			FileName: ".editorconfig",
			Writer:   editorConf,
		},
		{
			Dir:      jsServerPackageName,
			FileName: ".babelrc",
			Writer:   babelrc,
		},
		{
			Dir:      jsServerPackageName,
			FileName: ".eslintrc.json",
			Writer:   eslinter,
		},
		{
			Dir:      jsServerPackageName,
			FileName: "webpack.config.js",
			Writer:   wpConfig,
		},
		{
			Dir:      filepath.Join(jsServerPackageName, "config"),
			FileName: "webpack.config.node.js",
			Writer:   wpServerConfigNode,
		},
		{
			Dir:      filepath.Join(jsServerPackageName, "config"),
			FileName: "webpack.config.web.js",
			Writer:   wpConfigWeb,
		}, {
			Dir:      filepath.Join(jsServerPackageName, "config"),
			FileName: "webpack.preconfig.js",
			Writer:   wpConfigPre,
		},
		{
			Dir:      filepath.Join(jsServerPackageName, "src"),
			FileName: "app.js",
			Writer:   packageJSServer,
		},
	}

	jsClientDirectives := []gen.WriteDirective{
		{
			Dir:      jsPackageName,
			FileName: "package.json",
			Writer:   packageClientJSON,
		},
		{
			Dir:      jsPackageName,
			FileName: ".editorconfig",
			Writer:   editorConf,
		},
		{
			Dir:      jsPackageName,
			FileName: ".babelrc",
			Writer:   babelrc,
		},
		{
			Dir:      jsPackageName,
			FileName: ".eslintrc.json",
			Writer:   eslinter,
		},
		{
			Dir:      jsPackageName,
			FileName: "webpack.config.js",
			Writer:   wpConfig,
		},
		{
			Dir:      filepath.Join(jsPackageName, "config"),
			FileName: "webpack.config.node.js",
			Writer:   wpClientConfigNode,
		},
		{
			Dir:      filepath.Join(jsPackageName, "config"),
			FileName: "webpack.config.web.js",
			Writer:   wpConfigWeb,
		}, {
			Dir:      filepath.Join(jsPackageName, "config"),
			FileName: "webpack.preconfig.js",
			Writer:   wpConfigPre,
		},
		{
			Dir:      filepath.Join(jsPackageName, "src"),
			FileName: "app.js",
			Writer:   packageJSClient,
		},
	}

	goCode := []gen.WriteDirective{
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
	}

	directives := make([]gen.WriteDirective, 0, len(goCode)+len(jsClientDirectives)+len(jsServerDirectives))
	directives = append(directives, goCode...)

	if an.HasArg("js:client") {
		directives = append(directives, jsClientDirectives...)
	}

	if an.HasArg("js:server") {
		directives = append(directives, jsServerDirectives...)
	}

	return directives, nil
}

type rpBinding struct {
	ServiceName                    string
	IServiceName                   string
	TargetPackage                  string
	ImplPackageName                string
	ImplJSClientPackageName        string
	ImplJSServerPackageName        string
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
