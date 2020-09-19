package master

func Port() string {
	return ":5000"
}

func KeySha() string {
	return "Cl@v3_D3_tYp3_sh@_2cinK0"
}

func Host() string {
	return "redis-12203.c11.us-east-1-3.ec2.cloud.redislabs.com:12203"
}

func Password() string {
	return "1akfwnyutGZYHjpRDYh5Vp52KonI3efs"
}

func EmailReg() string {
	return "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
}

func PassReg() string {
	return "^[a-zA-Z0-9ÁÉÍÓÚÜÑáéíóúüñ.!#$:/%&*+/=?^_{|}~-]+"
}

