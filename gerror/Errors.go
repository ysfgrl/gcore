package gerror

var (
	UserNF = &Error{
		Code:   "api.user.nf",
		Detail: "User Not Found",
		Level:  "error",
	}
	SessionTO = &Error{
		Code:   "api.session_timeout",
		Detail: "Session Timeout",
		Level:  "error",
	}
	PassRequire = &Error{
		Code:   "api.user.pass_require",
		Detail: "Password Require",
		Level:  "error",
	}

	SaveFailed = &Error{
		Code:   "api.save_failed",
		Detail: "Save Failed",
		Level:  "error",
	}
	ListNF = &Error{
		Code:   "api.list_nf",
		Detail: "List Not Found",
		Level:  "error",
	}
	NF = &Error{
		Code:   "api.nf",
		Detail: "Not Found",
		Level:  "error",
	}
	IdNF = &Error{
		Code:   "api.id_nf",
		Detail: "Id Not Found",
		Level:  "error",
	}
	TypeNF = &Error{
		Code:   "api.type_nf",
		Detail: "Type Not Found",
		Level:  "error",
	}
	UrlNF = &Error{
		Code:   "api.url_nf",
		Detail: "Url Not Found",
		Level:  "error",
	}
	IdRequired = &Error{
		Code:   "api.id_required",
		Detail: "Id Required",
		Level:  "error",
	}
	PathParamRequired = &Error{
		Code:   "api.path_param_required",
		Detail: "Path Param Required",
		Level:  "error",
	}
	NotImplementedError = &Error{
		Code:   "api.not_implemented",
		Detail: "Not Implemented",
		Level:  "error",
	}
	PermissionDenied = &Error{
		Code:   "api.permission_denied",
		Detail: "Permission Denied",
		Level:  "error",
	}
	ForbiddenError = &Error{
		Code:   "api.forbidden",
		Detail: "Permission Denied",
		Level:  "error",
	}
	DecodeError = &Error{
		Code:   "api.decode_error",
		Detail: "Decode Failed",
		Level:  "error",
	}
	TokenKeyError = &Error{
		Code:   "api.token.key_error",
		Detail: "Key Not Found",
		Level:  "error",
	}
	TokenAlgError = &Error{
		Code:   "api.token.alg_error",
		Detail: "Algorithm Not Found",
		Level:  "error",
	}
	TokenUserError = &Error{
		Code:   "api.token.user_error",
		Detail: "Token User Not Found",
		Level:  "error",
	}
	TokenTokenError = &Error{
		Code:   "api.token.token_error",
		Detail: "Token Not Found",
		Level:  "error",
	}
	StorageNotOnlineError = &Error{
		Code:   "api.storage.not_online",
		Detail: "Storage Not online",
		Level:  "error",
	}
	StorageBucketNotExistError = &Error{
		Code:   "api.storage.bucket_not_exist",
		Detail: "Bucket Not Exist",
		Level:  "error",
	}
	StorageBucketCreateError = &Error{
		Code:   "api.storage.bucket_create_error",
		Detail: "Bucket Create Failed",
		Level:  "error",
	}
	MongoNotFoundDocument = &Error{
		Code:   "api.not_found_document",
		Detail: "Not Found",
		Level:  "error",
	}

	Errors = []*Error{
		UserNF,
		SessionTO,
		PassRequire,
		SaveFailed,
		ListNF,
		NF,
		IdNF,
		TypeNF,
		UrlNF,
		IdRequired,
		PathParamRequired,
		NotImplementedError,
		PermissionDenied,
		ForbiddenError,
		DecodeError,
		TokenKeyError,
		TokenAlgError,
		TokenUserError,
		TokenTokenError,
		StorageNotOnlineError,
		StorageBucketNotExistError,
		StorageBucketCreateError,
		MongoNotFoundDocument,
	}
)
