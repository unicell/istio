// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	"testing"

	"github.com/onsi/gomega"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestInstance_IsEmpty_False(t *testing.T) {
	g := gomega.NewWithT(t)

	e := Instance{
		Message: &emptypb.Empty{},
	}

	g.Expect(e.IsEmpty()).To(gomega.BeFalse())
}

func TestInstance_IsEmpty_True(t *testing.T) {
	g := gomega.NewWithT(t)
	e := Instance{}

	g.Expect(e.IsEmpty()).To(gomega.BeTrue())
}

func TestInstance_Clone_Empty(t *testing.T) {
	g := gomega.NewWithT(t)
	e := &Instance{}

	c := e.Clone()
	g.Expect(c).To(gomega.Equal(e))
}

func TestInstance_Clone_NonEmpty(t *testing.T) {
	g := gomega.NewWithT(t)

	e := &Instance{
		Message: &emptypb.Empty{},
	}

	c := e.Clone()
	g.Expect(c).To(gomega.Equal(e))
}
