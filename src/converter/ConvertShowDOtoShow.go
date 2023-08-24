package converter

import "homepage-server/src/aws"

type Show struct {
	date    string
	venue   string
	address string
}

func ConvertShowDOtoShow(s *aws.ShowDO) *Show {
	return nil
}
