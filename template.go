// Copyright 2014 izigo Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package izigo

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/izi-global/izigo/logs"
	"github.com/izi-global/izigo/utils"
)

var (
	izigoTplFuncMap           = make(template.FuncMap)
	iziViewPathTemplateLocked = false
	// iziViewPathTemplates caching map and supported template file extensions per view
	iziViewPathTemplates = make(map[string]map[string]*template.Template)
	templatesLock        sync.RWMutex
	// iziTemplateExt stores the template extension which will build
	iziTemplateExt = []string{"tpl", "html"}
	// iziTemplatePreprocessors stores associations of extension -> preprocessor handler
	iziTemplateEngines = map[string]templatePreProcessor{}
)

// ExecuteTemplate applies the template with name  to the specified data object,
// writing the output to wr.
// A template will be executed safely in parallel.
func ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return ExecuteViewPathTemplate(wr, name, BConfig.WebConfig.ViewsPath, data)
}

// ExecuteViewPathTemplate applies the template with name and from specific viewPath to the specified data object,
// writing the output to wr.
// A template will be executed safely in parallel.
func ExecuteViewPathTemplate(wr io.Writer, name string, viewPath string, data interface{}) error {
	if BConfig.RunMode == DEV {
		templatesLock.RLock()
		defer templatesLock.RUnlock()
	}
	if iziTemplates, ok := iziViewPathTemplates[viewPath]; ok {
		if t, ok := iziTemplates[name]; ok {
			var err error
			if t.Lookup(name) != nil {
				err = t.ExecuteTemplate(wr, name, data)
			} else {
				err = t.Execute(wr, data)
			}
			if err != nil {
				logs.Trace("template Execute err:", err)
			}
			return err
		}
		panic("can't find templatefile in the path:" + viewPath + "/" + name)
	}
	panic("Unknown view path:" + viewPath)
}

func init() {
	izigoTplFuncMap["dateformat"] = DateFormat
	izigoTplFuncMap["date"] = Date
	izigoTplFuncMap["compare"] = Compare
	izigoTplFuncMap["compare_not"] = CompareNot
	izigoTplFuncMap["not_nil"] = NotNil
	izigoTplFuncMap["not_null"] = NotNil
	izigoTplFuncMap["substr"] = Substr
	izigoTplFuncMap["html2str"] = HTML2str
	izigoTplFuncMap["str2html"] = Str2html
	izigoTplFuncMap["htmlquote"] = Htmlquote
	izigoTplFuncMap["htmlunquote"] = Htmlunquote
	izigoTplFuncMap["renderform"] = RenderForm
	izigoTplFuncMap["assets_js"] = AssetsJs
	izigoTplFuncMap["assets_css"] = AssetsCSS
	izigoTplFuncMap["config"] = GetConfig
	izigoTplFuncMap["map_get"] = MapGet

	// Comparisons
	izigoTplFuncMap["eq"] = eq // ==
	izigoTplFuncMap["ge"] = ge // >=
	izigoTplFuncMap["gt"] = gt // >
	izigoTplFuncMap["le"] = le // <=
	izigoTplFuncMap["lt"] = lt // <
	izigoTplFuncMap["ne"] = ne // !=

	izigoTplFuncMap["urlfor"] = URLFor // build a URL to match a Controller and it's method
}

// AddFuncMap let user to register a func in the template.
func AddFuncMap(key string, fn interface{}) error {
	izigoTplFuncMap[key] = fn
	return nil
}

type templatePreProcessor func(root, path string, funcs template.FuncMap) (*template.Template, error)

type templateFile struct {
	root  string
	files map[string][]string
}

// visit will make the paths into two part,the first is subDir (without tf.root),the second is full path(without tf.root).
// if tf.root="views" and
// paths is "views/errors/404.html",the subDir will be "errors",the file will be "errors/404.html"
// paths is "views/admin/errors/404.html",the subDir will be "admin/errors",the file will be "admin/errors/404.html"
func (tf *templateFile) visit(paths string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	}
	if f.IsDir() || (f.Mode()&os.ModeSymlink) > 0 {
		return nil
	}
	if !HasTemplateExt(paths) {
		return nil
	}

	replace := strings.NewReplacer("\\", "/")
	file := strings.TrimLeft(replace.Replace(paths[len(tf.root):]), "/")
	subDir := filepath.Dir(file)

	tf.files[subDir] = append(tf.files[subDir], file)
	return nil
}

// HasTemplateExt return this path contains supported template extension of izigo or not.
func HasTemplateExt(paths string) bool {
	for _, v := range iziTemplateExt {
		if strings.HasSuffix(paths, "."+v) {
			return true
		}
	}
	return false
}

// AddTemplateExt add new extension for template.
func AddTemplateExt(ext string) {
	for _, v := range iziTemplateExt {
		if v == ext {
			return
		}
	}
	iziTemplateExt = append(iziTemplateExt, ext)
}

// AddViewPath adds a new path to the supported view paths.
//Can later be used by setting a controller ViewPath to this folder
//will panic if called after izigo.Run()
func AddViewPath(viewPath string) error {
	if iziViewPathTemplateLocked {
		if _, exist := iziViewPathTemplates[viewPath]; exist {
			return nil //Ignore if viewpath already exists
		}
		panic("Can not add new view paths after izigo.Run()")
	}
	iziViewPathTemplates[viewPath] = make(map[string]*template.Template)
	return BuildTemplate(viewPath)
}

func lockViewPaths() {
	iziViewPathTemplateLocked = true
}

