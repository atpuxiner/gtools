package grapi

import (
	"encoding/json"
	"github.com/atpuxiner/gtools/gtcli/internal/gterr"
	"github.com/spf13/cobra"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type newParams struct {
	proj      string
	mod       string
	dir       string
	grapiData []struct {
		Path    string `json:"path"`
		IsDir   bool   `json:"isDir"`
		Content []byte `json:"content"`
	}
}

type newHandler struct {
	cmd  *cobra.Command
	args []string
	pars newParams
}

func (r *newHandler) init() {
	r.pars.proj, _ = r.cmd.Flags().GetString("proj")
	reProj := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]{0,64}$`)
	if !reProj.MatchString(r.pars.proj) {
		gterr.LogFatalf(gterr.ErrBadParam, "proj is required "+`^[a-zA-Z][a-zA-Z0-9_-]{0,64}$`)
	}
	r.pars.mod, _ = r.cmd.Flags().GetString("mod")
	reMod := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_./]{0,128}$`)
	if !reMod.MatchString(r.pars.mod) {
		gterr.LogFatalf(gterr.ErrBadParam, "mod is required "+`^[a-zA-Z][a-zA-Z0-9_./]{0,128}$`)
	}
	r.pars.mod = strings.TrimRight(r.pars.mod, "./")
	r.pars.dir, _ = r.cmd.Flags().GetString("dir")
	err := json.Unmarshal([]byte(grapiJson), &r.pars.grapiData)
	gterr.CheckErr(err, gterr.ErrJsonParse)
}

func (r *newHandler) handle() {
	rootPath := filepath.Join(r.pars.dir, r.pars.proj, r.pars.grapiData[0].Path)
	_, err := os.Stat(rootPath)
	if err == nil {
		gterr.LogFatalf(gterr.ErrProjExisted, rootPath+" already existed")
	}
	for _, info := range r.pars.grapiData {
		newPath := filepath.Join(r.pars.dir, r.pars.proj, info.Path)
		if info.IsDir {
			err := os.MkdirAll(newPath, fs.ModePerm)
			gterr.CheckErr(err, gterr.ErrMakeDir)
		} else {
			// replace
			newContent := strings.ReplaceAll(string(info.Content), grapiLink, r.pars.mod)
			if regexp.MustCompile(`^(.*[\\/])cmd([\\/])root\.go$`).MatchString(info.Path) {
				newContent = strings.ReplaceAll(newContent, grapiCmd, r.pars.proj)
			}
			if regexp.MustCompile(`^([\\/])main\.go$`).MatchString(info.Path) {
				newContent = strings.ReplaceAll(newContent, "@title grapi", "@title "+r.pars.proj)
			}
			if regexp.MustCompile(`^([\\/])README\.md$`).MatchString(info.Path) {
				newContent = strings.ReplaceAll(newContent, "# grapi", "# grapi ( => yourProj)")
			}
			err := os.WriteFile(newPath, []byte(newContent), fs.ModePerm)
			gterr.CheckErr(err, gterr.ErrCreateFile)
		}
	}
	log.Printf("%s written to: %s\n", r.pars.proj, rootPath)
}

func NewRun(cmd *cobra.Command, args []string) {
	h := newHandler{
		cmd:  cmd,
		args: args,
	}
	h.init()
	h.handle()
}
