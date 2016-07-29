package rie


const (
	Info = 0
	Warning = 1
	Error = 2
	Severe = 3
	Once = 128
)	


const (
	/* 
        Error Codes
         1 - 10         System and File Errors
        11 - 20         Program Limitations
        21 - 40         State Errors
        41 - 60         Parameter and Protocol Errors
        61 - 80         Execution Errors
	 */
	NoError = 0
	NoMem = 1
	System = 2
	NoFile = 3
	BadFile = 4
	Version = 5
	DiskFull = 6	
	Incapable = 11
	Unimplement = 12
	Limit = 13
	Bug = 14
	NotStarted = 23
	Nesting = 24
	NotOptions = 25
	NotAttribs = 26
	NotPrims = 27
	IllState = 28
	BadMotion = 29
	BadSolid = 30
	BadToken = 41
	Range = 42
	Consistency = 43
	BadHandle = 44
	NoShader = 45
	MissingData = 46
	Syntax = 47
	Skipping = 51
	Math = 61
)

