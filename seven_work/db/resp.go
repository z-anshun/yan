package db

type Resp struct {
	Message string`json:"message"`
	Code int`json:"code"`
}

func ErrorNameOrPasswd(msg string)*Resp{
	return &Resp{Message:msg,Code:400}
}

func ErrorSql(msg string)*Resp{
	return &Resp{Message:msg,Code:400}
}

func ErrorVerify(msg string)*Resp{
	return &Resp{Message:msg,Code:401}
}

func OK(msg string)*Resp{
	return &Resp{Message:msg,Code:204}
}