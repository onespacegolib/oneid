package oneid

func Init(host string, clientId, clientSecret, refId, redirectPath string) Context {
	return &context{
		oneAuth: oneAuth{
			clientId: clientId,
			clientSecret: clientSecret,
			refCode: refId,
			redirectPath: redirectPath,
		},
		host: host,
	}
}