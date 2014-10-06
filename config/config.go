package config

import "os"

const CONFIG_FILE_PATH = "$HOME/.jotrc"

const (
    DEFAULT_JOT_DIR = "$HOME/.jot"
    DEFAULT_EDITOR = "vim"
)

type Config struct {
    JotDir string
}

func GetJotEditor () string {
    return os.ExpandEnv(DEFAULT_EDITOR)
}

func GetJotDir () string {
    return os.ExpandEnv(DEFAULT_JOT_DIR)
}
