package main

import (
	"flag"

	"log"
	"os"
	"path/filepath"
	"runtime/pprof"

	"github.com/spf13/viper"
	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/internal/index"
	_ "opencoredata.org/ocdGarden/CSDCO/VaultWalker/internal/vault"
)

var viperVal string

// example source s3://noaa-nwm-retro-v2.0-pds/full_physics/2017/201708010001.CHRTOUT_DOMAIN1.comp
func init() {
	log.SetFlags(log.Lshortfile)
	// log.SetOutput(ioutil.Discard) // turn off all logging
	flag.StringVar(&viperVal, "cfg", "config.json", "Configuration file")

}

// Optional profile section
func pf() {
	f, err := os.Create("cpuprofile.txt")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()
}

func main() {
	flag.Parse() // parse any command line flags...

	// fmt.Println("test")
	// viper
	var v1 *viper.Viper
	var err error

	// Load the config file
	if isFlagPassed("cfg") {
		v1, err = readConfig(viperVal, map[string]interface{}{})
		if err != nil {
			log.Print("error")
			// panic(log.Errorf("error when reading config: %v", err))
		}
	}

	vault := v1.GetStringMapString("vault")

	log.Printf("Indexing: %s\n", vault["dir"])

	//mc := minio.MinioConnection(minioVal, portVal, accessVal, secretVal)

	// read directory to []string
	var files []string
	var va []vault.VaultItem

	// Get directory and walk it and put a paths in files string array

	err = filepath.Walk(vault["dir"], visit(&files))
	if err != nil {
		panic(err)
	}

	// get the various elements of the found files, walk the files string array
	for _, file := range files {
		v, err := index.PathInspection(vault["dir"], file)
		if err == nil {
			va = append(va, v)
			log.Println(v)
		}
	}

	log.Println(va)

}

func readConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.SetConfigType("json")
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
		}
		*files = append(*files, path)
		return nil
	}
}
