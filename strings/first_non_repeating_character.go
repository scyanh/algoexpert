package strings

type firstNonRepeatingCharacter struct {
}

func NewFirstNonRepeatingCharacter() firstNonRepeatingCharacter {
	return firstNonRepeatingCharacter{}
}

// FirstNonRepeatingCharacter
/*
<div class="html">
<p>
  Write a function that takes in a string of lowercase English-alphabet letters
  and returns the index of the string's first non-repeating character.
</p>
<p>
  The first non-repeating character is the first character in a string that
  occurs only once.
</p>
<p>
  If the input string doesn't have any non-repeating characters, your function
  should return <span>-1</span>.
</p>
<h3>Sample Input</h3>
<pre><span class="CodeEditor-promptParameter">string</span> = "abcdcaf"
</pre>
<h3>Sample Output</h3>
<pre>1 <span class="CodeEditor-promptComment">// The first non-repeating character is "b" and is found at index 1.</span>
</pre>
</div>
*/
func (firstNonRepeatingCharacter) FirstNonRepeatingCharacter(str string) int {
	memo := make(map[rune]int)

	for _, el := range str {
		memo[el]++
	}

	for i, el := range str {
		if memo[el] == 1 {
			return i
		}
	}

	return -1
}
