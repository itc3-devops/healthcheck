// Copyright © 2017 Heptio
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

package healthcheck

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTrigger(t *testing.T) {
	trigger := NewTrigger()
	assert.NoError(t, trigger.Check()(), "expected initial state to be happy")

	err := errors.New("test error")
	trigger.Trip(err)
	assert.EqualError(t, trigger.Check()(), "test error", "expected trigger to be tripped")

	trigger.Reset()
	assert.NoError(t, trigger.Check()(), "expected trigger to be happy after reset")
}
