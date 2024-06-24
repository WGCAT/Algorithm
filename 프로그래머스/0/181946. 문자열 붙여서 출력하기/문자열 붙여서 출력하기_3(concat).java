import java.util.Scanner;


// String.concat()
public class Solution {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String a = sc.next();
        String b = sc.next();
        
        String str = a.concat(b);
        System.out.println(str);
    }
}
