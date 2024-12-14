// Copyright © 2023 sealos.
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

package crypto

import "testing"

func TestRechargeBalance(t *testing.T) {
	data, _ := DecryptInt64("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJTZWFsb3MiLCJpYXQiOjE3MzMzMTQ4NzgsImV4cCI6MTc0MTA5MDg3OCwidHlwZSI6IkNsdXN0ZXIiLCJjbHVzdGVySUQiOiIyMjk1MDIwOSIsImRhdGEiOnsidG90YWxDUFUiOjgsInRvdGFsTWVtb3J5IjoxNiwibm9kZUNvdW50Ijo5OX19.kqBPSD8IcrbrSAYD3may2rM1eF-lWmy3dEUNrTjjMzqqyHTtTUAVZOPw7hU2rmTSShRi4P_iaAjvcnbQCw1w7_2J3Q2dtLdi7XUwTTtimc4fUYDnxiDhD3OUS2J1WydwnjPsPWAIIryqv8pGjdqqBBBL7-plDxcW9knzxT2zLDnBoH7YwiYoHiouZ5EIYWxXR3OlV2Fo8m4nONCSHaC75-Cr5Yk5M8bw4EdKagBcwmp0UWebPaFINVKHkK8ejTAWtm_KS18ElwTo0-SLtZeWI9gOZ1FhDPvmjbjpH52vxvJUvNoJuBjYglzsz-UZMARZkFlSaum8cx1bNbIyZCTrPQ")
	t.Log(data)
}
