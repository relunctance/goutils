package fc

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestViperFile(t *testing.T) {
	json := `{ "a":"1234", "b":2344234, "c":true }`
	configFilepath := os.TempDir() + "/xxxxxxxa.json"
	err := ioutil.WriteFile(configFilepath, []byte(json), 0755)
	defer os.Remove(configFilepath)
	if err != nil {
		t.Fatalf("write tmp file:[%s] faild", configFilepath)
	}
	err = ViperFile(configFilepath)
	if err != nil {
		t.Fatalf("should be nil")
	}

	if viper.GetString("a") != "1234" {
		t.Fatalf("should be == '1234'")
	}
	if viper.GetInt("b") != 2344234 {
		t.Fatalf("should be == 2344234")
	}
	if viper.GetBool("c") != true {
		t.Fatalf("should be true")
	}
	unExistsFilePath := configFilepath + ".ccc.json"
	err = ViperFile(unExistsFilePath)
	if err == nil {
		t.Fatalf("should be nil")
	}

}

func TestSplitPath(t *testing.T) {
	dir, name, tname := SplitPath("./a.json")

	if dir != "." {
		t.Fatalf("should be == .")
	}
	if name != "a" {
		t.Fatalf("should be == a")
	}

	if tname != "json" {
		t.Fatalf("should be == json")
	}

	dir, name, tname = SplitPath(".")

	if dir != "." {
		t.Fatalf("should be == .")
	}
	if name != "" {
		t.Fatalf("should be == a")
	}
	if tname != "" {
		t.Fatalf("should be == json")
	}
}
