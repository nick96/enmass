package witheditor

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

// WithEditor opens EDITOR with signature SIGNATURE and comment COMMENT
func WithEditor(editor string, signature string, comment string) (string, error) {
	tmpFh, err := ioutil.TempFile("", "with-editor-*")
	if err != nil {
		return "", err
	}
	defer func() {
		tmpFh.Close()
		os.Remove(tmpFh.Name())
	}()

	commentedComment := "# " + strings.Join(strings.Split(comment, "\n"), "\n# ")
	tplString := `
{{ .Comments }}


{{ .Signature }}
`
	tpl, err := template.New("witheditor").Parse(tplString)
	if err != nil {
		return "", err
	}

	writer := &bytes.Buffer{}
	err = tpl.Execute(writer, struct {
		Comments  string
		Signature string
	}{
		Comments:  commentedComment,
		Signature: signature,
	})

	tmpFh.WriteString(writer.String())

	line := len(strings.Split(comment, "\n"))
	cmd := exec.Command(editor, "+"+string(line), tmpFh.Name())
	if err = cmd.Run(); err != nil {
		return "", err
	}
	if err = cmd.Wait(); err != nil {
		return "", err
	}

	content, err := ioutil.ReadAll(tmpFh)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
