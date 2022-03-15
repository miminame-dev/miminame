package props

import (
	"github.com/miminame-dev/miminame/backend/pkg/config"
	"github.com/miminame-dev/miminame/backend/service"
)

type Props struct {
	Config              *config.Config
	ProcessVideoService service.IProcessVideoService
}
