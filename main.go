package main

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type structTags struct {
	tags   []reflect.StructTag // 各フィールドのタグ
	length map[string]int      // キーのタグ名の最も長いタグの長さ
	order  []string            // タグ名の順序
}

var reTagName = regexp.MustCompile(`(\w+):`)

func newStructTags(tags []string) *structTags {
	st := structTags{
		tags:   make([]reflect.StructTag, len(tags)),
		length: map[string]int{},
	}
	for i, tag := range tags {
		if tag == "" {
			continue
		}

		rst := reflect.StructTag(tag)
		for _, match := range reTagName.FindAllStringSubmatch(tag, -1) {
			tagname := match[1]
			length, ok := st.length[tagname]
			if !ok {
				// 初めて出現したタグ名を登録
				st.order = append(st.order, tagname)
			}
			if l := len(tagstr(tagname, rst.Get(tagname))); l > length {
				// 最も長いタグの長さを更新
				st.length[tagname] = l
			}
		}
		st.tags[i] = rst
	}
	return &st
}

func (st *structTags) aligned(index int) string {
	b := new(bytes.Buffer)
	for _, tagname := range st.order {
		var t string
		if value, ok := st.tags[index].Lookup(tagname); ok {
			t = tagstr(tagname, value)
		}
		b.WriteString(t)

		// 一番後ろに少なくとも1つはスペースを入れる
		b.WriteString(strings.Repeat(" ", st.length[tagname]-len(t)+1))
	}
	return strings.TrimRight(b.String(), " ")
}

func Align(tags []string) []string {
	result := make([]string, len(tags))

	st := newStructTags(tags)
	for i := range tags {
		if tags[i] != "" {
			result[i] = st.aligned(i)
		}
	}

	return result
}

func tagstr(tagname, value string) string {
	return tagname + `:"` + value + `"`
}

func unquote(tag string) string {
	s, err := strconv.Unquote(string(tag))
	if err != nil {
		panic(err) // 不正なタグはParseFileでエラーになる
	}
	return s
}

func quote(tag string) string {
	return "`" + tag + "`"
}

func main() {
	// ソースコード情報の取得（From ファイル）
	filename := "./src/domain/user.go"

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	// ast.Print(fset, file)

	// 抽象構文木の探索
	ast.Inspect(file, func(node ast.Node) bool {
		s, ok := node.(*ast.StructType)
		if !ok {
			return true
		}
		// ast.Print(fset, s)

		tags := make([]string, len(s.Fields.List))
		for i, f := range s.Fields.List {
			if f.Tag == nil {
				continue
			}
			tags[i] = unquote(f.Tag.Value) // ここでバッククオートを除いておかないと後続の処理がうまくいかないため注意
		}

		for i, tag := range Align(tags) {
			if s.Fields.List[i].Tag == nil {
				continue
			}
			s.Fields.List[i].Tag.Value = quote(tag)
		}

		return true
	})

	format.Node(os.Stdout, fset, file)

	// fmt.Println(tags)
}
