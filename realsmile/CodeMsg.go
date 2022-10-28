package realsmile

type codeMsg struct {
	// Success 以下系统整体级别返回编号
	Success            int32
	UnknownError       int32
	ParameterJsonError int32
	MethodParamsError  int32

	// TokenExpires 以下登陆相关编号
	TokenExpires             int32
	LoginNameOrPasswordError int32
	DuplicationLoginName     int32
	LoginNameFormat          int32
	PasswordFormat           int32
	// DatabaseConnectionError 数据库连接错误代码
	DatabaseConnectionError int32
	// SqlNotAffectedRow sql影响的行数为0，预期至少为1.
	SqlNotAffectedRow int32

	// SqlNoSuchDataOnParameters 条件查询失败
	SqlNoSuchDataOnParameters int32
	// SqlError sql执行出错
	SqlError int32

	TokenMsg                     *string
	JsonErrorMsg                 *string
	LoginNameFormatErrorMsg      *string
	PasswordErrorMsg             *string
	LoginNameOrPasswordErrorMsg  *string
	SuccessMsg                   *string
	ServiceErrorMsg              *string
	DuplicationLoginNameMsg      *string
	TokenExpiresMsg              *string
	SqlNotAffectedRowMsg         *string
	SqlNoSuchDataOnParametersMsg *string
	SqlErrorMsg                  *string
	RequestedFailedMsg           *string
	UnknownErrorMsg              *string
	MethodParamsErrorMsg         *string
	DatabaseConnectionErrorMsg   *string
}

func init() {
	TokenMsg := "token! "
	JsonErrorMsg := "Please send the correct JSON string! "
	LoginNameFormatErrorMsg := "Please enter the correct login name,It consists of a number of characters, words, or underscores! "
	PasswordErrorMsg := "Please enter the correct password! "
	LoginNameOrPasswordErrorMsg := "Please enter the correct login name or password! "
	SuccessMsg := "The service you requested has been executed successfully! "
	ServiceErrorMsg := "Your request caused an internal error in the service. Please contact the administrator! "
	DuplicationLoginNameMsg := "login name already exists! "
	TokenExpiresMsg := "Your token expires or is not signed in! "
	SqlNotAffectedRowMsg := "The service you called was executed successfully, but the number of rows affected is zero! "
	SqlNoSuchDataOnParametersMsg := "No data found based on the passed in parameters! "
	SqlErrorMsg := "SQL execution error! "
	RequestedFailedMsg := "The service you requested failed! "
	UnknownErrorMsg := "An unknown error occurred in the service you called! "
	MethodParamsErrorMsg := "Wrong argument to function! "
	DatabaseConnectionErrorMsg := "Database connection error! "
	CodeMsg = codeMsg{
		// Success 以下系统整体级别返回编号
		Success:            10000,
		UnknownError:       10001,
		ParameterJsonError: 10002,
		MethodParamsError:  10003,

		// TokenExpires 以下登陆相关编号
		TokenExpires:             10003,
		LoginNameOrPasswordError: 10010,
		DuplicationLoginName:     10011,
		LoginNameFormat:          10012,
		PasswordFormat:           10013,
		// DatabaseConnectionError 数据库连接错误代码
		DatabaseConnectionError: 1100,
		// SqlNotAffectedRow sql影响的行数为0，预期至少为1.
		SqlNotAffectedRow: 11001,

		// SqlNoSuchDataOnParameters 条件查询失败
		SqlNoSuchDataOnParameters: 11002,
		// SqlError sql执行出错
		SqlError: 11003,

		TokenMsg:                     &TokenMsg,
		JsonErrorMsg:                 &JsonErrorMsg,
		LoginNameFormatErrorMsg:      &LoginNameFormatErrorMsg,
		PasswordErrorMsg:             &PasswordErrorMsg,
		LoginNameOrPasswordErrorMsg:  &LoginNameOrPasswordErrorMsg,
		SuccessMsg:                   &SuccessMsg,
		ServiceErrorMsg:              &ServiceErrorMsg,
		DuplicationLoginNameMsg:      &DuplicationLoginNameMsg,
		TokenExpiresMsg:              &TokenExpiresMsg,
		SqlNotAffectedRowMsg:         &SqlNotAffectedRowMsg,
		SqlNoSuchDataOnParametersMsg: &SqlNoSuchDataOnParametersMsg,
		SqlErrorMsg:                  &SqlErrorMsg,
		RequestedFailedMsg:           &RequestedFailedMsg,
		UnknownErrorMsg:              &UnknownErrorMsg,
		MethodParamsErrorMsg:         &MethodParamsErrorMsg,
		DatabaseConnectionErrorMsg:   &DatabaseConnectionErrorMsg,
	}
}

var (
	CodeMsg codeMsg
)
