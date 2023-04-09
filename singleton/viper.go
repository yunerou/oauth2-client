package singleton

import (
	"sync"

	"github.com/spf13/viper"
)

// WrapViper ..Wrapper of *viper.Viper.
type wrapViper struct {
	*viper.Viper
}

var (
	onceViper       sync.Once
	vipperSingleton *wrapViper
)

// GetViper ..
func GetViper() *wrapViper {
	onceViper.Do(
		func() {
			vp := viper.New()
			vp.AutomaticEnv()
			vipperSingleton = &wrapViper{
				vp,
			}
		})
	return vipperSingleton
}

// LoadConfigFile ..Viper auto realize the file type in yaml|toml|json, higher priority to the latter of `path`.
func (v *wrapViper) LoadConfigFile(path []string, filenameWithoutExtension string) {
	if len(path) == 0 {
		panic("path can not null or blank")
	}

	v.SetConfigName(filenameWithoutExtension)
	for _, p := range path {
		v.AddConfigPath(p)
		err := v.MergeInConfig()
		if err != nil {
			panic(err)
		}
	}
}
