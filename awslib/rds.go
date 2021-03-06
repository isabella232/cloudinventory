/*
Copyright 2019 Adobe. All rights reserved.
This file is licensed to you under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License. You may obtain a copy
of the License at http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR REPRESENTATIONS
OF ANY KIND, either express or implied. See the License for the specific language
governing permissions and limitations under the License.
*/

package awslib

import (
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/jpillora/backoff"
)

//GetAllDBInstances resturns a complete list of DBInstances for a given session
func GetAllDBInstances(sess *session.Session) ([]*rds.DBInstance, error) {
	rdsc := rds.New(sess)
	allInstancesDone := false
	var allInstances []*rds.DBInstance
	input := rds.DescribeDBInstancesInput{}
	b := &backoff.Backoff{
		//These are the defaults
		Min:    10 * time.Millisecond,
		Max:    30 * time.Second,
		Factor: 2,
		Jitter: false,
	}
	for !allInstancesDone {
		// Describe instances with no filters
		result, err := rdsc.DescribeDBInstances(&input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				// Retry with backoff incase Rate has been exceeded
				if aerr.Code() == "RateExceeded" {
					time.Sleep(b.Duration())
					continue
				}
				return allInstances, aerr
			}
			return allInstances, err
		}
		if err != nil {
			return nil, errors.New("Error Describing Instances")
		}
		allInstances = append(allInstances, result.DBInstances...)
		if result.Marker == nil {
			allInstancesDone = true
			continue
		}
		input.SetMarker(*result.Marker)
	}
	return allInstances, nil
}
