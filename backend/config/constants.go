package config

import "os"

const UploadDir = "uploads"
const MaxFileSize = 8 << 20
const DefaultDirPermissions os.FileMode = 0755
