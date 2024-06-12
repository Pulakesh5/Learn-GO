package isogram
func IsIsogram(word string) bool {
	Count := make([]int, 26)
    for _,ch := range word {
        if (ch>='a' && ch<='z') {
			Count[int(ch-'a')]++;   
        } else if (ch>='A' && ch<='Z') {
			Count[int(ch-'A')]++;
        }    	
    }
    for i:=0; i<26; i++ {
        if(Count[i]>1) {
            return false
        }
    }
    return true
}
