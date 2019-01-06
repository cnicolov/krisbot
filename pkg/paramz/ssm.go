package paramz

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

type SSMProvider struct {
	client ssmiface.SSMAPI
	config *Config
}

func NewSSMClient() ssmiface.SSMAPI {
	sess := session.Must(session.NewSession())
	return ssm.New(sess)
}

func (s *SSMProvider) MustGetString(path string, withDecryption bool) string {
	key := s.config.Prefix + "/" + path
	inp := &ssm.GetParameterInput{
		Name:           aws.String(key),
		WithDecryption: aws.Bool(withDecryption),
	}

	out, err := s.client.GetParameter(inp)

	if err != nil {
		awsErr, ok := err.(awserr.Error)

		if !ok {
			panic(err)
		}
		switch awsErr.Code() {
		case ssm.ErrCodeParameterNotFound:
			log.Printf("Parameter %s not found", key)
			panic(err)
		default:
			panic(err)
		}
	}
	return *out.Parameter.Value
}
