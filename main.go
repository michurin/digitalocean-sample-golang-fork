package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
)

const startupMessage = `[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;54;48;5;39m [38;5;54;48;5;39m [38;5;54;48;5;39m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;39;48;5;39m [38;5;21;48;5;45m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;92;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;93;48;5;45m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;39;48;5;39m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;92;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;92;48;5;45m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;99;48;5;39m [38;5;39;48;5;39m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;31;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;204;48;5;17m [38;5;204;48;5;17m [38;5;92;48;5;45m [38;5;54;48;5;39m [38;5;92;48;5;45m [38;5;92;48;5;45m [38;5;92;48;5;45m [38;5;92;48;5;45m [38;5;92;48;5;45m [38;5;92;48;5;45m [38;5;212;48;5;24m [38;5;204;48;5;17m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;125;48;5;24m [38;5;39;48;5;39m [38;5;117;48;5;39m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;92;48;5;45m [38;5;55;48;5;39m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;92;48;5;45m [38;5;92;48;5;39m [38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;69;48;5;18m�[38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;99;48;5;45m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;39;48;5;39m [38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;25;48;5;33m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;38;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;231;48;5;231m�[38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;227;48;5;227m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;105;48;5;122m�[38;5;32;48;5;159m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;211;48;5;234m�[38;5;39;48;5;159m�[38;5;1;48;5;16m [38;5;57;48;5;51m�[38;5;57;48;5;51m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;27;48;5;159m [38;5;117;48;5;159m�[38;5;39;48;5;159m�[38;5;75;48;5;159m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;75;48;5;159m�[38;5;25;48;5;159m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;31;48;5;39m [38;5;31;48;5;39m [38;5;56;48;5;239m�[38;5;56;48;5;239m�[38;5;56;48;5;239m�[38;5;56;48;5;239m�[38;5;56;48;5;239m�[38;5;56;48;5;239m�[38;5;56;48;5;239m�[38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;32;48;5;39m [38;5;32;48;5;33m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;45;48;5;81m [38;5;231;48;5;231m�[38;5;226;48;5;226m [38;5;62;48;5;17m�[38;5;4;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;226;48;5;226m [38;5;117;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;171;48;5;39m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;69;48;5;122m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;32;48;5;159m�[38;5;195;48;5;195m�[38;5;33;48;5;159m�[38;5;75;48;5;159m�[38;5;25;48;5;195m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;211;48;5;234m [38;5;1;48;5;16m [38;5;99;48;5;73m�[38;5;99;48;5;159m�[38;5;161;48;5;26m�[38;5;21;48;5;87m�[38;5;159;48;5;159m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;136;48;5;253m�[38;5;136;48;5;253m�[38;5;136;48;5;253m�[38;5;178;48;5;253m�[38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;231;48;5;231m�[38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;25;48;5;33m [38;5;117;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;231;48;5;231m�[38;5;227;48;5;227m [38;5;226;48;5;226m [38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;93;48;5;239m�[38;5;230;48;5;230m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;75;48;5;159m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;117;48;5;159m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;75;48;5;159m�[38;5;21;48;5;159m�[38;5;39;48;5;159m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;135;48;5;45m�[38;5;57;48;5;87m�[38;5;21;48;5;87m�[38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;123;48;5;123m [38;5;195;48;5;195m [38;5;195;48;5;195m [38;5;195;48;5;195m [38;5;195;48;5;195m [38;5;195;48;5;195m [38;5;195;48;5;195m [38;5;195;48;5;195m [38;5;136;48;5;253m�[38;5;136;48;5;253m�[38;5;221;48;5;224m�[38;5;94;48;5;253m�[38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;117;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;117;48;5;39m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;57;48;5;87m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;55;48;5;45m�[38;5;57;48;5;87m�[38;5;57;48;5;51m�[38;5;57;48;5;87m�[38;5;99;48;5;87m�[38;5;141;48;5;87m�[38;5;57;48;5;51m�[38;5;166;48;5;249m�[38;5;166;48;5;249m�[38;5;166;48;5;249m�[38;5;43;48;5;145m�[38;5;119;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;62;48;5;17m�[38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;26;48;5;75m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;38;48;5;45m [38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;87;48;5;255m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;167;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;69;48;5;18m�[38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;219;48;5;219m [38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;109;48;5;231m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;61;48;5;24m�[38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;111;48;5;24m [38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;61;48;5;60m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;162;48;5;205m [38;5;162;48;5;205m [38;5;162;48;5;205m [38;5;89;48;5;212m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;27;48;5;123m�[38;5;1;48;5;16m [38;5;135;48;5;51m�[38;5;204;48;5;233m [38;5;1;48;5;16m [38;5;135;48;5;51m�[38;5;135;48;5;51m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;165;48;5;32m [38;5;135;48;5;39m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;162;48;5;205m [38;5;212;48;5;205m [38;5;162;48;5;205m [38;5;162;48;5;205m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;84;48;5;212m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;135;48;5;51m [38;5;45;48;5;231m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;99;48;5;51m�[38;5;1;48;5;16m [38;5;55;48;5;45m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;56;48;5;51m�[38;5;45;48;5;231m�[38;5;56;48;5;51m [38;5;1;48;5;16m [38;5;111;48;5;87m�[38;5;69;48;5;87m�[38;5;1;48;5;16m [38;5;87;48;5;87m [38;5;141;48;5;51m�[38;5;39;48;5;39m [38;5;117;48;5;39m [38;5;39;48;5;39m [38;5;24;48;5;39m [38;5;39;48;5;39m [38;5;117;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;92;48;5;39m�[38;5;165;48;5;33m [38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;199;48;5;212m [38;5;212;48;5;206m [38;5;225;48;5;225m [38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;56;48;5;51m�[38;5;141;48;5;87m�[38;5;123;48;5;231m�[38;5;31;48;5;195m [38;5;57;48;5;73m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;51;48;5;51m [38;5;56;48;5;51m�[38;5;81;48;5;195m�[38;5;55;48;5;45m [38;5;99;48;5;51m�[38;5;177;48;5;38m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;123;48;5;255m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;51;48;5;255m�[38;5;168;48;5;241m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[0m`

