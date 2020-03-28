package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/hashicorp/hcl"
	"github.com/magiconair/properties"
	"github.com/mitchellh/mapstructure"
	"github.com/pelletier/go-toml"
	"github.com/recallsong/go-utils/reflectx"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
)

// TrimBOM .
func TrimBOM(f []byte) []byte {
	return bytes.TrimPrefix(f, []byte("\xef\xbb\xbf"))
}

var (
	envVarRe      = regexp.MustCompile(`\$(\w+)|\$\{(\w+)(:[^}]*)?\}`)
	envVarEscaper = strings.NewReplacer(
		`"`, `\"`,
		`\`, `\\`,
	)
)

// EscapeEnv .
func EscapeEnv(contents []byte) []byte {
	params := envVarRe.FindAllSubmatch(contents, -1)
	for _, param := range params {
		if len(param) != 4 {
			continue
		}
		var key, defval []byte
		if len(param[1]) > 0 {
			key = param[1]
		} else if len(param[2]) > 0 {
			key = param[2]
		} else {
			continue
		}
		if len(param[3]) > 0 {
			defval = param[3][1:]
		}
		val, ok := os.LookupEnv(strings.TrimPrefix(reflectx.BytesToString(key), "$"))
		if !ok {
			val = string(defval)
		}
		val = envVarEscaper.Replace(val)
		contents = bytes.Replace(contents, param[0], reflectx.StringToBytes(val), 1)
	}
	return contents
}

// ParseError denotes failing to parse configuration file.
type ParseError struct {
	err error
}

// Error returns the formatted configuration error.
func (pe ParseError) Error() string {
	return fmt.Sprintf("While parsing config: %s", pe.err.Error())
}

// UnmarshalToMap .
func UnmarshalToMap(in io.Reader, typ string, c map[string]interface{}) (err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(in)
	switch strings.ToLower(typ) {
	case "yaml", "yml":
		if err = yaml.Unmarshal(buf.Bytes(), &c); err != nil {
			return ParseError{err}
		}
	case "json":
		if err = json.Unmarshal(buf.Bytes(), &c); err != nil {
			return ParseError{err}
		}
	case "hcl":
		obj, err := hcl.Parse(reflectx.BytesToString(buf.Bytes()))
		if err != nil {
			return ParseError{err}
		}
		if err = hcl.DecodeObject(&c, obj); err != nil {
			return ParseError{err}
		}
	case "toml":
		tree, err := toml.LoadReader(buf)
		if err != nil {
			return ParseError{err}
		}
		tmap := tree.ToMap()
		for k, v := range tmap {
			c[k] = v
		}
	case "properties", "props", "prop":
		props := properties.NewProperties()
		var err error
		if err = props.Load(buf.Bytes(), properties.UTF8); err != nil {
			return ParseError{err}
		}
		for _, key := range props.Keys() {
			value, _ := props.Get(key)
			// recursively build nested maps
			path := strings.Split(key, ".")
			lastKey := strings.ToLower(path[len(path)-1])
			deepestMap := deepSearch(c, path[0:len(path)-1])
			// set innermost value
			deepestMap[lastKey] = value
		}
	case "ini":
		cfg := ini.Empty()
		err = cfg.Append(buf.Bytes())
		if err != nil {
			return ParseError{err}
		}
		sections := cfg.Sections()
		for i := 0; i < len(sections); i++ {
			section := sections[i]
			keys := section.Keys()
			for j := 0; j < len(keys); j++ {
				key := keys[j]
				value := cfg.Section(section.Name()).Key(key.Name()).String()
				c[section.Name()+"."+key.Name()] = value
			}
		}
	}
	insensitiviseMap(c)
	return nil
}

func insensitiviseMap(m map[string]interface{}) {
	for key, val := range m {
		switch val.(type) {
		case map[interface{}]interface{}:
			val, _ = toStringMap(val)
			insensitiviseMap(val.(map[string]interface{}))
		case map[string]interface{}:
			insensitiviseMap(val.(map[string]interface{}))
		}
		lower := strings.ToLower(key)
		if key != lower {
			// remove old key (not lower-cased)
			delete(m, key)
		}
		// update map
		m[lower] = val
	}
}

// toStringMap casts an interface to a map[string]interface{} type.
func toStringMap(i interface{}) (map[string]interface{}, error) {
	var m = map[string]interface{}{}
	switch v := i.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			m[fmt.Sprint(k)] = val
		}
		return m, nil
	case map[string]interface{}:
		return v, nil
	case string:
		err := json.Unmarshal(reflectx.StringToBytes(v), &m)
		return m, err
	default:
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]interface{}", i, i)
	}
}

// deepSearch scans deep maps, following the key indexes listed in the
// sequence "path".
// The last value is expected to be another map, and is returned.
//
// In case intermediate keys do not exist, or map to a non-map value,
// a new map is created and inserted, and the search continues from there:
// the initial map "m" may be modified!
func deepSearch(m map[string]interface{}, path []string) map[string]interface{} {
	for _, k := range path {
		m2, ok := m[k]
		if !ok {
			// intermediate key does not exist
			// => create it and continue from there
			m3 := make(map[string]interface{})
			m[k] = m3
			m = m3
			continue
		}
		m3, ok := m2.(map[string]interface{})
		if !ok {
			// intermediate key is a value
			// => replace with a new map
			m3 = make(map[string]interface{})
			m[k] = m3
		}
		// continue search from here
		m = m3
	}
	return m
}

// ConvertData .
func ConvertData(input, output interface{}, tag string) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           output,
		WeaklyTypedInput: true,
		TagName:          tag,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			mapstructure.StringToTimeHookFunc("2006-01-02 15:04:05"),
		),
	})
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}

// LoadFile .
func LoadFile(path string) ([]byte, error) {
	byts, err := ioutil.ReadFile(path)
	if err == nil {
		byts = TrimBOM(byts)
		byts = EscapeEnv(byts)
	}
	return byts, err
}

// LoadToMap .
func LoadToMap(path string, c map[string]interface{}) error {
	typ := filepath.Ext(path)
	if len(typ) <= 0 {
		return fmt.Errorf("%s unknown file extension", path)
	}
	byts, err := LoadFile(path)
	if err != nil {
		return err
	}
	return UnmarshalToMap(bytes.NewReader(byts), typ[1:], c)
}

// LoadEnvFile .
func LoadEnvFile() {
	regex := regexp.MustCompile(`\s+\#`)
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	path := filepath.Join(wd, ".env")
	byts, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		return
	}
	content := reflectx.BytesToString(byts)
	for _, line := range strings.Split(content, "\n") {
		if strings.HasPrefix(line, "#") {
			continue
		}
		loc := regex.FindIndex(reflectx.StringToBytes(line))
		if len(loc) > 0 {
			line = line[0:loc[0]]
		}
		idx := strings.Index(line, "=")
		if idx <= 0 {
			continue
		}
		key := strings.TrimSpace(line[:idx])
		if len(key) <= 0 {
			continue
		}
		val := strings.TrimSpace(line[idx+1:])
		os.Setenv(key, val)
	}
}
