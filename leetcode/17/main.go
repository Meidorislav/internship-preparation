func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }

    btns := map[rune]string{
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

    result := []string{""}
    for _, digit := range digits {
        var nextResult []string
        for _, existingCombo := range result {
            for _, letter := range btns[digit] {
                nextResult = append(nextResult, existingCombo + string(letter))
            }
        }
        result = nextResult
    }
    return result
}