// BuildTemplate will build all template files in a directory.
// it makes izigo can render any template file in view directory.
func BuildTemplate(dir string, files ...string) error {
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return errors.New("dir open err")
	}
	iziTemplates, ok := iziViewPathTemplates[dir]
	if !ok {
		panic("Unknown view path: " + dir)
	}
	self := &templateFile{
		root:  dir,
		files: make(map[string][]string),
	}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		return self.visit(path, f, err)
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
		return err
	}
	buildAllFiles := len(files) == 0
	for _, v := range self.files {
		for _, file := range v {
			if buildAllFiles || utils.InSlice(file, files) {
				templatesLock.Lock()
				ext := filepath.Ext(file)
				var t *template.Template
				if len(ext) == 0 {
					t, err = getTemplate(self.root, file, v...)
				} else if fn, ok := iziTemplateEngines[ext[1:]]; ok {
					t, err = fn(self.root, file, izigoTplFuncMap)
				} else {
					t, err = getTemplate(self.root, file, v...)
				}
				if err != nil {
					logs.Error("parse template err:", file, err)
					templatesLock.Unlock()
					return err
				}
				iziTemplates[file] = t
				templatesLock.Unlock()
			}
		}
	}
	return nil
}

func getTplDeep(root, file, parent string, t *template.Template) (*template.Template, [][]string, error) {
	var fileAbsPath string
	var rParent string
	if filepath.HasPrefix(file, "../") {
		rParent = filepath.Join(filepath.Dir(parent), file)
		fileAbsPath = filepath.Join(root, filepath.Dir(parent), file)
	} else {
		rParent = file
		fileAbsPath = filepath.Join(root, file)
	}
	if e := utils.FileExists(fileAbsPath); !e {
		panic("can't find template file:" + file)
	}
	data, err := ioutil.ReadFile(fileAbsPath)
	if err != nil {
		return nil, [][]string{}, err
	}
	t, err = t.New(file).Parse(string(data))
	if err != nil {
		return nil, [][]string{}, err
	}
	reg := regexp.MustCompile(BConfig.WebConfig.TemplateLeft + "[ ]*template[ ]+\"([^\"]+)\"")
	allSub := reg.FindAllStringSubmatch(string(data), -1)
	for _, m := range allSub {
		if len(m) == 2 {
			tl := t.Lookup(m[1])
			if tl != nil {
				continue
			}
			if !HasTemplateExt(m[1]) {
				continue
			}
			_, _, err = getTplDeep(root, m[1], rParent, t)
			if err != nil {
				return nil, [][]string{}, err
			}
		}
	}
	return t, allSub, nil
}

func getTemplate(root, file string, others ...string) (t *template.Template, err error) {
	t = template.New(file).Delims(BConfig.WebConfig.TemplateLeft, BConfig.WebConfig.TemplateRight).Funcs(izigoTplFuncMap)
	var subMods [][]string
	t, subMods, err = getTplDeep(root, file, "", t)
	if err != nil {
		return nil, err
	}
	t, err = _getTemplate(t, root, subMods, others...)

	if err != nil {
		return nil, err
	}
	return
}

func _getTemplate(t0 *template.Template, root string, subMods [][]string, others ...string) (t *template.Template, err error) {
	t = t0
	for _, m := range subMods {
		if len(m) == 2 {
			tpl := t.Lookup(m[1])
			if tpl != nil {
				continue
			}
			//first check filename
			for _, otherFile := range others {
				if otherFile == m[1] {
					var subMods1 [][]string
					t, subMods1, err = getTplDeep(root, otherFile, "", t)
					if err != nil {
						logs.Trace("template parse file err:", err)
					} else if len(subMods1) > 0 {
						t, err = _getTemplate(t, root, subMods1, others...)
					}
					break
				}
			}
			//second check define
			for _, otherFile := range others {
				var data []byte
				fileAbsPath := filepath.Join(root, otherFile)
				data, err = ioutil.ReadFile(fileAbsPath)
				if err != nil {
					continue
				}
				reg := regexp.MustCompile(BConfig.WebConfig.TemplateLeft + "[ ]*define[ ]+\"([^\"]+)\"")
				allSub := reg.FindAllStringSubmatch(string(data), -1)
				for _, sub := range allSub {
					if len(sub) == 2 && sub[1] == m[1] {
						var subMods1 [][]string
						t, subMods1, err = getTplDeep(root, otherFile, "", t)
						if err != nil {
							logs.Trace("template parse file err:", err)
						} else if len(subMods1) > 0 {
							t, err = _getTemplate(t, root, subMods1, others...)
						}
						break
					}
				}
			}
		}

	}
	return
}

// SetViewsPath sets view directory path in izigo application.
func SetViewsPath(path string) *App {
	BConfig.WebConfig.ViewsPath = path
	return IZIApp
}

// SetStaticPath sets static directory path and proper url pattern in izigo application.
// if izigo.SetStaticPath("static","public"), visit /static/* to load static file in folder "public".
func SetStaticPath(url string, path string) *App {
	if !strings.HasPrefix(url, "/") {
		url = "/" + url
	}
	if url != "/" {
		url = strings.TrimRight(url, "/")
	}
	BConfig.WebConfig.StaticDir[url] = path
	return IZIApp
}

// DelStaticPath removes the static folder setting in this url pattern in izigo application.
func DelStaticPath(url string) *App {
	if !strings.HasPrefix(url, "/") {
		url = "/" + url
	}
	if url != "/" {
		url = strings.TrimRight(url, "/")
	}
	delete(BConfig.WebConfig.StaticDir, url)
	return IZIApp
}

// AddTemplateEngine add a new templatePreProcessor which support extension
func AddTemplateEngine(extension string, fn templatePreProcessor) *App {
	AddTemplateExt(extension)
	iziTemplateEngines[extension] = fn
	return IZIApp
}
