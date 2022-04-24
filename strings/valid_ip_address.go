package strings

import (
	"strconv"
	"strings"
)

type validIPAddress struct {
}

func NewValidIPAddress() validIPAddress {
	return validIPAddress{}
}

// ValidIPAddresses
/*
<div class="html">
<p>
  You're given a string of length 12 or smaller, containing only digits. Write a
  function that returns all the possible IP addresses that can be created by
  inserting three <span>.</span>s in the string.
</p>
<p>
  An IP address is a sequence of four positive integers that are separated by
  <span>.</span>s, where each individual integer is within the range
  <span>0 - 255</span>, inclusive.
</p>
<p>
  An IP address isn't valid if any of the individual integers contains leading
  <span>0</span>s. For example, <span>"192.168.0.1"</span> is a valid IP
  address, but <span>"192.168.00.1"</span> and
  <span>"192.168.0.01"</span> aren't, because they contain <span>"00"</span> and
  <span>01</span>, respectively. Another example of a valid IP address is
  <span>"99.1.1.10"</span>; conversely, <span>"991.1.1.0"</span> isn't valid,
  because <span>"991"</span> is greater than 255.
</p>
<p>
  Your function should return the IP addresses in string format and in no
  particular order. If no valid IP addresses can be created from the string,
  your function should return an empty list.
</p>
<p>
  Note: check out our Systems Design Fundamentals on SystemsExpert to learn more
  about IP addresses!
</p>
<h3>Sample Input</h3>
<pre><span class="CodeEditor-promptParameter">string</span> = "1921680"
</pre>
<h3>Sample Output</h3>
<pre>[
  "1.9.216.80",
  "1.92.16.80",
  "1.92.168.0",
  "19.2.16.80",
  "19.2.168.0",
  "19.21.6.80",
  "19.21.68.0",
  "19.216.8.0",
  "192.1.6.80",
  "192.1.68.0",
  "192.16.8.0"
]
<span class="CodeEditor-promptComment">// The IP addresses could be ordered differently.</span>
</pre>
</div>
*/
func (v validIPAddress) ValidIPAddresses(str string) []string {
	var ipAddresses []string
	currentIPAddress := []string{"", "", "", ""}

	for i := 1; i < 4; i++ {
		if i > len(str) {
			continue
		}

		currentIPAddress = []string{"", "", "", ""}

		currentIPAddress[0] = str[:i]
		if !v.isValid(currentIPAddress[0]) {
			continue
		}

		for j := i + 1; j < i+4; j++ {
			if j > len(str) {
				continue
			}

			currentIPAddress[1] = str[i:j]
			if !v.isValid(currentIPAddress[1]) {
				continue
			}

			for k := j + 1; k < j+4; k++ {
				if k > len(str) {
					continue
				}

				currentIPAddress[2] = str[j:k]

				if !v.isValid(currentIPAddress[2]) {
					continue
				}

				currentIPAddress[3] = str[k:]
				if !v.isValid(currentIPAddress[3]) {
					continue
				}

				ipAddresses = append(ipAddresses, strings.Join(currentIPAddress, "."))
			}
		}

	}

	return ipAddresses
}

func (validIPAddress) min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

func (validIPAddress) isValid(str string) bool {
	number, err := strconv.Atoi(str)
	if err != nil {
		return false
	}

	if number < 0 || number > 255 {
		return false
	}

	if len(str) != len(strconv.Itoa(number)) {
		return false
	}

	return true
}
