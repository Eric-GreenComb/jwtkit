package jwt

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gokit/jwtkit/static"
	"github.com/influx6/moz/ast"
	"github.com/influx6/moz/gen"
	"runtime"
)

// JWTGen generates a package for jwt authentication and api.
func JWTGen(toPackage string, an ast.AnnotationDeclaration, str ast.StructDeclaration, pkgDeclr ast.PackageDeclaration, pkg ast.Package) ([]gen.WriteDirective, error) {
	packageName := fmt.Sprintf("%sjwt", strings.ToLower(str.Object.Name.Name))
	packageFinalPath := filepath.Join(toPackage, packageName)

	contractName := an.Param("Contract")
	if contractName == "" {
		return nil, errors.New("annotation must have 'Contract' type for struct SSO Login claim")
	}

	var contract ast.StructDeclaration
	action, ok := pkg.StructFor(contractName)
	if !ok {
		return nil, errors.New(`Struct of type "` + contractName + `" not found`)
	}

	if action.Declr != nil {
		contract = action
	}

	jwtGen := gen.Block(
		gen.Commentary(
			gen.SourceText("jwt.packageTitle", `Package `+packageName+` provides a auto-generated package which contains a API for authentication using JWT.`, str),
		),
		gen.Package(
			gen.Name(packageName),
			gen.Imports(
				gen.Import("time", ""),
				gen.Import("errors", ""),
				gen.Import("strings", ""),
				gen.Import("encoding/base64", ""),
				gen.Import("context", ""),
				gen.Import("github.com/dgrijalva/jwt-go", "jwt"),
				gen.Import("github.com/rs/xid", ""),
				gen.Import("github.com/gokit/tokens", ""),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					"jwt.api",
					string(static.MustReadFile("jwt-api.tml", true)),
					gen.ToTemplateFuncs(
						ast.ASTTemplatFuncs,
						template.FuncMap{
							"hasFunc":       pkgDeclr.HasFunctionFor,
							"mapRandomJSON": ast.MapOutFieldsWithRandomValuesToJSON,
						},
					),
					struct {
						Pkg      *ast.PackageDeclaration
						Struct   ast.StructDeclaration
						Contract ast.StructDeclaration
					}{
						Pkg:      &pkgDeclr,
						Struct:   str,
						Contract: contract,
					},
				),
			),
		),
	)

	jwtMockGen := gen.Block(
		gen.Package(
			gen.Name("mock"),
			gen.Imports(
				gen.Import("fmt", ""),
				gen.Import("strings", ""),
				gen.Import("errors", ""),
				gen.Import("context", ""),
				gen.Import("github.com/gokit/tokens", ""),
				gen.Import(join(toPackage, packageName), packageName),
			),
			gen.Block(
				gen.SourceTextWith(
					"jwt.mock",
					string(static.MustReadFile("jwt-mock.tml", true)),
					gen.ToTemplateFuncs(
						ast.ASTTemplatFuncs,
						template.FuncMap{
							"hasFunc":       pkgDeclr.HasFunctionFor,
							"mapRandomJSON": ast.MapOutFieldsWithRandomValuesToJSON,
						},
					),
					struct {
						PackageName string
						PackagePath string
						Pkg         *ast.PackageDeclaration
						Struct      ast.StructDeclaration
						Contract    ast.StructDeclaration
					}{
						Pkg:         &pkgDeclr,
						Struct:      str,
						Contract:    contract,
						PackagePath: packageFinalPath,
						PackageName: packageName,
					},
				),
			),
		),
	)

	jwtTestGen := gen.Block(
		gen.Package(
			gen.Name(fmt.Sprintf("%s_test", packageName)),
			gen.Imports(
				gen.Import("fmt", ""),
				gen.Import("testing", ""),
				gen.Import("time", ""),
				gen.Import("encoding/json", ""),
				gen.Import("context", ""),
				gen.Import("github.com/influx6/faux/tests", ""),
				gen.Import("github.com/dgrijalva/jwt-go", "jwt"),
				gen.Import(join(toPackage, packageName), packageName),
				gen.Import(join(toPackage, packageName, "mock"), ""),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					"jwt.test",
					string(static.MustReadFile("jwt-test.tml", true)),
					gen.ToTemplateFuncs(
						ast.ASTTemplatFuncs,
						template.FuncMap{
							"hasFunc":       pkgDeclr.HasFunctionFor,
							"mapRandomJSON": ast.MapOutFieldsWithRandomValuesToJSON,
						},
					),
					struct {
						PackageName string
						PackagePath string
						Pkg         *ast.PackageDeclaration
						Struct      ast.StructDeclaration
						Contract    ast.StructDeclaration
					}{
						Pkg:         &pkgDeclr,
						Struct:      str,
						Contract:    contract,
						PackagePath: packageFinalPath,
						PackageName: packageName,
					},
				),
			),
		),
	)

	//jwtHTTPGen := gen.Block(
	//	gen.Package(
	//		gen.Name(packageName),
	//		gen.Imports(
	//			gen.Import("errors", ""),
	//			gen.Import("fmt", ""),
	//			gen.Import("context", ""),
	//			gen.Import("net/http", ""),
	//			gen.Import("encoding/json", ""),
	//			gen.Import("github.com/influx6/faux/metrics", ""),
	//			gen.Import("github.com/influx6/faux/httputil", ""),
	//			gen.Import(str.Path, ""),
	//		),
	//		gen.Block(
	//			gen.SourceTextWith(
	//				string(static.MustReadFile("jwt-api-http.tml", true)),
	//				gen.ToTemplateFuncs(
	//					ast.ASTTemplatFuncs,
	//					template.FuncMap{
	//						"hasFunc":       pkgDeclr.HasFunctionFor,
	//						"mapRandomJSON": ast.MapOutFieldsWithRandomValuesToJSON,
	//					},
	//				),
	//				struct {
	//					Pkg      *ast.PackageDeclaration
	//					Struct   ast.StructDeclaration
	//					Contract ast.StructDeclaration
	//				}{
	//					Pkg:      &pkgDeclr,
	//					Struct:   str,
	//					Contract: contract,
	//				},
	//			),
	//		),
	//	),
	//)

	jwtReadmeGen := gen.Block(
		gen.Block(
			gen.SourceTextWith(
				"jwt.readme",
				string(static.MustReadFile("jwt-api-readme.tml", true)),
				gen.ToTemplateFuncs(
					ast.ASTTemplatFuncs,
					template.FuncMap{
						"hasFunc":       pkgDeclr.HasFunctionFor,
						"mapRandomJSON": ast.MapOutFieldsWithRandomValuesToJSON,
					},
				),
				struct {
					PackageName string
					PackagePath string
					Pkg         *ast.PackageDeclaration
					Struct      ast.StructDeclaration
					Contract    ast.StructDeclaration
				}{
					PackagePath: packageFinalPath,
					PackageName: packageName,
					Pkg:         &pkgDeclr,
					Struct:      str,
					Contract:    contract,
				},
			),
		),
	)

	writers := []gen.WriteDirective{
		{
			Writer:   jwtReadmeGen,
			FileName: "readme.md",
			Dir:      packageName,
		},
		{
			Writer:   jwtGen,
			FileName: fmt.Sprintf("%s.go", packageName),
			Dir:      packageName,
		},
		{
			Writer:   jwtTestGen,
			FileName: fmt.Sprintf("%s_test.go", packageName),
			Dir:      packageName,
		},
		{
			Writer:   jwtMockGen,
			FileName: fmt.Sprintf("%s.go", packageName),
			Dir:      filepath.Join(packageName, "mock"),
		},
		//{
		//	Writer:   (jwtHTTPGen),
		//	FileName: fmt.Sprintf("%s_api.go", packageName),
		//	Dir:      packageName,
		//},
	}

	return writers, nil
}

func join(s ...string) string {
	ss := filepath.Join(s...)
	if runtime.GOOS == "windows" {
		return filepath.ToSlash(ss)
	}
	return ss
}
