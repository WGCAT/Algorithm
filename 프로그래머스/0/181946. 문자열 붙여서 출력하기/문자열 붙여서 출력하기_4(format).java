import java.util.Scanner;

// String.format()
public class Solution {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String a = sc.next();
        String b = sc.next();
        
        String str = String.format("%s%s", a, b);
        System.out.println(str);
    }
}