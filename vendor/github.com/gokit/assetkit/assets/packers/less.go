// +build !js

package packers

import (
	"bytes"
	"context"
	"fmt"
	"os"
	gexec "os/exec"
	"path/filepath"
	"strings"
	"time"

	"io/ioutil"

	"github.com/gokit/assetkit/assets"
	"github.com/gokit/zexec"
)

// LessPacker defines an implementation for parsing .less files into css files using the less compiler in nodejs.
// WARNING: Requires Nodejs to be installed. Else original file contents will be copied as is.
type LessPacker struct {
	MainFile string
	Options  map[string]string
}

// Pack process all files present in the FileStatment slice and returns WriteDirectives
// which conta ins expected outputs for these files.
func (less LessPacker) Pack(statements []assets.FileStatement, dir assets.DirStatement) ([]assets.WriteDirective, error) {
	var directives []assets.WriteDirective

	// If main less file has being set then attempt to find main file.
	if less.MainFile == "" {
		for _, statement := range statements {
			if err := processStatement(statement, less, &directives); err != nil {
				return nil, err
			}
		}

		return directives, nil
	}

	for _, statement := range statements {
		if statement.Path != less.MainFile {
			continue
		}

		if err := processStatement(statement, less, &directives); err != nil {
			return nil, err
		}
	}

	return directives, nil
}

func processStatement(statement assets.FileStatement, less LessPacker, directives *[]assets.WriteDirective) error {
	fileExt := filepath.Ext(statement.Path)
	cssFileName := fileJoin(filepath.Dir(statement.Path), strings.Replace(filepath.Base(statement.Path), fileExt, ".css", 1))
	cssAbsFileName := fileJoin(filepath.Dir(statement.AbsPath), strings.Replace(filepath.Base(statement.Path), fileExt, ".css", 1))

	cssFileName = strings.Replace(cssFileName, "less/", "css/", 1)
	cssAbsFileName = strings.Replace(cssAbsFileName, "less/", "css/", 1)

	var args []string

	for option, value := range less.Options {
		args = append(args, option, value)
	}

	args = append(args, filepath.Clean(statement.AbsPath))

	node, err := gexec.LookPath("node")
	if err != nil {
		return err
	}

	os.Setenv("node", node)

	command := fmt.Sprintf("%s %s", fileJoin(nodeBin, "lessc"), strings.Join(args, " "))

	var errBuf, outBuf bytes.Buffer
	cleanCmd := zexec.New(
		zexec.Async(),
		zexec.Command(command),
		zexec.Output(&outBuf),
		zexec.Err(&errBuf),
	)

	ctx, cancl := context.WithTimeout(context.Background(), time.Minute)
	defer cancl()

	if _, err := cleanCmd.Exec(ctx); err != nil {
		fmt.Printf("Command Execution Failed: %q\n Response: %q\n Command: %q\n", err, errBuf.String(), command)

		content, err := ioutil.ReadFile(statement.AbsPath)
		if err != nil {
			return err
		}

		*directives = append(*directives, assets.WriteDirective{
			OriginPath:    statement.Path,
			OriginAbsPath: statement.AbsPath,
			Writer:        bytes.NewReader(content),
		})
		return nil
	}

	*directives = append(*directives, assets.WriteDirective{
		OriginPath:    cssFileName,
		OriginAbsPath: cssAbsFileName,
		Writer:        bytes.NewReader(outBuf.Bytes()),
	})

	return nil
}

func fileJoin(s ...string) string {
	return filepath.ToSlash(filepath.Join(s...))
}
