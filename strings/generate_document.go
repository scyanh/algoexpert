package strings

type generateDocument struct {
}

func NewGenerateDocument() generateDocument {
	return generateDocument{}
}

// GenerateDocument
/*
<div class="html">
<p>
  You're given a string of available characters and a string representing a
  document that you need to generate. Write a function that determines if you
  can generate the document using the available characters. If you can generate
  the document, your function should return <span>true</span>; otherwise, it
  should return <span>false</span>.
</p>
<p>
  You're only able to generate the document if the frequency of unique
  characters in the characters string is greater than or equal to the frequency
  of unique characters in the document string. For example, if you're given
  <span>characters = "abcabc"</span> and <span>document = "aabbccc"</span> you
  <b>cannot</b> generate the document because you're missing one <span>c</span>.
</p>
<p>
  The document that you need to create may contain any characters, including
  special characters, capital letters, numbers, and spaces.
</p>
<p>Note: you can always generate the empty string (<span>""</span>).</p>
<h3>Sample Input</h3>
<pre><span class="CodeEditor-promptParameter">characters</span> = "Bste!hetsi ogEAxpelrt x "
<span class="CodeEditor-promptParameter">document</span> = "AlgoExpert is the Best!"
</pre>
<h3>Sample Output</h3>
<pre>true
</pre>
</div>
 */
func (generateDocument) GenerateDocument(characters string, document string) bool {
	memo := make(map[rune]int)

	for _, el := range characters {
		memo[el]++
	}

	for _, el := range document {
		if memo[el] > 0 {
			memo[el]--
		} else {
			return false
		}
	}

	return true
}
