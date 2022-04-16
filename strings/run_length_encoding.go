package strings

import "strconv"

type runLengthEncoding struct {
}

func NewRunLengthEncoding() runLengthEncoding {
	return runLengthEncoding{}
}

// RunLengthEncoding
/*<div class="html">
<p>
  Write a function that takes in a non-empty string and returns its run-length
  encoding.
</p>
<p>
  From Wikipedia, "run-length encoding is a form of lossless data compression in
  which runs of data are stored as a single data value and count, rather than as
  the original run." For this problem, a run of data is any sequence of
  consecutive, identical characters. So the run <span>"AAA"</span> would be
  run-length-encoded as <span>"3A"</span>.
</p>
<p>
  To make things more complicated, however, the input string can contain all
  sorts of special characters, including numbers. And since encoded data must be
  decodable, this means that we can't naively run-length-encode long runs. For
  example, the run <span>"AAAAAAAAAAAA"</span> (12 <span>A</span>s), can't
  naively be encoded as <span>"12A"</span>, since this string can be decoded as
  either <span>"AAAAAAAAAAAA"</span> or <span>"1AA"</span>. Thus, long runs (runs
  of 10 or more characters) should be encoded in a split fashion; the
  aforementioned run should be encoded as <span>"9A3A"</span>.
</p>
<h3>Sample Input</h3>
<pre><span class="CodeEditor-promptParameter">string</span> = "AAAAAAAAAAAAABBCCCCDD"
</pre>
<h3>Sample Output</h3>
<pre>"9A4A2B4C2D"
</pre>
</div>
 */
func (runLengthEncoding) RunLengthEncoding(str string) string {
	var encodedStringChar []byte
	currentLenght := 1

	for i := 1; i < len(str); i++ {
		if str[i] != str[i-1] || currentLenght == 9 {
			el := strconv.Itoa(currentLenght)
			encodedStringChar = append(encodedStringChar, el[0])
			encodedStringChar = append(encodedStringChar, str[i-1])
			currentLenght = 0
		}

		currentLenght++
	}

	encodedStringChar = append(encodedStringChar, strconv.Itoa(currentLenght)[0])
	encodedStringChar = append(encodedStringChar, str[len(str)-1])

	return string(encodedStringChar)
}
