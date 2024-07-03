import java.util.*;
class Solution {
    public String solution(int[] numbers) {
        String[] strNum = Arrays.stream(numbers)
                .mapToObj(String::valueOf)
                .toArray(String[]::new);
        
        Arrays.sort(strNum, (x, y) -> (y + x).compareTo(x + y));
        
        if (strNum[0].equals("0")) {
            return "0";
        }
        
        StringBuilder r = new StringBuilder();
        for (String s : strNum) {
            r.append(s);
        }
        return r.toString();
    }
}