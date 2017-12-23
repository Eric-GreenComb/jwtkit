package jwt

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/influx6/faux/fmtwriter"

	"github.com/gokit/jwtkit/static"
	"github.com/influx6/moz/ast"
	"github.com/influx6/moz/gen"
)

// JWTGen generates a package for jwt authentication and api.
func JWTGen(toPackage string, an ast.AnnotationDeclaration, str ast.StructDeclaration, pkgDeclr ast.PackageDeclaration, pkg ast.Package) ([]gen.WriteDirective, error) {
	packageName := fmt.Sprintf("%sapi", strings.ToLower(str.Object.Name.Name))
	packageFinalPath := filepath.Join(toPackage, packageName)

	contractName := an.Param("Contract")
	if contractName == "" {
		return nil, errors.New("Annotation must have 'Contract' type pointing to struct for SSO Login")
	}

	var contract ast.StructDeclaration
	if action, err := ast.FindStructType(pkgDeclr, contractName); err == nil && action.Declr != nil {
		contract = action
	}

	jwtGen := gen.Block(
		gen.Commentary(
			gen.SourceText(`Package `+packageName+` provides a auto-generated package which contains a jwt restful CRUD API for the specific {{.Object.Name}} struct in package {{.Package}}.`, str),
		),
		gen.Package(
			gen.Name(packageName),
			gen.Imports(
				gen.Import("fmt", ""),
				gen.Import("context", ""),
				gen.Import("net/jwt", ""),
				gen.Import("encoding/json", ""),
				gen.Import("github.com/dimfeld/jwttreemux", ""),
				gen.Import("github.com/influx6/faux/metrics", ""),
				gen.Import("github.com/influx6/faux/jwtutil", ""),
				gen.Import("github.com/influx6/faux/metrics/custom", ""),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("jwt-api.tml", true)),
					gen.ToTemplateFuncs(
						ast.ASTTemplatFuncs,
						template.FuncMap{
							"hasFunc": pkgDeclr.HasFunctionFor,
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

	jwtReadmeGen := gen.Block(
		gen.Block(
			gen.SourceTextWith(
				string(static.MustReadFile("jwt-api-readme.tml", true)),
				gen.ToTemplateFuncs(
					ast.ASTTemplatFuncs,
					template.FuncMap{
						"hasFunc": pkgDeclr.HasFunctionFor,
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

	jwtJSONGen := gen.Block(
		gen.Package(
			gen.Name("fixtures"),
			gen.Imports(
				gen.Import("encoding/json", ""),
				gen.Import(str.Path, ""),
			),
			gen.Block(
				gen.SourceTextWith(
					string(static.MustReadFile("jwt-api-json.tml", true)),
					gen.ToTemplateFuncs(
						ast.ASTTemplatFuncs,
						template.FuncMap{
							"hasFunc": pkgDeclr.HasFunctionFor,
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

	writers := []gen.WriteDirective{
		{
			Writer:   jwtReadmeGen,
			FileName: "readme.md",
			Dir:      packageName,
		},
		{
			Writer:       fmtwriter.New(jwtJSONGen, true, true),
			FileName:     fmt.Sprintf("%s_fixtures.go", packageName),
			Dir:          filepath.Join(packageName, "fixtures"),
			DontOverride: true,
		},
		{
			Writer:   fmtwriter.New(jwtGen, true, true),
			FileName: fmt.Sprintf("%s.go", packageName),
			Dir:      packageName,
		},
	}

	return writers, nil
}
