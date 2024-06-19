package pkgConfig

type AuthConfig struct {
	ActiveDirectory AuthConfigActiveDirectory `yaml:"active_directory"`
}

type AuthConfigActiveDirectory struct {
	Network             string                             `yaml:"network"`
	Address             []AuthConfigActiveDirectoryAddress `yaml:"address"`
	Username            string                             `yaml:"username"`
	Password            string                             `yaml:"password"`
	BaseDistinguishName string                             `yaml:"base_distinguish_name"`
}

type AuthConfigActiveDirectoryAddress struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
