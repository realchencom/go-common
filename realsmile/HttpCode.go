package realsmile

type http struct {
	UnauthorizedStatus  int32
	ForbiddenStatus     int32
	SuccessStatus       int32
	InternalServerError int32
}

var (
	Http = http{
		UnauthorizedStatus:  401,
		ForbiddenStatus:     403,
		SuccessStatus:       200,
		InternalServerError: 500,
	}
)