func logRequest(r *http.Request) {
	uri := r.RequestURI
	method := r.Method
	fmt.Println("Got request!", method, uri)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		fmt.Fprintf(w, "Hello! you've requested %s\n", r.URL.Path)
	})

	http.HandleFunc("/cached", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		query := r.URL.Query()
		maxAgeParam := query.Get("max-age")
		if len(maxAgeParam) > 0 {
			maxAge, _ := strconv.Atoi(maxAgeParam)
			w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", maxAge))
		}
		responseHeaderParams, ok := query["headers"]
		if ok {
			for _, header := range responseHeaderParams {
				h, v, ok := strings.Cut(header, ":")
				if !ok {
					continue
				}
				w.Header().Set(h, strings.TrimSpace(v))
			}
		}
		statusCodeParam := query.Get("status")
		if len(statusCodeParam) > 0 {
			statusCode, _ := strconv.Atoi(statusCodeParam)
			if statusCode >= 200 && statusCode < 600 {
				w.WriteHeader(statusCode)
			}
		}
		requestID := uuid.Must(uuid.NewV4())
		fmt.Fprint(w, requestID.String())
	})

	http.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		key := r.URL.Query().Get("key")
		if len(key) > 0 {
			fmt.Fprint(w, r.Header.Get(key))
			return
		}
		headers := []string{}
		headers = append(headers, fmt.Sprintf("host=%s", r.Host))
		for key, values := range r.Header {
			headers = append(headers, fmt.Sprintf("%s=%s", key, strings.Join(values, ",")))
		}
		fmt.Fprint(w, strings.Join(headers, "\n"))
	})

	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		keys, ok := r.URL.Query()["key"]
		if ok && len(keys) > 0 {
			fmt.Fprint(w, os.Getenv(keys[0]))
			return
		}
		envs := []string{}
		envs = append(envs, os.Environ()...)
		fmt.Fprint(w, strings.Join(envs, "\n"))
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		codeParams, ok := r.URL.Query()["code"]
		if ok && len(codeParams) > 0 {
			statusCode, _ := strconv.Atoi(codeParams[0])
			if statusCode >= 200 && statusCode < 600 {
				w.WriteHeader(statusCode)
			}
		}
		requestID := uuid.Must(uuid.NewV4())
		fmt.Fprint(w, requestID.String())
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	encodedRouteString := os.Getenv("ROUTES")
	if encodedRouteString != "" {
		for _, encodedRoute := range strings.Split(encodedRouteString, ",") {
			path, body, ok := strings.Cut(encodedRoute, "=")
			if !ok {
				fmt.Printf("Skip routing %q: wrong format", encodedRoute)
				continue
			}
			http.HandleFunc("/"+path, func(w http.ResponseWriter, _ *http.Request) {
				fmt.Fprint(w, body)
			})
		}
	}

	bindAddr := fmt.Sprintf(":%s", port)
	lines := strings.Split(startupMessage, "\n")
	fmt.Println()
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Println()
	fmt.Printf("==> Server listening at %s 🚀\n", bindAddr)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}
}
