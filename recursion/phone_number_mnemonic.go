package recursion

type phoneNumberMnemonics struct {
}

func NewPhoneNumberMnemonics() phoneNumberMnemonics {
	return phoneNumberMnemonics{}
}

func (p phoneNumberMnemonics) PhoneNumberMnemonics(phoneNumber string) []string {
	currentMnemonic := make([]byte, len(phoneNumber))

	for i := range phoneNumber {
		currentMnemonic[i] = '0'
	}

	mnemonicsFound := make([]string, 0)

	p.phoneNumberMnemoicHelper(0, phoneNumber, currentMnemonic, &mnemonicsFound)

	return mnemonicsFound
}

func (p phoneNumberMnemonics) phoneNumberMnemoicHelper(idx int, phoneNumber string, currentMnemonic []byte, mnemonicsFound *[]string) {
	if idx == len(phoneNumber) {
		mnemonic := string(currentMnemonic)
		*mnemonicsFound = append(*mnemonicsFound, mnemonic)
		return
	}

	digit := phoneNumber[idx]
	letters := DigitLetters[digit]

	for _, el := range letters {
		currentMnemonic[idx] = el
		p.phoneNumberMnemoicHelper(idx+1, phoneNumber, currentMnemonic, mnemonicsFound)
	}

}

var DigitLetters = map[byte][]byte{
	'0': {'0'},
	'1': {'1'},
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}
