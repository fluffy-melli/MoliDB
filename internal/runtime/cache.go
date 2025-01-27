package runtime

import (
	"time"

	"github.com/fluffy-melli/MoliDB/pkg/cache"
	"github.com/fluffy-melli/MoliDB/pkg/config"
)

var DB = cache.LastLoad()
var Configs = config.ReadConfigFile()
var StartTime = time.Now().UnixMilli()
