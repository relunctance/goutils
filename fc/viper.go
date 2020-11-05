package fc

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func ViperFile(filepath string) error {
	if !IsExist(filepath) {
		return fmt.Errorf("file not exists: %s", filepath)
	}
	dirname, name, typename := SplitPath(filepath)
	viper.SetConfigName(name)
	viper.SetConfigType(typename)
	viper.AddConfigPath(dirname)
	return viper.ReadInConfig()
}

func SplitPath(path string) (dir, name, typename string) {
	arr := strings.Split(filepath.Base(path), ".")
	return filepath.Dir(path), arr[0], arr[1]
}
