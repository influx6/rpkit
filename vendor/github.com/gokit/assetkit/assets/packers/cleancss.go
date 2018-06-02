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

var (
	inGOPATH         = os.Getenv("GOPATH")
	inGOPATHSrc      = filepath.ToSlash(fileJoin(inGOPATH, "src"))
	guSrc            = filepath.ToSlash(fileJoin(inGOPATHSrc, "github.com/gokit/assetkit"))
	guSrcNodeModules = filepath.ToSlash(fileJoin(inGOPATHSrc, "github.com/gokit/assetkit/node_modules"))
	nodeBin          = filepath.ToSlash(fileJoin(inGOPATHSrc, "github.com/gokit/assetkit/node_modules/.bin/"))
)

// CleanCSSPacker defines an implementation for parsing css files.
// WARNING: Requires Nodejs to be installed. Else original file contents will be copied as is.
type CleanCSSPacker struct {
	Args []string
}

// Pack process all files present in the FileStatment slice and returns WriteDirectives
// which conta ins expected outputs for these files.
func (cess CleanCSSPacker) Pack(statements []assets.FileStatement, dir assets.DirStatement) ([]assets.WriteDirective, error) {
	var directives []assets.WriteDirective

	for _, statement := range statements {
		if err := processCleanStatement(statement, cess, &directives); err != nil {
			return nil, err
		}
	}

	return directives, nil
}

func processCleanStatement(statement assets.FileStatement, cess CleanCSSPacker, directives *[]assets.WriteDirective) error {
	args := append([]string{}, cess.Args...)
	args = append(args, statement.AbsPath)

	node, err := gexec.LookPath("node")
	if err != nil {
		return err
	}

	os.Setenv("node", node)

	command := fmt.Sprintf("%s %s", fileJoin(nodeBin, "cleancss"), strings.Join(args, " "))

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
		fmt.Printf("Command Execution Failed: \n%s\n Response: %s\n Command: %s\n", err, errBuf.String(), command)

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
		OriginPath:    statement.Path,
		OriginAbsPath: statement.AbsPath,
		Writer:        bytes.NewReader(outBuf.Bytes()),
	})

	return nil
}
