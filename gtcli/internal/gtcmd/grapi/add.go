package grapi

import (
	"bufio"
	"errors"
	"github.com/atpuxiner/gtools/gtcli/internal/gtcmd/grapi/addtpl"
	"github.com/atpuxiner/gtools/gtcli/internal/gterr"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

type addParams struct {
	api string
	ver string
}

type addHandler struct {
	cmd  *cobra.Command
	args []string
	pars addParams
}

func (r *addHandler) init() {
	r.pars.api, _ = r.cmd.Flags().GetString("api")
	reApi := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{0,64}$`)
	if !reApi.MatchString(r.pars.api) {
		gterr.LogFatalf(gterr.ErrBadParam, "api is required "+`^[a-zA-Z][a-zA-Z0-9_]{0,64}$`)
	}
	r.pars.ver, _ = r.cmd.Flags().GetString("ver")
	reVer := regexp.MustCompile(`^v[a-zA-Z0-9_]{0,10}$`)
	if !reVer.MatchString(r.pars.ver) {
		gterr.LogFatalf(gterr.ErrBadParam, "ver is required "+`^v[a-zA-Z0-9_]{0,10}$`)
	}
}

func (r *addHandler) handle() {
	workDir, _ := os.Getwd()
	modName, err := r.getModName(workDir)
	if err != nil {
		gterr.LogFatalf(gterr.ErrModNotExist, err)
	}
	// check
	for k, _ := range addtpl.Tpls {
		checkPath := filepath.Join(workDir, strings.ReplaceAll(k, "tpl", r.pars.api))
		checkPath = strings.ReplaceAll(checkPath, "vn", r.pars.ver)
		checkDir := filepath.Dir(checkPath)
		if strings.Contains(k, "business") || strings.Contains(k, "entity") {
			checkDir = filepath.Dir(checkDir)
		}
		if _, err := os.Stat(checkDir); err != nil {
			if os.IsNotExist(err) {
				gterr.LogFatalf(gterr.ErrProjCheck, checkDir+" does not exist")
			}
			gterr.LogFatalf(gterr.ErrProjCheck, err)
		}
		_, err := os.Stat(checkPath)
		if err == nil {
			gterr.LogFatalf(gterr.ErrProjCheck, checkPath+" already existed")
		} else if !os.IsNotExist(err) {
			gterr.LogFatalf(gterr.ErrProjCheck, err)
		}
	}
	// add
	for k, v := range addtpl.Tpls {
		newPath := filepath.Join(workDir, strings.ReplaceAll(k, "tpl", r.pars.api))
		newPath = strings.ReplaceAll(newPath, "vn", r.pars.ver)
		newDir := filepath.Dir(newPath)
		// check
		if strings.Contains(k, "business") || strings.Contains(k, "entity") {
			_, err := os.Stat(newDir)
			if err != nil {
				if os.IsNotExist(err) {
					err = os.MkdirAll(newDir, fs.ModePerm)
					if err != nil {
						gterr.LogFatalf(gterr.ErrMakeDir, err)
					}
				} else {
					gterr.LogFatalf(gterr.ErrProjCheck, err)
				}
			}
		}
		newContent := strings.TrimLeft(v, "\r\n")
		newContent = strings.ReplaceAll(newContent, grapiLink, modName)
		newContent = strings.ReplaceAll(newContent, "tpl", r.pars.api)
		newContent = strings.ReplaceAll(newContent, "Tpl", string(unicode.ToUpper(rune(r.pars.api[0])))+r.pars.api[1:])
		newContent = strings.ReplaceAll(newContent, "vn", r.pars.ver)
		newContent = strings.ReplaceAll(newContent, "Vn", string(unicode.ToUpper(rune(r.pars.ver[0])))+r.pars.ver[1:])
		err := os.WriteFile(newPath, []byte(newContent), fs.ModePerm)
		if err != nil {
			gterr.LogFatalf(gterr.ErrCreateFile, err)
		}
	}
}

func (r *addHandler) getModName(modDir string) (string, error) {
	modPath := filepath.Join(modDir, "go.mod")
	file, err := os.Open(modPath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}
	return "", errors.New(modPath + " file does not specify a module")
}

func AddRun(cmd *cobra.Command, args []string) {
	h := addHandler{
		cmd:  cmd,
		args: args,
	}
	h.init()
	h.handle()
}
