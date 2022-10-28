package sys

const (
	// Success 以下系统整体级别返回编号
	Success            = 10000
	UnknownError       = 10001
	ParameterJsonError = 10002
	MethodParamsError  = 10003

	// TokenExpires 以下登陆相关编号
	TokenExpires             = 10003
	LoginNameOrPasswordError = 10010
	DuplicationLoginName     = 10011
	LoginNameFormat          = 10012
	PasswordFormat           = 10013
	// DatabaseConnectionError 数据库连接错误代码
	DatabaseConnectionError = 1100
	// SqlNotAffectedRow sql影响的行数为0，预期至少为1.
	SqlNotAffectedRow = 11001

	// SqlNoSuchDataOnParameters 条件查询失败
	SqlNoSuchDataOnParameters = 11002
	// SqlError sql执行出错
	SqlError = 11003

	TokenMsg                     = "token! "
	JsonErrorMsg                 = "Please send the correct JSON string! "
	LoginNameFormatErrorMsg      = "Please enter the correct login name,It consists of a number of characters, words, or underscores! "
	PasswordErrorMsg             = "Please enter the correct password! "
	LoginNameOrPasswordErrorMsg  = "Please enter the correct login name or password! "
	SuccessMsg                   = "The service you requested has been executed successfully! "
	ServiceErrorMsg              = "Your request caused an internal error in the service. Please contact the administrator! "
	DuplicationLoginNameMsg      = "login name already exists! "
	TokenExpiresMsg              = "Your token expires or is not signed in! "
	SqlNotAffectedRowMsg         = "The service you called was executed successfully, but the number of rows affected is zero! "
	SqlNoSuchDataOnParametersMsg = "No data found based on the passed in parameters! "
	SqlErrorMsg                  = "SQL execution error! "
	RequestedFailedMsg           = "The service you requested failed! "
	UnknownErrorMsg              = "An unknown error occurred in the service you called! "
	MethodParamsErrorMsg         = "Wrong argument to function! "
	DatabaseConnectionErrorMsg   = "Database connection error! "
)
