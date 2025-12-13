package util

// WildcardMatch reports whether the string s matches the pattern p.
// Supported wildcards:
// - '*' matches any sequence of characters (including empty)
// - '?' matches any single character
// No character escaping is supported.
func WildcardMatch(p, s string) bool {
    // Simple DP implementation
    m, n := len(p), len(s)
    dp := make([][]bool, m+1)
    for i := range dp {
        dp[i] = make([]bool, n+1)
    }
    dp[0][0] = true
    for i := 1; i <= m; i++ {
        if p[i-1] == '*' {
            dp[i][0] = dp[i-1][0]
        }
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            switch p[i-1] {
            case '*':
                dp[i][j] = dp[i-1][j] || dp[i][j-1]
            case '?':
                dp[i][j] = dp[i-1][j-1]
            default:
                dp[i][j] = dp[i-1][j-1] && p[i-1] == s[j-1]
            }
        }
    }
    return dp[m][n]
}
