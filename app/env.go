package app

type Env struct {
	name string
}

func InitEnv() {

	envName := G.Cfg.GetString("env")

	env := new(Env)
	env.name = envName

	G.Env = env
}