class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int i = 0, j = 0, len = s.length();
        int maxlen = 0;
        int book[256] = {0};
        while (j < len) {
            if (book[s[j]] == 1) {
                maxlen = max(maxlen, j - i);
                while(s[i] != s[j]) {
                    book[s[i]] = 0;
                    ++i;
                }
                ++i;
            } else {
                book[s[j]] = 1;
            }
            ++j;
        }
        maxlen = max(maxlen, j - i);
        return maxlen;
    }
};
