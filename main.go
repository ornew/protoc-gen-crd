/*
Copyright Arata Furukawa

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	Version = "(unknown)"

	SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
)

func GetCompilerVersion(req *pluginpb.CodeGeneratorRequest) string {
	version := "(unknown)"
	if v := req.GetCompilerVersion(); v != nil {
		version = fmt.Sprintf("v%v.%v.%v", v.GetMajor(), v.GetMinor(), v.GetPatch())
		if s := v.GetSuffix(); s != "" {
			version += "-" + s
		}
	}
	return version
}

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = SupportedFeatures
		genFiles := map[string]*protogen.File{}
		for path, file := range gen.FilesByPath {
			if file.Generate {
				genFiles[path] = file
			}
		}
		// for _, file := range genFiles {
		// 	filename := file.GeneratedFilenamePrefix + "_crd_gen.go"
		// 	_ = gen.NewGeneratedFile(filename, file.GoImportPath)
		// }
		return nil
	})
}
